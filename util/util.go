package util

import (
	"io/ioutil"
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

func LoadMusic(path string) ([]byte, error) {
	if BuildTime == "true" {

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			log.Fatalf("Asset Error: %v\n", err)
		}

		path = dir + path
		path = strings.ReplaceAll(path, "./", "/")
	}

	songBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return songBytes, nil
}

// Smoother Cam
func CamLerp(a float64, b float64) float64 {
	t := 0.125

	return a*(1-t) + b*t
}

