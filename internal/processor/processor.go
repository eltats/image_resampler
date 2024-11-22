package processor

import (
	"bytes"
	"errors"
	"image"

	"time"

	"github.com/nfnt/resize"
)

func Process(content []byte, width, height uint) (image.Image, int64, bool, error) {
	startTime := time.Now()

	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, 0, false, errors.New("failed to decode image data")
	}
	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	processingTime := time.Since(startTime).Milliseconds()

	return resizedImg, processingTime, false, nil
}
