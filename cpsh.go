package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

const Version = "0.1.1"

func main() {
	filePath := flag.String("f", "", "file path to copy its content")
	showVersion := flag.Bool("v", false, "show version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("cpsh version: %s\n", Version)
		return
	}

	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Failed to initialize clipboard: %v\n", err)
		os.Exit(1)
	}

	// Handle direct content from args
	if len(flag.Args()) > 0 {
		content := flag.Args()[0]
		clipboard.Write(clipboard.FmtText, []byte(content))
		printSuccess()
		return
	}

	// Handle file path
	if *filePath != "" {
		ext := strings.ToLower(filepath.Ext(*filePath))
		if ext == ".txt" || ext == "" {
			data, err := os.ReadFile(*filePath)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}
			clipboard.Write(clipboard.FmtText, data)
		} else {
			absPath, err := filepath.Abs(*filePath)
			if err != nil {
				fmt.Printf("Error getting absolute path: %v\n", err)
				os.Exit(1)
			}
			clipboard.Write(clipboard.FmtText, []byte(absPath))
		}
		printSuccess()
		return
	}

	// Handle pipe input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		clipboard.Write(clipboard.FmtText, data)
		printSuccess()
		return
	}

	flag.Usage()
	os.Exit(1)
}

func printSuccess() {
	fmt.Print("\rCopied! ")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\r       \r")
}
