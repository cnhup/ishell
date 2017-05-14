package ishell_test

import (
	"testing"

	"github.com/abiosoft/ishell"
	"github.com/stretchr/testify/assert"
)

func newCmd(name string, help string) *ishell.Cmd {
	return &ishell.Cmd{
		Name: name,
		Help: help,
	}
}
func TestAddCommand(t *testing.T) {
	cmd := newCmd("root", "")
	assert.Equal(t, len(cmd.Children()), 0, "should be empty")
	cmd.AddCmd(newCmd("child", ""))
	assert.Equal(t, len(cmd.Children()), 1, "should include one child command")
}

func TestDeleteCommand(t *testing.T) {
	cmd := newCmd("root", "")
	cmd.AddCmd(newCmd("child", ""))
	assert.Equal(t, len(cmd.Children()), 1, "should include one child command")
	cmd.DeleteCmd("child")
	assert.Equal(t, len(cmd.Children()), 0, "should be empty")
}

func TestFindCmd(t *testing.T) {
	cmd := newCmd("root", "")
	cmd.AddCmd(newCmd("child1", ""))
	cmd.AddCmd(newCmd("child2", ""))
	res, err := cmd.FindCmd([]string{"child1"})
	if err != nil {
		t.Error("finding should work")
	}
	assert.Equal(t, res.Name, "child1")

	res, err = cmd.FindCmd([]string{"child2"})
	if err != nil {
		t.Error("finding should work")
	}
	assert.Equal(t, res.Name, "child2")

	res, err = cmd.FindCmd([]string{"child3"})
	if err == nil {
		t.Error("should not find this child!")
	}
	assert.Nil(t, res)
}

func TestHelpText(t *testing.T) {
	cmd := newCmd("root", "help for root command")
	cmd.AddCmd(newCmd("child1", "help for child1 command"))
	cmd.AddCmd(newCmd("child2", "help for child2 command"))
	res := cmd.HelpText()
	expected := "\nhelp for root command\n\nCommands:\n  child1      help for child1 command\n  child2      help for child2 command\n\n"
	assert.Equal(t, res, expected)
}
