package main

import (
	"context"
	"github.com/spencer-p/geneticdj/sexpr"
	"io"
	"os/exec"
)

func readAll(file *io.Reader) []sexpr.Fn {
	return nil
}

func play(ctx context.Context, fn sexpr.Fn) {
	cmd := exec.CommandContext(ctx, "/usr/bin/pacat",
		"--rate", "8000",
		"--format", "u8")

	// Retrieve pipe info before starting. Weird things happen otherwise.
	pipe, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	// Start the player
	cmd.Start()

	// procResult will send the err (or nil) of the process when it is done
	procResult := make(chan error)
	go func() {
		procResult <- cmd.Wait()
	}()

	// Write bytes to the process
	go func() {
		defer pipe.Close()
		for t := 0; ; t++ {
			select {
			case <-ctx.Done():
				return
			default:
				pipe.Write([]byte{byte(fn.Eval(t))})
			}
		}
	}()

	if err := <-procResult; err != nil {
		if err.Error() != "signal: killed" {
			panic(err)
		}
	}
}
