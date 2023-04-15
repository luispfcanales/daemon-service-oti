package entity

import (
	"os/exec"
	"strings"
)

type Command struct{}

func NewCommand() *Command {
	return &Command{}
}
func (c *Command) GetInfoCMD(command string) []string {

	outCommand := c.MustExecCmd(command)
	bodyValue := string(outCommand)
	rowValue := strings.SplitAfter(bodyValue, "\n")
	lineValue := strings.TrimSpace(rowValue[2])
	values := strings.Split(lineValue, ",")
	return values
}
func (c *Command) GetInfoPOWERSHELL(command string) string {
	line := c.MustExecPowershell(command)
	return string(line)
}

func (c *Command) MustExecCmd(cm string) []byte {
	cmd := exec.Command("cmd", "/c", cm)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return out
}

func (c *Command) MustExecPowershell(cm string) []byte {
	cmd := exec.Command("powershell", cm)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return out
}
