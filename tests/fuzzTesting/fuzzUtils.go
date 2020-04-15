package fuzzTesting

import (
	"fmt"
	"github.com/amitdavidson234/go-uniq/pkg/utils"
	"math/rand"
	"os"
	"time"
)

func writeToFuzzFile(path string, data []byte) {
	// write fuzzing data to a file
	f, err := utils.CreateFile(path)
	if err != nil {
		fmt.Println("Couldn't create file")
		os.Exit(-1)
	}
	err = utils.WriteFile(f, data)
	if err != nil {
		fmt.Println("Couldn't write to file")
		os.Exit(-1)
	}
	f.Close()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
