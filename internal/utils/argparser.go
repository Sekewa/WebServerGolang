package utils

import (
	"log"
	"strconv"
)

func ArgsParser(args []string) map[string]string {
	lenArgs := len(args)

	mapParam := make(map[string]string)

	mapParam["port"] = ":8080"

	for i := 0; i < lenArgs; i++ {

		switch s := args[i]; s {
		case "--p", "-port":

			if i+1 <= lenArgs-1 {
				// on recupere l'index d'apres pour avoir le parametres
				portTemp := args[i+1]

				// conversion string -> int
				_, err := strconv.Atoi(portTemp)

				// on regarde si ce n'est pas un autre chose comme un des parametres
				if contains(params, portTemp) {
					log.Println("autre paramètre", portTemp)
					mapParam["port"] = ":8080"
					break
				}

				// on regarde si il n'y aurait pas un probleme avec le port
				if err != nil {
					log.Println("Impossible de trouver un port utilisable, veuillez le rentrer dans ce format : xxxx, où x est un chiffre")
					mapParam["port"] = ":8080"
					break
				}

				// on avance la boucle d'1 sachant qu'on a traite correctement les informations
				i += 1
				mapParam["port"] = ":" + portTemp
				log.Printf("Port : %s\n", mapParam["port"])
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
