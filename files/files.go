package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		color.Red("Error!")
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		color.Red("Error!")
		return
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		color.Red("Error!")
		return
	}
	color.Green("File writing successful!")
}
