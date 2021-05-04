package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spencer-p/geneticdj/pkg/sexpr"
)

const defaultTimeout = 10 * time.Second

func play(ctx context.Context, fn sexpr.Fn) {
	cmd := exec.CommandContext(ctx, "/usr/bin/pacat",
		"--rate", "8000",
		"--format", "u8")

	// Retrieve pipe info before any forking.
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

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s '<s-expression>'\n", os.Args[0])
		return
	}

	fn, err := sexpr.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("invalid expression: %v\n", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), defaultTimeout)
	play(ctx, fn)
}
