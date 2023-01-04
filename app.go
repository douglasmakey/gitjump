package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/muesli/coral"
)

var (
	to, branch string
)

func runGitCommand(out io.Writer, params ...string) error {
	c := exec.Command("git", params...)
	c.Stdout = out
	return c.Run()
}

func goToNth(n int) error {
	var buffer bytes.Buffer
	if err := runGitCommand(os.Stdout, "checkout", branch); err != nil {
		return fmt.Errorf("git checkout failed: %w", err)
	}

	if err := runGitCommand(&buffer, "rev-list", "--count", branch); err != nil {
		return err
	}

	commitsQty, err := strconv.Atoi(string(buffer.Bytes()[0 : buffer.Len()-1]))
	if err != nil {
		return err
	}

	buffer.Reset()
	if err := runGitCommand(&buffer, "log", "--reverse", "--oneline"); err != nil {
		return err
	}

	r := bufio.NewReader(&buffer)
	commits := make([]string, 0, commitsQty)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		split := strings.Split(string(line), " ")
		commits = append(commits, split[0])
	}

	if n < 0 {
		n = (commitsQty - 1) - int(math.Abs(float64(n)))
	}

	return runGitCommand(os.Stdout, "checkout", commits[n])
}

func run(to string) error {
	switch to {
	case "last":
		if err := runGitCommand(os.Stdout, "checkout", "main"); err != nil {
			return fmt.Errorf("git checkout failed: %w", err)
		}

		return nil
	default:
		n, err := strconv.Atoi(to)
		if err != nil {
			return fmt.Errorf("error parsing: %w", err)
		}

		return goToNth(n)
	}
}

func root() *coral.Command {
	root := &coral.Command{
		Use:           "gitjump",
		Short:         "",
		SilenceErrors: true,
		SilenceUsage:  false,
		RunE: func(cmd *coral.Command, args []string) error {
			return run(to)
		},
	}

	root.Flags().StringVar(&to, "goto", "0", "to go to specific commit")
	root.Flags().StringVar(&branch, "branch", "main", "from specific branch")
	return root
}
