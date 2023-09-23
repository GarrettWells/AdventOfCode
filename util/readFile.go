package util

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadFile(filename string) string {
	// Gets the path of the calling function
	_, dir, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	body, err := os.ReadFile(path.Join(path.Dir(dir), filename))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	// Trims a potential newline character added from a paste
	return strings.TrimSuffix(string(body), "\n")
}
