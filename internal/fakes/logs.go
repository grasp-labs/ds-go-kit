package fakes

import (
	"bytes"
	"log"
)

func CaptureLogs(fn func()) string {
	var buf bytes.Buffer
	old := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(old)

	fn()
	return buf.String()
}
