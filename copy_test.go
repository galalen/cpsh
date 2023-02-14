package cpsh

import (
	"testing"

	"golang.design/x/clipboard"
)

func TestCopyToClipboard(t *testing.T) {
	data := "data"
	ok := CopyToClipboard(data)
	if !ok {
		t.Error("data should be copied to clipboard")
	}

	msg := string(clipboard.Read(clipboard.FmtText))
	if msg != data {
		t.Errorf("Excepted %s, found %s", data, msg)
	}
}
