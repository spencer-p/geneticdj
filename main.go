package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spencer-p/geneticdj/sexpr"
)

const (
	// Some samples that sound good.
	RYG    = "(+ (& 127 (| (/ (* t (% (^ (>> t 12) (- (>> t 12) 2)) 11)) 4) (>> t 13))) (& (/ (* t (& (i [51 54 51 54 52 54 56 57] (& (>> t 13) 7)) 15)) 12) 128))"
	VIZNUT = "(| (& (* t 5) (>> t 7)) (& (* t 3) (>> t 10)))"
)

var (
	fGenerate   = flag.Bool("generate", false, "generate and play a function")
	fCustomFunc = flag.String("play", "", "your own function to play")
	fDuration   = flag.String("duration", "30s", "duration to play (30s, 1h, etc)")
	fWriteFile  = flag.String("o", "", "write the audio to a file")
)

func main() {
	flag.Parse()
	timeout, err := time.ParseDuration(*fDuration)

	var fn sexpr.Fn
	if *fGenerate || *fCustomFunc != "" {

		// Create the function desired
		switch {
		case *fGenerate:
			rand.Seed(time.Now().UnixNano())
			fn = sexpr.Preset{
				MaxDepth:      10,
				NumRange:      [2]int{1, 10},
				ArrayLenRange: [2]int{3, 7},
				OddsBranch:    0.75,
				OddsVar:       0.5,
			}.Generate()
		case *fCustomFunc != "":
			var err error // to avoid shadowing fn.
			fn, err = sexpr.Parse(*fCustomFunc)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	// Print it for the user to see
	fmt.Println(fn)

	if *fWriteFile != "" {
		file, err := os.Create(*fWriteFile)
		if err != nil {
			fmt.Println("Could not open", *fWriteFile, "for writing:", err)
			os.Exit(1)
		}
		defer file.Close()
		err = sexpr.WriteWav(fn, file, int(timeout.Seconds()))
		if err != nil {
			fmt.Println("Could not write audio:", err)
			os.Exit(1)
		}
	} else {
		// Set up context
		if err != nil {
			fmt.Println("Invalid duration:", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Play it
		play(ctx, fn)
	}
}
