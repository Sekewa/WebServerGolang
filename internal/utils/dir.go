package utils

import (
	"log"
	"net/http"
	"os"
	"runtime"
)

/*
This function will return the directory to use on the root path.
If there is no static directory in this project it will just take the example on to return.
If none a present it will stop the Web Server.
*/
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
