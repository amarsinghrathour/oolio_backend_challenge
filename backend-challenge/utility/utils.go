package utility

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ValidatePromoCode(code string) (bool, error) {
	if len(code) < 8 || len(code) > 10 {
		return false, fmt.Errorf("promo code must be between 8 and 10 characters")
	}

	files := []string{"couponbase1.gz", "couponbase2.gz", "couponbase3.gz"}
	matchCount := 0

	for _, file := range files {
		found, err := checkPromoInFile(file, code)
		if err != nil {
			return false, err
		}
		if found {
			matchCount++
		}
		if matchCount >= 2 {
			return true, nil // Short-circuit as soon as 2 matches are found
		}
	}

	return false, fmt.Errorf("promo code is not valid or not found in at least two files")
}

func checkPromoInFile(filename, code string) (bool, error) {
	filePath := filepath.Join("utility", filename)
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return false, fmt.Errorf("could not create gzip reader for file %s: %v", filename, err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), code) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error while scanning file %s: %v", filename, err)
	}

	return false, nil
}
