package logger

//This is a simple log framework.
//The common standard output.
//the common standard file output.
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

//extreme case
func TestStdLogger(t *testing.T) {
	logger := NewStdLogger(false, false, false, false, false)

	flags := logger.logger.Flags()
	if flags != 0 {
		t.Fatalf("Expected %q, received %q\n", 0, flags)
	}

	if logger.debug {
		t.Fatalf("Expected %t, received %t\n", false, logger.debug)
	}

	if logger.trace {
		t.Fatalf("Expected %t, received %t\n", false, logger.trace)
	}

	logger1 := NewStdLogger(true, true, true, true, true)

	flags1 := logger1.logger.Flags()
	if !logger1.debug {
		t.Fatalf("Expected %t, received %t\n", true, logger.debug)
	}

	logger.Tracef("show me the debug info %s, %d", ":input error", flags1)
	//
	logger.Noticef("show me the debug info %s, %d", ":input error", flags1)
 }

 //Debug trace and time
func TestStdLoggerWithDebugTraceAndTime(t *testing.T) {
	logger := NewStdLogger(true, true, true, false, false)

	flags := logger.logger.Flags()
	if flags != log.LstdFlags|log.Lmicroseconds {
		t.Fatalf("Expected %d, received %d\n", log.LstdFlags, flags)
	}

	if !logger.debug {
		t.Fatalf("Expected %t, received %t\n", true, logger.debug)
	}

	if !logger.trace {
		t.Fatalf("Expected %t, received %t\n", true, logger.trace)
	}
}

// func Logger notice
func TestStdLoggerNotice(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Noticef("foo")
	}, "[INF] foo\n")
}

// func Logger notice with color
func TestStdLoggerNoticeWithColor(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, true, false)
		logger.Noticef("foo")
	}, "[\x1b[32mINF\x1b[0m] foo\n")
}

// func Logger debug
func TestStdLoggerDebug(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, true, false, false, false)
		logger.Debugf("foo %s", "bar")
	}, "[DBG] foo bar\n")
}
// func Logger without debug
func TestStdLoggerDebugWithOutDebug(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Debugf("foo")
	}, "")
}
// func Logger trace
func TestStdLoggerTrace(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, true, false, false)
		logger.Tracef("foo")
	}, "[TRC] foo\n")
}

//func trace without debug
func TestStdLoggerTraceWithOutDebug(t *testing.T) {
	expectOutput(t, func() {
		logger := NewStdLogger(false, false, false, false, false)
		logger.Tracef("foo")
	}, "")
}

//log file 1
func TestFileLogger(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "_test1")
	if err != nil {
		t.Fatal("Could not create tmp dir")
	}
	defer os.RemoveAll(tmpDir)

	file, err := ioutil.TempFile(tmpDir, "t_:log_")
	if err != nil {
		t.Fatalf("Could not create the temp file: %v", err)
	}
	file.Close()

	logger := NewFileLogger(file.Name(), false, false, false, false)
	logger.Noticef("foo")

	buf, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Could not read logfile: %v", err)
	}
	if len(buf) <= 0 {
		t.Fatal("Expected a non-zero length logfile")
	}

	if string(buf) != "[INF] foo\n" {
		t.Fatalf("Expected '%s', received '%s'\n", "[INFO] foo", string(buf))
	}

	file, err = ioutil.TempFile(tmpDir, "t_:log_")
	if err != nil {
		t.Fatalf("Could not create the temp file: %v", err)
	}
	file.Close()

	logger = NewFileLogger(file.Name(), true, true, true, true)
	logger.Errorf("foo")

	buf, err = ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Could not read logfile: %v", err)
	}
	if len(buf) <= 0 {
		t.Fatal("Expected a non-zero length logfile")
	}
	str := string(buf)
	errMsg := fmt.Sprintf("Expected '%s', received '%s'\n", "[pid] <date> [ERR] foo", str)
	pidEnd := strings.Index(str, " ")
	infoStart := strings.LastIndex(str, "[ERR]")
	if pidEnd == -1 || infoStart == -1 {
		t.Fatalf("%v", errMsg)
	}
	pid := str[0:pidEnd]
	if pid[0] != '[' || pid[len(pid)-1] != ']' {
		t.Fatalf("%v", errMsg)
	}
	//TODO: Parse date.
	if !strings.HasSuffix(str, "[ERR] foo\n") {
		t.Fatalf("%v", errMsg)
	}
}

func expectOutput(t *testing.T, f func(), expected string) {
	old := os.Stderr // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stderr = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	os.Stderr.Close()
	os.Stderr = old // restoring the real stdout
	out := <-outC
	if out != expected {
		t.Fatalf("Expected '%s', received '%s'\n", expected, out)
	}
}

//created the log file with debug pattern
func TestNewFileWithSpecifiedFileName(t *testing.T){
	file, err := ioutil.TempFile("./", "test1-")
	if err != nil {
		t.Fatalf("Could not create the temp file: %v", err)
	}
	file.Close()

	logger := NewFileLogger(file.Name(), true, true, true, true)
	logger.Noticef("Good ! PERFECT")
}

//created the log file with debug pattern
func TestNewFileWithHardcodedName(t *testing.T){

	logger := NewFileLogger("test--1.log", true, true, true, true)
	logger.Noticef("Good ! PERFECT")
}

