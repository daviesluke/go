package utils

// Standard imports

import "bufio"
import "os"
import "os/signal"
import "regexp"
import "strings"
import "strconv"
import "syscall"

// Local imports

import "github.com/daviesluke/logger"

// Local functions 

type fn func() 

// Global functions 

func CheckRegEx(checkString string, regEx string) bool {
	logger.Tracef("Set regular expression -> %s", regEx)
	
	logger.Tracef("Checking to see if string %s matches regex", checkString)
	stringMatch, err := regexp.MatchString(regEx, checkString)
	if err != nil {
		logger.Errorf("Invalid regular exporession %s", regEx)
	}

	if stringMatch {
		logger.Trace("String matched!")
	} else {
		logger.Trace("String NOT matched!")
	}

	logger.Tracef("Returning %s", strconv.FormatBool(stringMatch))

	return stringMatch
}

func RemovePassword(checkString string) string {
	var userName        string

	logger.Info("Removing any passwords found ...")

	logger.Trace("Checking string for passwords ...")

	if CheckRegEx(checkString, ".+/[^@]+") {
		logger.Warn("Password in connection string - consider using SEPS")

		logger.Debug("Removing password for display ...")
	}

	upTokens := strings.SplitN(checkString, "/", 2)
	logger.Debugf("Split string using / delimiter into %d tokens.  Username is %s", len(upTokens), upTokens[0])

	userName = upTokens[0]
	logger.Tracef("Username set to %s", userName)

	passWord := upTokens[1]
	logger.Trace("Password set")

	if strings.Index(passWord,"@") != -1 {
		logger.Debug("Password portion contains a TNS Alias")

		passTokens := strings.SplitN(passWord, "@", 2)
		logger.Debugf("Split string using @ delimiter into %d tokens.  TNS Alias is %s", len(passTokens), passTokens[1])

		tnsAlias := passTokens[1]
		logger.Tracef("TNS Alias set to %s", tnsAlias)

		userName = strings.Join([]string{userName, tnsAlias},"@")
		logger.Tracef("Username now set to %s", userName)
	}

	logger.Infof("Process complete - returning %s", userName)

	return userName
}

func TrapSignal(runFunction fn) {
	logger.Infof("Trapping signals ...")
	
	channel := make(chan os.Signal)
	logger.Debug("Channel set for signals")

	signal.Notify(channel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	logger.Debug("Interrupt generated for signals 1,2,3 and 15")
	go func() {
		logger.Debug("About to block on signal input (In threaded process) ...")
		signalRecieved := <-channel
		logger.Infof("Received signal %d", signalRecieved)
		runFunction()
		logger.Debug("Exiting ...")
		os.Exit(1)
	}()

	logger.Infof("Process complete")
}

func LookupFile(searchFileName string, searchString string, searchIndex int, returnIndex int, delimiter string) string {
	logger.Info("Looking up string file")
	logger.Infof("Searching for %s in position %d in file %s demilited by %s ...", searchString, searchIndex, searchFileName, delimiter)

	logger.Tracef("Trying to open file %s ...", searchFileName)

	searchFile, err := os.Open(searchFileName)
	if err != nil {
		logger.Errorf("Unable to find file %s", searchFileName)
	}

	// Defer the close to auto close at the end of the procedure 

	defer searchFile.Close()
	logger.Tracef("Deferred closing of file %s at end of function", searchFileName)

	searchScanner := bufio.NewScanner(searchFile)
	logger.Tracef("Set up scanner for search file. Entering loop ...")

	lineNo := 0
	logger.Tracef("Set up variable LineNo and set to %d", lineNo)

	for searchScanner.Scan() {
		lineNo++
		logger.Tracef("Line number incremented to %d", lineNo)

		// Ignore blank lines and comments

		searchLine := strings.TrimSpace(searchScanner.Text())
		logger.Debugf("Trimmed search file line number %d contents - %s", lineNo, searchLine)

		if searchLine == "" || searchLine[0] == '#' {
			logger.Trace("Comment or blank line - ignoring line")
			continue
		}

		variableTokens := strings.Split(searchLine, delimiter)
		logger.Tracef("Line split into %d tokens using %s as delimiter", len(variableTokens), delimiter)

		if len(variableTokens) == 0 {
			logger.Trace("No tokens - ignoring line")
			continue
		}

		if strings.TrimSpace(variableTokens[searchIndex-1]) == searchString {
			if strings.TrimSpace(variableTokens[returnIndex-1]) != "" {
				logger.Infof("Search criteria found.  Returning entry %d => %s",  returnIndex, strings.TrimSpace(variableTokens[returnIndex-1]))
				return strings.TrimSpace(variableTokens[returnIndex-1])
			} else {
				logger.Debugf("NULL string in position %d in file %s for search string %s", returnIndex, searchFileName, searchString)
			}
		} else {
			logger.Tracef("String %s does not match search %s", strings.TrimSpace(variableTokens[searchIndex-1]), searchString)
		}
	}

	logger.Debug("Nothing found to return.  Returning empty string")
	return ""
}
