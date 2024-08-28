package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Define the src directory
	srcDir := filepath.Join(cwd, "src")

	// Walk through the src directory
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if it's not a directory or if it's the top-level src directory
		if !info.IsDir() || path == srcDir {
			return nil
		}

		// Create index file for the current directory
		createIndexFile(path)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
}

func createIndexFile(dir string) {
	indexPath := filepath.Join(dir, "index.ts")
	file, err := os.OpenFile(indexPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating index file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	existingExports := make(map[string]bool)

	for _, f := range files {
		if f.IsDir() || f.Name() == "index.ts" {
			continue
		}

		if strings.HasSuffix(f.Name(), ".tsx") || strings.HasSuffix(f.Name(), ".jsx") || strings.HasSuffix(f.Name(), ".ts") || strings.HasSuffix(f.Name(), ".js") {
			filePath := filepath.Join(dir, f.Name())
			if hasExports(filePath) {
				baseName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				exportLine := fmt.Sprintf("export * from './%s';\n", baseName)

				if !existingExports[exportLine] {
					_, err := writer.WriteString(exportLine)
					if err != nil {
						fmt.Println("Error writing to index file:", err)
					}
					existingExports[exportLine] = true
				}
			}
		}
	}

	writer.Flush()
}

func hasExports(filePath string) bool {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	// Regular expression to match export statements
	exportRegex := regexp.MustCompile(`(?m)^export\s+`)
	return exportRegex.Match(content)
}
