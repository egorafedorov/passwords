package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		color.Red("Error!")
		return
	}
	fmt.Println(string(data))
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
