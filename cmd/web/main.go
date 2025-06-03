package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/Sekewa/WebServerGolang/internal/utils"
)

var (
	port   string = ":8080"
	params        = make([]string, 2)
)

/*
Type de base pour tester
*/
type countHandler struct {
	mu sync.Mutex
	n  int
}

/*
Exemple de base sur le site de golang
*/
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

/*
Exemple supplementaire
*/
func Exit(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "BYE !!")
	log.Fatal("BYE !!")
}

func PopulateParams() {
	params[0] = "--p"
	params[1] = "-port"
}

func ArgsParser() {
	lenArgs := len(os.Args)

	for i := 0; i < lenArgs; i++ {

		switch s := os.Args[i]; s {
		case "--p", "-port":

			if i+1 <= lenArgs-1 {
				// on recupere l'index d'apres pour avoir le parametres
				portTemp := os.Args[i+1]

				// conversion string -> int
				_, err := strconv.Atoi(portTemp)

				// on regarde si ce n'est pas un autre chose comme un des parametres
				if slices.Contains(params, portTemp) {
					log.Println("autre paramètre")
					break
				}

				// on regarde si il n'y aurait pas un probleme avec le port
				if err != nil {
					log.Println("Impossible de trouver un port utilisable, veuillez le rentrer dans ce format : xxxx, où x est un chiffre")

				}

				// on avance la boucle d'1 sachant qu'on a traite correctement les informations
				i += 1
				port = ":" + portTemp
				log.Printf("Port : %s\n", port)
			} else {
				log.Println("il manque le port, par défaut 8080 sera utilisé")
			}
		}

	}
}

func main() {
	PopulateParams()
	ArgsParser()

	logger := utils.NewLogger()

	mux := http.NewServeMux()

	mux.Handle("/count", new(countHandler))
	mux.HandleFunc("/exit", Exit)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	loggedMux := utils.LoggingMiddleware(mux, logger)

	srv := &http.Server{
		Addr:         port,
		Handler:      loggedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Start of serveur on http://localhost", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server stopped : %v", err)
	}
}
