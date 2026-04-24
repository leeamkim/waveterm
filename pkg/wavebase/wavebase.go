// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// Package wavebase provides core utilities and base configuration for WaveTerm.
package wavebase

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const WaveTermAppName = "waveterm"
const WaveTermVersion = "v0.1.0"
const WaveHomeDirName = ".waveterm"
const WaveDBDirName = "db"
const WaveLogDirName = "log"

var globalLock sync.Mutex
var waveHomeDir string

// GetHomeDir returns the user's home directory.
func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "/"
	}
	return homeDir
}

// GetWaveHomeDir returns the WaveTerm home directory (~/.waveterm by default).
// The directory is created if it does not exist.
func GetWaveHomeDir() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	if waveHomeDir != "" {
		return waveHomeDir
	}
	if override := os.Getenv("WAVETERM_HOME"); override != "" {
		waveHomeDir = override
	} else {
		waveHomeDir = filepath.Join(GetHomeDir(), WaveHomeDirName)
	}
	return waveHomeDir
}

// EnsureWaveHomeDir creates the WaveTerm home directory and required subdirectories.
func EnsureWaveHomeDir() error {
	homeDir := GetWaveHomeDir()
	subDirs := []string{
		homeDir,
		filepath.Join(homeDir, WaveDBDirName),
		filepath.Join(homeDir, WaveLogDirName),
	}
	for _, dir := range subDirs {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return fmt.Errorf("failed to create directory %q: %w", dir, err)
		}
	}
	return nil
}

// GetWaveDBDir returns the path to the WaveTerm database directory.
func GetWaveDBDir() string {
	return filepath.Join(GetWaveHomeDir(), WaveDBDirName)
}

// GetWaveLogDir returns the path to the WaveTerm log directory.
func GetWaveLogDir() string {
	return filepath.Join(GetWaveHomeDir(), WaveLogDirName)
}

// GetOSInfo returns a human-readable string describing the current OS and architecture.
func GetOSInfo() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

// IsValidHomeDir checks whether the given path exists and is a directory.
func IsValidHomeDir(path string) error {
	if path == "" {
		return errors.New("home directory path is empty")
	}
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("home directory does not exist: %q", path)
		}
		return fmt.Errorf("cannot stat home directory %q: %w", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("home directory path is not a directory: %q", path)
	}
	return nil
}
