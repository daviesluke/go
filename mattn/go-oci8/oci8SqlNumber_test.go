package oci8

import (
	"testing"
)

// TestSelectCastNumber checks cast x from dual number types
func TestSelectCastNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	// https://ss64.com/ora/syntax-datatypes.html

	queryResults := []testQueryResults{

		// NUMBER(38,10)
		testQueryResults{
			query: "select cast (:1 as NUMBER(38,10)) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-99999999999999999999999999.9999999999)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.9873046875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.76171875)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.76171875)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.9873046875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(99999999999999999999999999.9999999999)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-99999999999999999999999999.9999999999)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.9873046875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.76171875)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.76171875)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.9873046875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(99999999999999999999999999.9999999999)}},
			},
		},

		// DEC(38,10)
		testQueryResults{
			query: "select cast (:1 as DEC(38,10)) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-99999999999999999999999999.9999999999)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.9873046875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.76171875)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.76171875)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.9873046875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(99999999999999999999999999.9999999999)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-99999999999999999999999999.9999999999)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.9873046875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.76171875)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.76171875)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.9873046875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(99999999999999999999999999.9999999999)}},
			},
		},

		// DECIMAL(38,10)
		testQueryResults{
			query: "select cast (:1 as DECIMAL(38,10)) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-99999999999999999999999999.9999999999)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.9873046875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.76171875)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.76171875)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.9873046875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(99999999999999999999999999.9999999999)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-99999999999999999999999999.9999999999)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.9873046875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.76171875)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.76171875)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.9873046875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(99999999999999999999999999.9999999999)}},
			},
		},

		// NUMERIC(38,10)
		testQueryResults{
			query: "select cast (:1 as NUMERIC(38,10)) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-99999999999999999999999999.9999999999)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.9873046875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.76171875)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.76171875)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.9873046875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(99999999999999999999999999.9999999999)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-99999999999999999999999999.9999999999)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.9873046875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.76171875)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.76171875)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.9873046875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(99999999999999999999999999.9999999999)}},
			},
		},

		// FLOAT
		testQueryResults{
			query: "select cast (:1 as FLOAT) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-288230381928101358902502915674136903680)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.99999988079071044921875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.00415134616196155548095703125)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.00415134616196155548095703125)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.99999988079071044921875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(288230381928101358902502915674136903680)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-288230381928101358902502915674136903680)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(288230381928101358902502915674136903680)}},
			},
		},

		// INTEGER
		testQueryResults{
			query: "select cast (:1 as INTEGER) from dual",
			args: [][]interface{}{
				[]interface{}{int64(-2147483648)},
				[]interface{}{int64(-1)},
				[]interface{}{int64(0)},
				[]interface{}{int64(1)},
				[]interface{}{int64(2147483647)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-2147483648)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
			},
		},

		// INT
		testQueryResults{
			query: "select cast (:1 as INT) from dual",
			args: [][]interface{}{
				[]interface{}{int64(-2147483648)},
				[]interface{}{int64(-1)},
				[]interface{}{int64(0)},
				[]interface{}{int64(1)},
				[]interface{}{int64(2147483647)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-2147483648)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
			},
		},

		// SMALLINT
		testQueryResults{
			query: "select cast (:1 as SMALLINT) from dual",
			args: [][]interface{}{
				[]interface{}{int64(-2147483648)},
				[]interface{}{int64(-1)},
				[]interface{}{int64(0)},
				[]interface{}{int64(1)},
				[]interface{}{int64(2147483647)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-2147483648)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
			},
		},

		// REAL
		testQueryResults{
			query: "select cast (:1 as REAL) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-288230381928101358902502915674136903680)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.99999988079071044921875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.00415134616196155548095703125)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.00415134616196155548095703125)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.99999988079071044921875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(288230381928101358902502915674136903680)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-288230381928101358902502915674136903680)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(288230381928101358902502915674136903680)}},
			},
		},

		// BINARY_FLOAT
		testQueryResults{
			query: "select cast (:1 as BINARY_FLOAT) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-288230381928101358902502915674136903680)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.99999988079071044921875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.00415134616196155548095703125)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.00415134616196155548095703125)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.99999988079071044921875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483648)},
				[]interface{}{float64(288230381928101358902502915674136903680)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-288230381928101358902502915674136903680)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483648)}},
				[][]interface{}{[]interface{}{float64(288230381928101358902502915674136903680)}},
			},
		},

		// BINARY_DOUBLE
		testQueryResults{
			query: "select cast (:1 as BINARY_DOUBLE) from dual",
			args: [][]interface{}{
				[]interface{}{float64(-288230381928101358902502915674136903680)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-123456792)},
				[]interface{}{float64(-1.99999988079071044921875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.00415134616196155548095703125)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.00415134616196155548095703125)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.99999988079071044921875)},
				[]interface{}{float64(123456792)},
				[]interface{}{float64(2147483647)},
				[]interface{}{float64(288230381928101358902502915674136903680)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-288230381928101358902502915674136903680)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-123456792)}},
				[][]interface{}{[]interface{}{float64(-1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.00415134616196155548095703125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.99999988079071044921875)}},
				[][]interface{}{[]interface{}{float64(123456792)}},
				[][]interface{}{[]interface{}{float64(2147483647)}},
				[][]interface{}{[]interface{}{float64(288230381928101358902502915674136903680)}},
			},
		},
	}

	testRunQueryResults(t, queryResults)
}

// TestSelectGoTypesNumber is select :1 from dual for each number Go Type
func TestSelectGoTypesNumber(t *testing.T) {
	if TestDisableDatabase {
		t.SkipNow()
	}

	// https://tour.golang.org/basics/11

	queryResults := []testQueryResults{
		// bool
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{true},
				[]interface{}{false},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(0)}},
			},
		},

		// int8: -128 to 127
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{int16(-128)},
				[]interface{}{int16(-1)},
				[]interface{}{int16(0)},
				[]interface{}{int16(1)},
				[]interface{}{int16(127)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-128)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
			},
		},
		// int16: -32768 to 32767
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{int16(-32768)},
				[]interface{}{int16(-128)},
				[]interface{}{int16(-1)},
				[]interface{}{int16(0)},
				[]interface{}{int16(1)},
				[]interface{}{int16(127)},
				[]interface{}{int16(32767)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-32768)}},
				[][]interface{}{[]interface{}{int64(-128)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
			},
		},
		// int32: -2147483648 to 2147483647
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{int32(-2147483648)},
				[]interface{}{int32(-32768)},
				[]interface{}{int32(-128)},
				[]interface{}{int32(-1)},
				[]interface{}{int32(0)},
				[]interface{}{int32(1)},
				[]interface{}{int32(127)},
				[]interface{}{int32(32767)},
				[]interface{}{int32(2147483647)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-2147483648)}},
				[][]interface{}{[]interface{}{int64(-32768)}},
				[][]interface{}{[]interface{}{int64(-128)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
			},
		},
		// int64: -9223372036854775808 to 9223372036854775807
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{int64(-9223372036854775808)},
				[]interface{}{int64(-2147483648)},
				[]interface{}{int64(-32768)},
				[]interface{}{int64(-128)},
				[]interface{}{int64(-1)},
				[]interface{}{int64(0)},
				[]interface{}{int64(1)},
				[]interface{}{int64(127)},
				[]interface{}{int64(32767)},
				[]interface{}{int64(2147483647)},
				[]interface{}{int64(9223372036854775807)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(-9223372036854775808)}},
				[][]interface{}{[]interface{}{int64(-2147483648)}},
				[][]interface{}{[]interface{}{int64(-32768)}},
				[][]interface{}{[]interface{}{int64(-128)}},
				[][]interface{}{[]interface{}{int64(-1)}},
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
				[][]interface{}{[]interface{}{int64(9223372036854775807)}},
			},
		},

		// uint8: 0 to 255
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{uint32(0)},
				[]interface{}{uint32(1)},
				[]interface{}{uint32(127)},
				[]interface{}{uint32(255)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(255)}},
			},
		},
		// uint16: 0 to 65535
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{uint32(0)},
				[]interface{}{uint32(1)},
				[]interface{}{uint32(127)},
				[]interface{}{uint32(32767)},
				[]interface{}{uint32(65535)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
				[][]interface{}{[]interface{}{int64(65535)}},
			},
		},
		// uint32: 0 to 4294967295
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{uint32(0)},
				[]interface{}{uint32(1)},
				[]interface{}{uint32(127)},
				[]interface{}{uint32(32767)},
				[]interface{}{uint32(2147483647)},
				[]interface{}{uint32(4294967295)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
				[][]interface{}{[]interface{}{int64(4294967295)}},
			},
		},
		// uint64: 0 to 18446744073709551615
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{uint64(0)},
				[]interface{}{uint64(1)},
				[]interface{}{uint64(127)},
				[]interface{}{uint64(32767)},
				[]interface{}{uint64(2147483647)},
				[]interface{}{uint64(9223372036854775807)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{int64(0)}},
				[][]interface{}{[]interface{}{int64(1)}},
				[][]interface{}{[]interface{}{int64(127)}},
				[][]interface{}{[]interface{}{int64(32767)}},
				[][]interface{}{[]interface{}{int64(2147483647)}},
				[][]interface{}{[]interface{}{int64(9223372036854775807)}},
			},
		},

		// float32
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{float32(-9223372036854775808)},
				[]interface{}{float32(-2147483648)},
				[]interface{}{float32(-32767.123046875)},
				[]interface{}{float32(-32767)},
				[]interface{}{float32(-128.1234588623046875)},
				[]interface{}{float32(-128)},
				[]interface{}{float32(-1.12345683574676513671875)},
				[]interface{}{float32(-1)},
				[]interface{}{float32(-0.12345679104328155517578125)},
				[]interface{}{float32(0)},
				[]interface{}{float32(0.12345679104328155517578125)},
				[]interface{}{float32(1)},
				[]interface{}{float32(1.12345683574676513671875)},
				[]interface{}{float32(128)},
				[]interface{}{float32(128.1234588623046875)},
				[]interface{}{float32(32767)},
				[]interface{}{float32(32767.123046875)},
				[]interface{}{float32(2147483648)},
				[]interface{}{float32(9223372036854775808)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-9223372036854775808)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-32767.123046875)}},
				[][]interface{}{[]interface{}{float64(-32767)}},
				[][]interface{}{[]interface{}{float64(-128.1234588623046875)}},
				[][]interface{}{[]interface{}{float64(-128)}},
				[][]interface{}{[]interface{}{float64(-1.12345683574676513671875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.12345679104328155517578125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.12345679104328155517578125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.12345683574676513671875)}},
				[][]interface{}{[]interface{}{float64(128)}},
				[][]interface{}{[]interface{}{float64(128.1234588623046875)}},
				[][]interface{}{[]interface{}{float64(32767)}},
				[][]interface{}{[]interface{}{float64(32767.123046875)}},
				[][]interface{}{[]interface{}{float64(2147483648)}},
				[][]interface{}{[]interface{}{float64(9223372036854775808)}},
			},
		},

		// float64
		testQueryResults{
			query: "select :1 from dual",
			args: [][]interface{}{
				[]interface{}{float64(-9223372036854775808)},
				[]interface{}{float64(-2147483648)},
				[]interface{}{float64(-32767.123046875)},
				[]interface{}{float64(-32767)},
				[]interface{}{float64(-128.1234588623046875)},
				[]interface{}{float64(-128)},
				[]interface{}{float64(-1.12345683574676513671875)},
				[]interface{}{float64(-1)},
				[]interface{}{float64(-0.12345679104328155517578125)},
				[]interface{}{float64(0)},
				[]interface{}{float64(0.12345679104328155517578125)},
				[]interface{}{float64(1)},
				[]interface{}{float64(1.12345683574676513671875)},
				[]interface{}{float64(128)},
				[]interface{}{float64(128.1234588623046875)},
				[]interface{}{float64(32767)},
				[]interface{}{float64(32767.123046875)},
				[]interface{}{float64(2147483648)},
				[]interface{}{float64(9223372036854775808)},
			},
			results: [][][]interface{}{
				[][]interface{}{[]interface{}{float64(-9223372036854775808)}},
				[][]interface{}{[]interface{}{float64(-2147483648)}},
				[][]interface{}{[]interface{}{float64(-32767.123046875)}},
				[][]interface{}{[]interface{}{float64(-32767)}},
				[][]interface{}{[]interface{}{float64(-128.1234588623046875)}},
				[][]interface{}{[]interface{}{float64(-128)}},
				[][]interface{}{[]interface{}{float64(-1.12345683574676513671875)}},
				[][]interface{}{[]interface{}{float64(-1)}},
				[][]interface{}{[]interface{}{float64(-0.12345679104328155517578125)}},
				[][]interface{}{[]interface{}{float64(0)}},
				[][]interface{}{[]interface{}{float64(0.12345679104328155517578125)}},
				[][]interface{}{[]interface{}{float64(1)}},
				[][]interface{}{[]interface{}{float64(1.12345683574676513671875)}},
				[][]interface{}{[]interface{}{float64(128)}},
				[][]interface{}{[]interface{}{float64(128.1234588623046875)}},
				[][]interface{}{[]interface{}{float64(32767)}},
				[][]interface{}{[]interface{}{float64(32767.123046875)}},
				[][]interface{}{[]interface{}{float64(2147483648)}},
				[][]interface{}{[]interface{}{float64(9223372036854775808)}},
			},
		},
	}

	testRunQueryResults(t, queryResults)
}

// TestDestructiveNumber checks insert, select, update, and delete of number types
func TestDestructiveNumber(t *testing.T) {
	if TestDisableDatabase || TestDisableDestructive {
		t.SkipNow()
	}

	// https://ss64.com/ora/syntax-datatypes.html

	// NUMBER negative
	err := testExec(t, "create table NUMBER_"+TestTimeString+
		" ( A NUMBER(10,2), B NUMBER(20,4), C NUMBER(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table NUMBER_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into NUMBER_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults := []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMBER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from NUMBER_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMBER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// test truncate
	err = testExec(t, "truncate table NUMBER_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMBER_" + TestTimeString,
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMBER positive
	err = testExecRows(t, "insert into NUMBER_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMBER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from NUMBER_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMBER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DEC negative
	err = testExec(t, "create table DEC_"+TestTimeString+
		" ( A DEC(10,2), B DEC(20,4), C DEC(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table DEC_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into DEC_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DEC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from DEC_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DEC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DEC positive
	err = testExec(t, "truncate table DEC_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into DEC_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DEC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from DEC_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DEC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DECIMAL negative
	err = testExec(t, "create table DECIMAL_"+TestTimeString+
		" ( A DECIMAL(10,2), B DECIMAL(20,4), C DECIMAL(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table DECIMAL_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into DECIMAL_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DECIMAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from DECIMAL_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DECIMAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// DECIMAL positive
	err = testExec(t, "truncate table DECIMAL_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into DECIMAL_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DECIMAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from DECIMAL_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from DECIMAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMERIC negative
	err = testExec(t, "create table NUMERIC_"+TestTimeString+
		" ( A NUMERIC(10,2), B NUMERIC(20,4), C NUMERIC(38,8) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table NUMERIC_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into NUMERIC_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMERIC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from NUMERIC_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMERIC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// NUMERIC positive
	err = testExec(t, "truncate table NUMERIC_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into NUMERIC_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMERIC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from NUMERIC_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from NUMERIC_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// FLOAT negative
	err = testExec(t, "create table FLOAT_"+TestTimeString+
		" ( A FLOAT(28), B FLOAT(32), C FLOAT(38) )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table FLOAT_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into FLOAT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from FLOAT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// FLOAT positive
	err = testExec(t, "truncate table FLOAT_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into FLOAT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from FLOAT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INTEGER negative
	err = testExec(t, "create table INTEGER_"+TestTimeString+
		" ( A INTEGER, B INTEGER, C INTEGER )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table INTEGER_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into INTEGER_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INTEGER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					[]interface{}{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from INTEGER_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INTEGER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INTEGER positive
	err = testExec(t, "truncate table INTEGER_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into INTEGER_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INTEGER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					[]interface{}{int64(12345679), int64(123456792), int64(123456792)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from INTEGER_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{10000000},
			[]interface{}{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INTEGER_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INT negative
	err = testExec(t, "create table INT_"+TestTimeString+
		" ( A INT, B INT, C INT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table INT_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into INT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					[]interface{}{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from INT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INT positive
	err = testExec(t, "truncate table INT_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into INT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					[]interface{}{int64(12345679), int64(123456792), int64(123456792)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from INT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{10000000},
			[]interface{}{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from INT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// SMALLINT negative
	err = testExec(t, "create table SMALLINT_"+TestTimeString+
		" ( A SMALLINT, B SMALLINT, C SMALLINT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table SMALLINT_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into SMALLINT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from SMALLINT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-21474836), int64(-2147483648), int64(-2147483648)},
					[]interface{}{int64(-10000000), int64(-1000000000000000), int64(-1000000000000000)},
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from SMALLINT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-10000000},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from SMALLINT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(-1234567), int64(-123456792), int64(-123456792)},
					[]interface{}{int64(-2), int64(-2), int64(-2)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
					[]interface{}{int64(-1), int64(-1), int64(-1)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// INT positive
	err = testExec(t, "truncate table SMALLINT_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into SMALLINT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from SMALLINT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(10000000), int64(1000000000000000), int64(1000000000000000)},
					[]interface{}{int64(12345679), int64(123456792), int64(123456792)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from SMALLINT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{10000000},
			[]interface{}{12345679},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from SMALLINT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{int64(0), int64(0), int64(0)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(1), int64(1), int64(1)},
					[]interface{}{int64(2), int64(2), int64(2)},
					[]interface{}{int64(21474836), int64(2147483647), int64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// REAL negative
	err = testExec(t, "create table REAL_"+TestTimeString+
		" ( A REAL, B REAL, C REAL )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table REAL_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into REAL_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from REAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from REAL_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from REAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// REAL positive
	err = testExec(t, "truncate table REAL_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into REAL_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from REAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from REAL_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from REAL_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT negative
	err = testExec(t, "create table BINARY_FLOAT_"+TestTimeString+
		" ( A BINARY_FLOAT, B BINARY_FLOAT, C BINARY_FLOAT )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table BINARY_FLOAT_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into BINARY_FLOAT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680)},
			[]interface{}{-2147483648, -2147483648, -2147483648},
			[]interface{}{-123456792, -123456792, -123456792},
			[]interface{}{-1.99999988079071044921875, -1.99999988079071044921875, -1.99999988079071044921875},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.00415134616196155548095703125, -0.00415134616196155548095703125, -0.00415134616196155548095703125},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680), float64(-288230381928101358902502915674136903680)},
					[]interface{}{float64(-2147483648), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-123456792), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.99999988079071044921875), float64(-1.99999988079071044921875), float64(-1.99999988079071044921875)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from BINARY_FLOAT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-2147483648},
			[]interface{}{float64(-288230381928101358902502915674136903680)},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-123456792), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.99999988079071044921875), float64(-1.99999988079071044921875), float64(-1.99999988079071044921875)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125), float64(-0.00415134616196155548095703125)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_FLOAT positive
	err = testExec(t, "truncate table BINARY_FLOAT_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into BINARY_FLOAT_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.00415134616196155548095703125, 0.00415134616196155548095703125, 0.00415134616196155548095703125},
			[]interface{}{1, 1, 1},
			[]interface{}{1.99999988079071044921875, 1.99999988079071044921875, 1.99999988079071044921875},
			[]interface{}{123456792, 123456792, 123456792},
			[]interface{}{2147483648, 2147483648, 2147483648},
			[]interface{}{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.99999988079071044921875), float64(1.99999988079071044921875), float64(1.99999988079071044921875)},
					[]interface{}{float64(123456792), float64(123456792), float64(123456792)},
					[]interface{}{float64(2147483648), float64(2147483648), float64(2147483648)},
					[]interface{}{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from BINARY_FLOAT_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_FLOAT_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125), float64(0.00415134616196155548095703125)},
					[]interface{}{float64(1.99999988079071044921875), float64(1.99999988079071044921875), float64(1.99999988079071044921875)},
					[]interface{}{float64(123456792), float64(123456792), float64(123456792)},
					[]interface{}{float64(2147483648), float64(2147483648), float64(2147483648)},
					[]interface{}{float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680), float64(288230381928101358902502915674136903680)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE negative
	err = testExec(t, "create table BINARY_DOUBLE_"+TestTimeString+
		" ( A BINARY_DOUBLE, B BINARY_DOUBLE, C BINARY_DOUBLE )", nil)
	if err != nil {
		t.Fatal("create table error:", err)
	}

	defer func() {
		err = testExec(t, "drop table BINARY_DOUBLE_"+TestTimeString, nil)
		if err != nil {
			t.Error("drop table error:", err)
		}
	}()

	err = testExecRows(t, "insert into BINARY_DOUBLE_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{-9999999.99, -999999999999999.9999, -9999999999999999999999999.99999999},
			[]interface{}{-21474836, -2147483648, -2147483648},
			[]interface{}{-1234567, -123456792, -123456792},
			[]interface{}{-1.98, -1.9873, -1.98730468},
			[]interface{}{-1, -1, -1},
			[]interface{}{-0.76, -0.7617, -0.76171875},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_DOUBLE_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-21474836), float64(-2147483648), float64(-2147483648)},
					[]interface{}{float64(-9999999.99), float64(-999999999999999.9999), float64(-9999999999999999999999999.99999999)},
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from BINARY_DOUBLE_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{-21474836},
			[]interface{}{-9999999.99},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_DOUBLE_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(-1234567), float64(-123456792), float64(-123456792)},
					[]interface{}{float64(-1.98), float64(-1.9873), float64(-1.98730468)},
					[]interface{}{float64(-1), float64(-1), float64(-1)},
					[]interface{}{float64(-0.76), float64(-0.7617), float64(-0.76171875)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	// BINARY_DOUBLE positive
	err = testExec(t, "truncate table BINARY_DOUBLE_"+TestTimeString, nil)
	if err != nil {
		t.Error("truncate error:", err)
	}

	err = testExecRows(t, "insert into BINARY_DOUBLE_"+TestTimeString+" ( A, B, C ) values (:1, :2, :3)",
		[][]interface{}{
			[]interface{}{0, 0, 0},
			[]interface{}{0.76, 0.7617, 0.76171875},
			[]interface{}{1, 1, 1},
			[]interface{}{1.98, 1.9873, 1.98730468},
			[]interface{}{12345679, 123456792, 123456792},
			[]interface{}{21474836, 2147483647, 2147483647},
			[]interface{}{9999999.99, 999999999999999.9999, 99999999999999999999999999.99999999},
		})
	if err != nil {
		t.Error("insert error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_DOUBLE_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0), float64(0), float64(0)},
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1), float64(1), float64(1)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)

	err = testExecRows(t, "delete from BINARY_DOUBLE_"+TestTimeString+" where A = :1",
		[][]interface{}{
			[]interface{}{0},
			[]interface{}{1},
		})
	if err != nil {
		t.Error("delete error:", err)
	}

	queryResults = []testQueryResults{
		testQueryResults{
			query: "select A, B, C from BINARY_DOUBLE_" + TestTimeString + " order by A",
			args:  [][]interface{}{[]interface{}{}},
			results: [][][]interface{}{
				[][]interface{}{
					[]interface{}{float64(0.76), float64(0.7617), float64(0.76171875)},
					[]interface{}{float64(1.98), float64(1.9873), float64(1.98730468)},
					[]interface{}{float64(9999999.99), float64(999999999999999.9999), float64(99999999999999999999999999.99999999)},
					[]interface{}{float64(12345679), float64(123456792), float64(123456792)},
					[]interface{}{float64(21474836), float64(2147483647), float64(2147483647)},
				},
			},
		},
	}
	testRunQueryResults(t, queryResults)
}
