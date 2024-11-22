package config

import (
	"flag"
	"os"
)

type Config struct {
	OrigDir string
	ResDir  string
	Width   uint
	Height  uint
}

func ParseFlags() *Config {
	cfg := &Config{}

	// Define CLI arguments
	flag.StringVar(&cfg.OrigDir, "path-orig", "/tmp/img_orig", "Directory for original images")
	flag.StringVar(&cfg.ResDir, "path-res", "/tmp/img_res", "Directory for processed images")
	flag.UintVar(&cfg.Width, "width", 200, "Resampling width")
	flag.UintVar(&cfg.Height, "height", 200, "Resampling height")

	flag.Parse()

	os.MkdirAll(cfg.OrigDir, os.ModePerm)
	os.MkdirAll(cfg.ResDir, os.ModePerm)

	return cfg
}
