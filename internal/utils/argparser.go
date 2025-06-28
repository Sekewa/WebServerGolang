package utils

import (
	"log"
	"strconv"
)

/*
This function will return a map[string]string for all parameters possible
*/
func ArgsParser(args []string) map[string]string {
	lenArgs := len(args)

	mapParam := make(map[string]string)
	// in case there is no port the user specified
	mapParam["port"] = ":8080"

	for i := 0; i < lenArgs; i++ {

		switch s := args[i]; s {
		case "--p", "-port":

			if i+1 <= lenArgs-1 {
				// we get the second args
				portTemp := args[i+1]

				// conversion string -> int
				_, err := strconv.Atoi(portTemp)

				// we check if the second args is not a existing params
				if contains(params, portTemp) {
					log.Println("autre paramètre", portTemp)
					mapParam["port"] = ":8080"
					break
				}

				// we check if the port is parsable
				if err != nil {
					log.Println("Impossible de trouver un port utilisable, veuillez le rentrer dans ce format : xxxx, où x est un chiffre")
					mapParam["port"] = ":8080"
					break
				}

				// arrived there we can +1 because the port is good
				i += 1
				mapParam["port"] = ":" + portTemp
				//log.Printf("Port : %s\n", mapParam["port"])
			} else {
				log.Println("il manque le port, par défaut 8080 sera utilisé")
			}
		}

	}

	return mapParam
}

func contains(haystack []string, needle string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			log.Println(haystack[i], "==", needle)
			return true
		}
	}
	return false
}
