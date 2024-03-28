package lib

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func FileHash(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Create a new SHA256 hash
	hash := sha256.New()

	// Copy the file data into the hash
	_, err = io.Copy(hash, file)
	if err != nil {
		fmt.Println("Error copying file data:", err)
		return "", err
	}

	// Get the hashed value as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert byte slice to a hexadecimal string
	hashString := fmt.Sprintf("%x", hashBytes)

	// fmt.Println("SHA256 Hash:", hashString)
	return hashString, nil
}
