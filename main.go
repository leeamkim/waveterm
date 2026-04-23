// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/wavetermdev/waveterm/pkg/wavebase"
)

var WaveVersion = "v0.0.1"
var BuildTime = "0"

func main() {
	wavebase.WaveVersion = WaveVersion
	wavebase.BuildTime = BuildTime

	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func runMain() error {
	// Print version info to stderr so it doesn't interfere with stdout output
	fmt.Fprintf(os.Stderr, "WaveTerm %s (build %s)\n", WaveVersion, BuildTime)
	return nil
}
