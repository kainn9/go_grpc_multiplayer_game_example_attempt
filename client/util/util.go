package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	BuildTime string
)

/*
	Helper for loading relative images
*/
func LoadImg(path string) *ebiten.Image {

	if BuildTime == "true" {

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			log.Fatalf("Asset Error: %v\n", err)
		}

		path = dir + path
		path = strings.ReplaceAll(path, "./", "/")
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
