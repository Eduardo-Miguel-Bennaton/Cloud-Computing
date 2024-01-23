// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"main/processes"
// 	"path/filepath"
// )

// func main() {
// 	directory := `C:\Users\eduar\Desktop\Final Project\mockfiles\Type A`

// 	files, err := ioutil.ReadDir(directory)
// 	if err != nil {
// 		log.Fatal("Error reading directory:", err)
// 		return
// 	}

// 	for _, fileInfo := range files {
// 		filePath := filepath.Join(directory, fileInfo.Name())

// 		fileData, err := ioutil.ReadFile(filePath)
// 		if err != nil {
// 			fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
// 			continue
// 		}
// 		processes.ProcessA(fileData)
// 	}
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/processes"
	"path/filepath"
)

func main() {
	directory := `C:\Users\eduar\Desktop\Final Project\mockfiles`

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal("Error reading directory:", err)
		return
	}

	for _, fileInfo := range files {
		
		if (fileInfo.Name() == "Type A"){
			directory := `C:\Users\eduar\Desktop\Final Project\mockfiles\Type A`

			files, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatal("Error reading directory:", err)
				return
			}

			for _, fileInfo := range files {
				filePath := filepath.Join(directory, fileInfo.Name())

				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
					continue
				}
				processes.ProcessA(fileData)
			}			
		} else if (fileInfo.Name() == "Type B"){
			// directory := `C:\Users\eduar\Desktop\Final Project\mockfiles\Type B`

			// files, err := ioutil.ReadDir(directory)
			// if err != nil {
			// 	log.Fatal("Error reading directory:", err)
			// 	return
			// }

			// for _, fileInfo := range files {
			// 	filePath := filepath.Join(directory, fileInfo.Name())

			// 	fileData, err := ioutil.ReadFile(filePath)
			// 	if err != nil {
			// 		fmt.Printf("Error reading file %s: %v\n", fileInfo.Name(), err)
			// 		continue
			// 	}
			// 	processes.ProcessB(fileData)
			// }
			fmt.Println("Found!")			
		} else if (fileInfo.Name() == "Type C"){
			directory := `C:\Users\eduar\Desktop\Final Project\mockfiles\Type C`

			files, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatal("Error reading directory:", err)
				return
			}

			for _, fileInfo := range files {
				filePath := filepath.Join(directory, fileInfo.Name())

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