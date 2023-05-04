package entity

import (
	"os/exec"
	"strings"
	"syscall"
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

	attr := &syscall.SysProcAttr{
		HideWindow: true,
	}
	cmd := exec.Command("cmd", "/c", cm)
	cmd.SysProcAttr = attr

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return out
}

func (c *Command) MustExecPowershell(cm string) []byte {
	attr := &syscall.SysProcAttr{
		HideWindow: true,
	}

	cmd := exec.Command("powershell", cm)
	cmd.SysProcAttr = attr

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return out
}
