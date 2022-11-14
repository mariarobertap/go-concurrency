package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	var wg sync.WaitGroup
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	expectedString := "testando"
	wg.Add(1)

	go printSomething(expectedString, &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, expectedString) {
		t.Errorf("Expected = %s; Got = %s", expectedString, output)
	}
}
