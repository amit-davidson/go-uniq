package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	f, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return content, err
}

func OpenFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}


func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(f *os.File, text []byte) error {
	dataWriter := bufio.NewWriter(f)
	_, err := dataWriter.Write(text)
	if err != nil {
		return err
	}
	err = dataWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(fileName string) (*os.File, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func WriteToStdout(text []byte) {
	_, err := os.Stdout.Write(text)
	if err != nil {
		ExitWithError(err)
	}
}
