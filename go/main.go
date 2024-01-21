package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/module/apmhttp/v2"
)

func main() {
	godotenv.Load()
	fmt.Println("ELASTIC_APM_SERVER_URL:", os.Getenv("ELASTIC_APM_SERVER_URL"))
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/log/{myvar}", logPage)
	router.HandleFunc("/error/{myvar}", errorPage)
	router.HandleFunc("/fatal/{myvar}", fatalPage)
	http.ListenAndServe(":" + os.Getenv("WEB_PORT"), apmhttp.Wrap(router))

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func logPage(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)
	fmt.Fprintf(w, fmt.Sprintf("Sending Log! Route: %v", url["myvar"]))
}
func errorPage(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)
	fmt.Fprintf(w, fmt.Sprintf("Sending Error! Route: %v", url["myvar"]))
	apm.CaptureError(r.Context(), errors.New("Error Route: " + url["myvar"])).Send()
}
func fatalPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sending Fatal!")
	url := mux.Vars(r)
	log.Fatal("Fatal: " + url["myvar"])
}
