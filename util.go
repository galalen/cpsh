package cpsh

import (
	"fmt"
	"os"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
	}
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(Red, err)
		os.Exit(1)
	}
}

func Println(msg string) {
	fmt.Println(Green, msg+Reset)
}
