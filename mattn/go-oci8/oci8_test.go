package oci8

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

// to run database tests
// go test -v github.com/mattn/go-oci8 -args -disableDatabase=false -hostValid type_hostname_here -username type_username_here -password "type_password_here"
// note minimum Go version for testing is 1.8

/* note that testing needs an Oracle user and the following:
create or replace function TYPE_USER_HERE.SLEEP_SECONDS (p_seconds number) return integer is
begin
  dbms_lock.sleep(p_seconds);
  return 1;
end SLEEP_SECONDS;
/
*/

var (
	TestDisableDatabase      bool
	TestHostValid            string
	TestHostInvalid          string
	TestUsername             string
	TestPassword             string
	TestContextTimeoutString string
	TestContextTimeout       time.Duration
	TestDatabase             string
	TestDisableDestructive   bool

	TestTimeString string

	TestDB *sql.DB

	TestTypeTime      = reflect.TypeOf(time.Time{})
	TestTypeByteSlice = reflect.TypeOf([]byte{})

	testString1    string
	testByteSlice1 []byte

	testTimeLocUTC *time.Location
	testTimeLocGMT *time.Location
	testTimeLocEST *time.Location
	testTimeLocMST *time.Location
	testTimeLocNZ  *time.Location
)

// testQueryResults is for testing a query
type testQueryResults struct {
	query   string
	args    [][]interface{}
	results [][][]interface{}
}

// TestMain sets up testing
func TestMain(m *testing.M) {
	code := setupForTesting()
	if code != 0 {
		os.Exit(code)
	}
	code = m.Run()

	if !TestDisableDatabase {
		err := TestDB.Close()
		if err != nil {
			fmt.Println("close error:", err)
			os.Exit(2)
		}
	}

	os.Exit(code)
}

// setupForTesting sets up flags and connects to test database
func setupForTesting() int {
	flag.BoolVar(&TestDisableDatabase, "disableDatabase", true, "set to true to disable the Oracle tests")
	flag.StringVar(&TestHostValid, "hostValid", "127.0.0.1", "a host where a Oracle database is running")
	flag.StringVar(&TestHostInvalid, "hostInvalid", "169.254.200.200", "a host where a Oracle database is not running")
	flag.StringVar(&TestUsername, "username", "", "the username for the Oracle database")
	flag.StringVar(&TestPassword, "password", "", "the password for the Oracle database")
	flag.StringVar(&TestContextTimeoutString, "contextTimeout", "30s", "the context timeout for queries")
	flag.BoolVar(&TestDisableDestructive, "disableDestructive", false, "set to true to disable the destructive Oracle tests")

	flag.Parse()

	var err error
	TestContextTimeout, err = time.ParseDuration(TestContextTimeoutString)
	if err != nil {
		fmt.Println("parse context timeout error:", err)
		return 4
	}

	if !TestDisableDatabase {
		TestDB = testGetDB()
		if TestDB == nil {
			return 6
		}
	}

	TestTimeString = time.Now().UTC().Format("20060102150405")

	var i int
	var buffer bytes.Buffer
	for i = 0; i < 1000; i++ {
		buffer.WriteRune(rune(i))
	}
	testString1 = buffer.String()

	testByteSlice1 = make([]byte, 2000)
	for i = 0; i < 2000; i++ {
		testByteSlice1[i] = byte(i)
	}

	testTimeLocUTC, _ = time.LoadLocation("UTC")
	testTimeLocGMT, _ = time.LoadLocation("GMT")
	testTimeLocEST, _ = time.LoadLocation("EST")
	testTimeLocMST, _ = time.LoadLocation("MST")
	testTimeLocNZ, _ = time.LoadLocation("NZ")

	return 0
}

// TestParseDSN tests parsing the DSN
func TestParseDSN(t *testing.T) {
	var (
		pacific *time.Location
		err     error
	)

	if pacific, err = time.LoadLocation("America/Los_Angeles"); err != nil {
		panic(err)
	}
	var dsnTests = []struct {
		dsnString   string
		expectedDSN *DSN
	}{
		{"oracle://xxmc:xxmc@107.20.30.169:1521/ORCL?loc=America%2FLos_Angeles", &DSN{Username: "xxmc", Password: "xxmc", Connect: "107.20.30.169:1521/ORCL", prefetchRows: 10, Location: pacific}},
		{"xxmc/xxmc@107.20.30.169:1521/ORCL?loc=America%2FLos_Angeles", &DSN{Username: "xxmc", Password: "xxmc", Connect: "107.20.30.169:1521/ORCL", prefetchRows: 10, Location: pacific}},
		{"sys/syspwd@107.20.30.169:1521/ORCL?loc=America%2FLos_Angeles&as=sysdba", &DSN{Username: "sys", Password: "syspwd", Connect: "107.20.30.169:1521/ORCL", prefetchRows: 10, Location: pacific, operationMode: 0x00000002}}, // with operationMode: 0x00000002 = C.OCI_SYDBA
		{"xxmc/xxmc@107.20.30.169:1521/ORCL", &DSN{Username: "xxmc", Password: "xxmc", Connect: "107.20.30.169:1521/ORCL", prefetchRows: 10, Location: time.Local}},
		{"xxmc/xxmc@107.20.30.169/ORCL", &DSN{Username: "xxmc", Password: "xxmc", Connect: "107.20.30.169/ORCL", prefetchRows: 10, Location: time.Local}},
	}

	for _, tt := range dsnTests {
		actualDSN, err := ParseDSN(tt.dsnString)

		if err != nil {
			t.Errorf("ParseDSN(%s) got error: %+v", tt.dsnString, err)
		}

		if !reflect.DeepEqual(actualDSN, tt.expectedDSN) {
			t.Errorf("ParseDSN(%s): expected %+v, actual %+v", tt.dsnString, tt.expectedDSN, actualDSN)
		}
	}
}

// TestIsBadConn tests bad connection error codes
func TestIsBadConn(t *testing.T) {
	var errorCode = "ORA-03114"
	if !isBadConnection(errorCode) {
		t.Errorf("TestIsBadConn: expected %+v, actual %+v", true, isBadConnection(errorCode))
	}
}

var sql1 = `create table foo(
	c1 varchar2(256),
	c2 nvarchar2(256),
	c3 number,
	c4 float,
	c6 date,
	c7 BINARY_FLOAT,
	c8 BINARY_DOUBLE,
	c9 TIMESTAMP,
	c10 TIMESTAMP WITH TIME ZONE,
	c11 TIMESTAMP WITH LOCAL TIME ZONE,
	c12 INTERVAL YEAR TO MONTH,
	c13 INTERVAL DAY TO SECOND,
	c14 RAW(80),
	c15 ROWID,
	c17 CHAR(15),
	c18 NCHAR(20),
	c19 CLOB,
	c21 BLOB,
	cend varchar2(12)
	)`

var sql12 = `insert( c1,c2,c3,c4,c6,c7,c8,c9,c10,c11,c12,c13,c14,c17,c18,c19,c20,c21,cend) into foo values(
:1,
:2,
:3,
:4,
:6,
:7,
:8,
:9,
:10,
:11,
NUMTOYMINTERVAL( :12, 'MONTH'),
NUMTODSINTERVAL( :13 / 1000000000, 'SECOND'),
:14,
:17,
:18,
:19,
:21,
'END'
)`
