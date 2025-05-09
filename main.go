package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openFile, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer openFile.Close()

	buffer := make([]byte, 8)

	for {
		n, err := openFile.Read(buffer)

		if n > 0 {
			fmt.Printf("read: %s\n", string(buffer[:n]))
		}

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}
