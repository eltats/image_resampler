package validation

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"strings"
)

const (
	MinSize = 1024
	MaxSize = 10 * 1024 * 1024
)

func getImageDimensions(base64Image string) (width, height int, err error) {
	if strings.HasPrefix(base64Image, "data:image/jpeg;base64,") {
		base64Image = strings.TrimPrefix(base64Image, "data:image/jpeg;base64,")
	}

	imgBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode base64: %v", err)
	}

	img, _, err := image.DecodeConfig(bytes.NewReader(imgBytes))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode image: %v", err)
	}

	return img.Width, img.Height, nil
}

func ValidateImagePayload(base64Image string) ([]byte, error) {
	if strings.HasPrefix(base64Image, "data:image/jpeg;base64,") {
		base64Image = strings.TrimPrefix(base64Image, "data:image/jpeg;base64,")
	}
	data, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 encoding: %w", err)
	}

	size := len(data)
	if size < MinSize {
		return nil, fmt.Errorf("image size too small: %d bytes", size)
	}
	if size > MaxSize {
		return nil, fmt.Errorf("image size too large: %d bytes", size)
	}

	reader := bytes.NewReader(data)
	img, format, err := image.Decode(reader)
	if err != nil || img == nil {
		return nil, fmt.Errorf("file is not a valid image: %w", err)
	}
	height, width, err := getImageDimensions(base64Image)
	if err != nil {
		return nil, err
	}
	if height == 0 && width == 0 {
		return nil, fmt.Errorf("file is 0x0")
	}
	if strings.ToLower(format) != "jpeg" || strings.ToLower(format) != "jpeg" {
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	return data, nil
}
