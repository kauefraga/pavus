package lib

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
)

// Verify if a file is an asset (.png, .jpg/.jpeg, .gif, .webp, .avif, .mp4)
func isAsset(filename string) bool {
	isAsset, err := regexp.MatchString(`^.*\.(png|jpg|jpeg|gif|webp|avif|mp4)$`, filename)
	if err != nil {
		color.Red("Error: %s", err)
		os.Exit(1)
	}

	return isAsset
}

func findAssetDirectory(root string) string {
	var assetDirectory string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if isAsset(info.Name()) {
			assetDirectory = filepath.Dir(path)
			return filepath.SkipAll
		}

		return err
	})
	if err != nil {
		color.Red("Error: %s", err)
		os.Exit(1)
	}

	return assetDirectory
}

// If the asset flag is present, return it,
// otherwise find and return the first asset directory
func GetAssetDirectory(assetPath string) string {
	if assetPath == "" {
		return findAssetDirectory(".")
	}

	_, err := os.Stat(assetPath)
	if err != nil {
		return findAssetDirectory(".")
	}

	// trust user input 💀
	return assetPath
}
