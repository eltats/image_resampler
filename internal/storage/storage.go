package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

func CheckCache(content []byte, width, height int, resDir string) bool {
	hash := generateCacheFilePath(content, width, height)
	resPath := filepath.Join(resDir, hash)
	_, err := os.Stat(resPath)
	if os.IsNotExist(err) {
		return true
	}
	return false
}
func generateCacheFilePath(content []byte, width, height int) string {
	hash := sha256.Sum256(content)
	return fmt.Sprintf("%s_%dx%d.jpg", hex.EncodeToString(hash[:]), width, height)
}
func SaveImages(processedImg image.Image, origImg []byte, resDir, origDir string, width, height int) error {
	cacheName := generateCacheFilePath(origImg, width, height)
	resFilename := filepath.Join(resDir, cacheName)
	outFile, err := os.Create(resFilename)
	if err != nil {
		return errors.New("failed to create processed image file")
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, processedImg, &jpeg.Options{Quality: 80}); err != nil {
		return errors.New("failed to save processed image")
	}
	origFilename := filepath.Join(origDir, cacheName)
	file, err := os.Create(origFilename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(origImg)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
