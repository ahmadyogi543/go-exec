package utils

import (
	"io"
	"os/exec"

	"github.com/google/uuid"
)

func CreateCode(code string, extension string) (string, error) {
	path := "code"
	filename := uuid.New().String() + "." + extension
	return WriteFile(path, filename, code)
}

func DeleteCode(path string) error {
	err := RemoveFile(path)
	return err
}

func ExecuteCode(path string, input string, executable string) (string, error) {
	var mainErr error
	execCmd := exec.Command(executable, path)

	execStdin, err := execCmd.StdinPipe()
	if err != nil {
		mainErr = err
	}
	execStdout, err := execCmd.StdoutPipe()
	if err != nil {
		mainErr = err
	}
	execStderr, err := execCmd.StderrPipe()
	if err != nil {
		mainErr = err
	}

	execCmd.Start()
	execStdin.Write([]byte(input))
	execStdin.Close()

	stdOutOutput, err := (io.ReadAll(execStdout))
	if err != nil {
		mainErr = err
	}
	stdErrOutput, err := (io.ReadAll(execStderr))
	if err != nil {
		mainErr = err
	}
	execCmd.Wait()

	codeOutput := string(stdOutOutput)
	if string(stdErrOutput) != "" {
		codeOutput = string(stdErrOutput)
	}

	return codeOutput, mainErr
}
