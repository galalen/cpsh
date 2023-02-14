package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/galalen/cpsh"
)

func main() {

	filePath := flag.String("f", "", "file path to copy its content")
	content := flag.String("c", "", "direct content to copy")
	flag.Parse()

	switch {
	case *filePath != "":
		cpsh.ReadFromFileAndCopy(*filePath)
	case *content != "":
		cpsh.CopyToClipboard(*content)
	default:
		// check for pipe data
		fi, err := os.Stdin.Stat()
		cpsh.HandleError(err)

		if fi.Mode()&os.ModeNamedPipe != 0 {
			var buf []byte
			scanner := bufio.NewScanner(os.Stdin)

			for scanner.Scan() {
				buf = append(buf, scanner.Bytes()...)
			}
			if err := scanner.Err(); err != nil {
				cpsh.HandleError(err)
			}

			cpsh.CopyToClipboard(string(buf))
		}
	}
	cpsh.Println("copied!")

}
