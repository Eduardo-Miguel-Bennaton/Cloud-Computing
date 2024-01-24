package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"main/processes"
	"strings"
)

func main() {
	directory := `..\mockfiles`

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal("Error reading directory:", err)
		return
	}

	for _, fileInfo := range files {
		fileName := fileInfo.Name()

		newFilePathA := filepath.Join(directory, "Type A", fileName)
		newFilePathB := filepath.Join(directory, "Type B", fileName)
		newFilePathC := filepath.Join(directory, "Type C", fileName)

		if strings.HasPrefix(fileName, "TypeA") {
			fmt.Printf("%s belongs to Type A\n", fileName)
			os.MkdirAll(filepath.Join(directory, "Type A"), os.ModePerm)
			err := os.Rename(filepath.Join(directory, fileName), newFilePathA)
			if err != nil {
				fmt.Printf("Error moving file %s: %v\n", fileName, err)
			}
		} else if strings.HasPrefix(fileName, "TypeB") {
			fmt.Printf("%s belongs to Type B\n", fileName)
			os.MkdirAll(filepath.Join(directory, "Type B"), os.ModePerm)
			err := os.Rename(filepath.Join(directory, fileName), newFilePathB)
			if err != nil {
				fmt.Printf("Error moving file %s: %v\n", fileName, err)
			}
		} else if strings.HasPrefix(fileName, "TypeC") {
			fmt.Printf("%s belongs to Type C\n", fileName)
			os.MkdirAll(filepath.Join(directory, "Type C"), os.ModePerm)
			err := os.Rename(filepath.Join(directory, fileName), newFilePathC)
			if err != nil {
				fmt.Printf("Error moving file %s: %v\n", fileName, err)
			}
		} else {
			fmt.Printf("%s does not match any known type\n", fileName)
		}

		if strings.HasPrefix(fileName, "Type A") {
			subDirectory := filepath.Join(directory, "Type A")
			files, err := ioutil.ReadDir(subDirectory)
			if err != nil {
				log.Fatal("Error reading directory:", err)
				return
			}

			for _, fileInfo := range files {
				filePath := filepath.Join(subDirectory, fileInfo.Name())

				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
					continue
				}
				processes.ProcessA(fileData)
			}
		} else if strings.HasPrefix(fileName, "Type B") {
			subDirectory := filepath.Join(directory, "Type B")
			files, err := ioutil.ReadDir(subDirectory)
			if err != nil {
				log.Fatal("Error reading directory:", err)
				return
			}

			for _, fileInfo := range files {
				filePath := filepath.Join(subDirectory, fileInfo.Name())

				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
					continue
				}
				processes.ProcessB(fileData)
			}
		} else if strings.HasPrefix(fileName, "Type C") {
			subDirectory := filepath.Join(directory, "Type C")
			files, err := ioutil.ReadDir(subDirectory)
			if err != nil {
				log.Fatal("Error reading directory:", err)
				return
			}

			for _, fileInfo := range files {
				filePath := filepath.Join(subDirectory, fileInfo.Name())

				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
					continue
				}
				processes.ProcessC(fileData)
			}
		}
	}
}
