package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

type IntermediateErr struct {
	MyError
}

func runJob(id string) IntermediateErr {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGlobalExec(jobBinPath)
	if err != nil {
		return IntermediateErr{wrapError(err, "")}
	} else if isExecutable == false {
		return IntermediateErr{wrapError(nil, "job binary is not executable")}
	} else {

	}

	return IntermediateErr{wrapError(exec.Command(jobBinPath, "--id="+id).Run(), "")}
}

type LowLevelErr struct {
	MyError
}

func isGlobalExec(path string) (bool, LowLevelErr) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(wrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, LowLevelErr{}
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)

	}
}
