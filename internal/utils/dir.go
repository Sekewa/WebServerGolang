package utils

import (
	"log"
	"net/http"
	"os"
	"runtime"
)

func SearchStatic() http.Dir {
	proj, err := os.Getwd()

	if err != nil {
		log.Fatalln("Erreur dans la récupération du chemin")
	}

	dir := http.Dir(proj)

	_, err = dir.Open("static")

	if err != nil {
		log.Println("static not found, search for example")
	} else {
		if runtime.GOOS == "windows" {
			dir += "\\static"
		} else if runtime.GOOS == "linux" {
			dir += "/static"
		}

		return dir
	}

	_, err = dir.Open("example")

	if err != nil {
		log.Fatalln("example not found")
	}

	if runtime.GOOS == "windows" {
		dir += "\\example"
	} else if runtime.GOOS == "linux" {
		dir += "/example"
	}

	return dir
}
