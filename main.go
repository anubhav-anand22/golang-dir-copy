package main

import (
	"fmt"
	"io"
	"main/lib"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ask(query string) (string, error) {
	var input string

	fmt.Println(query)
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}

func isExistsAndIsFolder(path string) bool {
	if r, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Error: Folder does not exist.")
	} else {
		if r.IsDir() {
			return true
		} else {
			fmt.Println("Error: Not a folder")
		}
	}
	return false
}

func main() {
	targetFolder, err := ask("Target folder (folder containing new files)")

	if err != nil {
		fmt.Println("Error: Unable to get target folder!")
		return
	}

	destFolder, err := ask("Destination folder (folder where files will be copied)")

	if err != nil {
		fmt.Println("Error: Unable to get target folder!")
		return
	}

	println("Target folder:- " + targetFolder)
	println("Destination folder:- " + destFolder)

	isTargetFolderExists := isExistsAndIsFolder(targetFolder)
	isDestFolderExists := isExistsAndIsFolder(destFolder)

	if !isDestFolderExists || !isTargetFolderExists {
		fmt.Println("Error: Folder does not exist. Exiting...")
		os.Exit(1)
		return
	}

	targetFolderItems, err := lib.ListFiles(targetFolder)
	if err != nil {
		fmt.Println("Error: unable to list target folders")
		os.Exit(1)
		return
	}
	destFolderItems, err := lib.ListFiles(destFolder)
	if err != nil {
		fmt.Println("Error: unable to list destination folders")
		os.Exit(1)
		return
	}

	var MissingFilesInDest []string

	for _, tItem := range targetFolderItems {
		exit := false
		for _, dItem := range destFolderItems {
			tFS := strings.Replace(tItem, targetFolder, "", 1)
			dFS := strings.Replace(dItem, destFolder, "", 1)
			// println(tFS + "  ---  " + dFS)
			if tFS == dFS {
				exit = true
			}
		}
		if !exit {
			MissingFilesInDest = append(MissingFilesInDest, tItem)
		}
	}

	fmt.Println("")
	fmt.Println(strconv.Itoa(len(MissingFilesInDest)) + " Items are missing from destination folder")
	fmt.Println("")

	toContinue, err := ask("Do you want to continue? (y/N)")
	if err != nil {
		fmt.Println("Error: getting answer", err)
		os.Exit(1)
		return
	}
	if strings.TrimSpace(strings.ToLower(toContinue)) != "y" {
		fmt.Println("Your answer was 'No'. Exiting...")
		os.Exit(1)
		return
	}

	var errFilePath []string

	for _, items := range MissingFilesInDest {
		dPath := filepath.Join(destFolder, strings.Replace(items, targetFolder, "", 1))
		dPathDirPath := filepath.Dir(dPath)

		if _, err := os.Stat(dPathDirPath); os.IsNotExist(err) {
			// Directory does not exist, create it
			err := os.MkdirAll(dPathDirPath, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				errFilePath = append(errFilePath, items)
				continue
			}
		} else if err != nil {
			fmt.Println("Error checking directory existence:", err)
			errFilePath = append(errFilePath, items)
			continue
		}

		err = copyFile(items, dPath)
		if err != nil {
			errFilePath = append(errFilePath, items)
			continue
		}
	}

	fmt.Println("")
	fmt.Println(strconv.Itoa(len(errFilePath)) + " Items were not copied due to error")
	fmt.Println("")
}

func copyFile(sourcePath string, destPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return err
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(destPath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return err
	}
	defer destination.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return err
	}

	return nil
}
