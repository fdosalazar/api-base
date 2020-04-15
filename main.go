package main

import (
	"api-base/app/bundles/apibasebundle"
	"api-base/app/core"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Aqui se declaran los controladores
	initLog()

	r := mux.NewRouter()

	// Aqui se define el nombre del Api
	s := r.PathPrefix("/api-base/").Subrouter()

	// Aqui se definen las rutas de los metodos y si son GET o POST
	for _, b := range initBundles(nil) {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	//Aqui se define el puerto de la Api
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadConfig() *core.Config {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found")
	}
	c := &core.Config{}
	c.Fetch()
	return c
}

func initBundles(db *gorm.DB) []core.Bundle {
	return []core.Bundle{apibasebundle.NewApiBaseBundle(db)}
}

func initLog() {
	f, _ := os.OpenFile(os.Getenv("PATH_LOG")+os.Args[0]+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)
}
