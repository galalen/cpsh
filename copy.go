package cpsh

import (
	"os"

	"golang.design/x/clipboard"
)

func Init() {
	err := clipboard.Init()
	HandleError(err)
}

func CopyToClipboard(data string) bool {
	clipboard.Write(clipboard.FmtText, []byte(data))
	return true
}

func ReadFromFileAndCopy(path string) {
	data, err := os.ReadFile(path)
	HandleError(err)
	CopyToClipboard(string(data))
}
