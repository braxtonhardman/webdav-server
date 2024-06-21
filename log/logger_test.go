package logger

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
	"log"
	"sync"
)

func TestStart(t *testing.T) {

    t.Cleanup(func() {
        user, err := user.Current()
        if err != nil {
            t.Fatalf("Error in getting current user: %s", err)
        }

        dir := filepath.Join(user.HomeDir, "CloudServer", "log")
        path1 := filepath.Join(dir, "error_log.txt")
        path2 := filepath.Join(dir, "system_log.txt")

        var wg sync.WaitGroup
        wg.Add(2) // We have two goroutines for file removal

        go func() {
            defer wg.Done()
            // Perform cleanup tasks for path1
            if err := os.Remove(path1); err != nil {
                log.Printf("Error removing file %s: %s", path1, err)
            }
        }()

        go func() {
            defer wg.Done()
            // Perform cleanup tasks for path2
            if err := os.Remove(path2); err != nil {
                log.Printf("Error removing file %s: %s", path2, err)
            }
        }()

        // Wait for all cleanup operations to finish before removing directories
        wg.Wait()

        // Now remove the directories themselves
        if err := os.Remove(dir); err != nil {
            log.Printf("Error removing directory %s: %s", dir, err)
        }

        if err := os.Remove(filepath.Join(user.HomeDir, "CloudServer")); err != nil {
            log.Printf("Error removing directory %s: %s", filepath.Join(user.HomeDir, "CloudServer"), err)
        }
    })

    done := make(chan struct{})
    defer close(done)

    // Call Start in a goroutine
    go func() {
        defer close(done)
        Start()
    }()

    // Wait for Start() to complete
    <-done

    // Add assertions or additional checks if needed

    // Add assertions if needed
}
