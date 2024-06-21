package logger

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"time"
	"fmt"
)

var errorPath string 
var systemPath string 

func Start() {

	var homeDir string = createDir()

    // rootDir := filepath.Join(homeDir, "webdav-server", "log")

	errorPath = filepath.Join(homeDir, "error_log.txt")
	systemPath = filepath.Join(homeDir, "system_log.txt")


    dir := createDir()

	var wg sync.WaitGroup

    // Add 2 to the WaitGroup for two goroutines
    wg.Add(3)

    // Start goroutines to create files
    go createFile(filepath.Join(dir, "error_log.txt"), &wg)
    go createFile(filepath.Join(dir, "system_log.txt"), &wg)

    // Wait for all goroutines to finish
    wg.Wait()

}

func createDir() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(currentUser.HomeDir, "webdav-server", "log")

	// Check if the directory already exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Directory does not exist, create it
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		// Some error other than directory not exists
		log.Fatal(err)
	} else {
		// Directory already exists, no need to create
		fmt.Printf("Directory %s already exists\n", dir)

	}

	return dir
}

func createFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Check if the file already exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Error creating file %s: %s", filePath, err)
		}
		defer file.Close()

		log.Printf("Created file: %s", filePath)
	} else if err != nil {
		// Error occurred while checking if file exists
		log.Fatalf("Error checking if file %s exists: %s", filePath, err)
	} else {
		// Clear File of pervious data 
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil { 
			log.Fatal(err)
		}

		defer file.Close()

		// File already exists
		log.Printf("File %s already exists, clearing data", filePath)

		
	}
}


func LogError(e error) {
	// Open the file for appending or create it if it doesn't exist
	file, err := os.OpenFile(errorPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error opening file", errorPath, ":", err)
		return
	}
	defer file.Close()

	// Write error message to the file
	_, err = file.WriteString(e.Error() + "\n")
	if err != nil {
		log.Println("Error writing to file", errorPath, ":", err)
	}
}

func LogSystem(data string) {
	// Open the file for appending or create it if it doesn't exist
	file, err := os.OpenFile(systemPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error opening file", systemPath, ":", err)
		return
	}
	defer file.Close()

	// Prepare data to write, including current date and time
	currentDateTime := time.Now().Format("2006-01-02 15:04:05")
	writeData := currentDateTime + " " + data + "\n"

	// Write system log message to the file
	_, err = file.WriteString(writeData)
	if err != nil {
		log.Println("Error writing to file", systemPath, ":", err)
	}
}

