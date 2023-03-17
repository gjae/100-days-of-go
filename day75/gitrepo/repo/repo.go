package repo

import (
	"bytes"
	"context"
	"io"
	"log"
	"os/exec"
	"time"
)

var (
	DefaultCommandTimeout = 60 * time.Second
)

// Command represents a Git command minimun instance
type Command struct {
	name string
	args []string
}

// NewCommand creates a new command object
func NewCommand(args ...string) *Command {
	return &Command{
		name: "git",
		args: args,
	}
}

// AddArguments adds a new arguments and params
// to command
func (c *Command) AddArguments(args ...string) *Command {
	c.args = append(c.args, args...)

	return c
}

func (c *Command) RunInDirPipeplineTimeout(dir string, timeout int, stdout, stderr io.Writer) error {

	var commandTimeout time.Duration

	if timeout == -1 {
		commandTimeout = DefaultCommandTimeout
	} else {
		commandTimeout = time.Second * time.Duration(timeout)
	}

	if len(dir) > 0 {
		log.Printf("To repo: %s", dir)
	} else {
		log.Printf("%s %v", dir, c)
	}
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.name, c.args...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return ctx.Err()

}

// Runs the command and get the output command linking stdut to
// a buffer
func (c *Command) Run(repoPath string, timeout int) string {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	err := c.RunInDirPipeplineTimeout(repoPath, timeout, stdout, stderr)

	if err != nil {
		panic(err)
	}

	return stdout.String()
}
