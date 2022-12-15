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

	// execStdin, _ := execCmd.StdinPipe()
	execStdout, _ := execCmd.StdoutPipe()
	execStderr, _ := execCmd.StderrPipe()

	execCmd.Start()
	// execStdin.Write([]byte(input))
	// execStdin.Close()

	stdOutOutput, err := io.ReadAll(execStdout)
	if err != nil {
		fmt.Println("Error di stdoutput", err)
	}
	stdErrOutput, err := io.ReadAll(execStderr)
	if err != nil {
		fmt.Println("Error di stderrput", err)
	}
	codeOutput := string(stdOutOutput)
	execCmd.Wait()

	if string(stdErrOutput) != "" {
		codeOutput = string(stdErrOutput)
	}

	return codeOutput, mainErr
}
