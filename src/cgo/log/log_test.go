package log

import (
	"runtime"
	"testing"
)

func TestMyLog(t *testing.T) {
	logname := "Test.test"
	myLog := NewMyLog(logname)
	defer CloseMyLog(myLog)

	_, filename, line, _ := runtime.Caller(0)
	md5 := "ccf2b043d9390a1e995306c4c75584ee"
	now := 1431854915
	cv := "windows2389"
	myLog.LogDebug(logname, filename, line, "%s\t%d\t%s", md5, now, cv)
}
