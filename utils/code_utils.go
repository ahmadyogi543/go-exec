package utils

import (
	"fmt"
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
		fmt.Println("Error di stdinpipe")
	}
	execStdout, err := execCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error di stdoutpipe")
	}
	execStderr, err := execCmd.StderrPipe()
	if err != nil {
		fmt.Println("Error di stderrpipe")
	}

	execCmd.Start()
	execStdin.Write([]byte(input))
	execStdin.Close()

	stdOutOutput, _ := io.ReadAll(execStdout)
	stdErrOutput, _ := io.ReadAll(execStderr)
	execCmd.Wait()

	codeOutput := string(stdOutOutput)
	if string(stdErrOutput) != "" {
		codeOutput = string(stdErrOutput)
	}

	return codeOutput, mainErr
}
