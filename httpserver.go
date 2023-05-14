package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
    "os"

    "github.com/gorilla/mux"
    "github.com/dotenv-org/godotenvvault"
)

type systemData struct{
    API_KEY string
    SYSTEM string
}

func main() {
    err := godotenvvault.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

    r := mux.NewRouter()

    homerouter := r.PathPrefix("").Subrouter()
    homerouter.HandleFunc("/", homePage).Methods("GET")

    systemrouter := r.PathPrefix("/systems").Subrouter()
    systemrouter.HandleFunc("/", systemViewer).Methods("GET")

    textures := http.StripPrefix("/textures/", http.FileServer(http.Dir("./textures/")))
    r.PathPrefix("/textures/").Handler(textures)
    
    fmt.Println("http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        log.Fatal("Error loading index.html")
    }
    tmpl.Execute(w, nil)
}

func systemViewer(w http.ResponseWriter, r *http.Request) {
    system := r.URL.Query().Get("system")

    systemData := systemData{
    API_KEY: os.Getenv("API_KEY"),
    SYSTEM: system}

    tmpl, err := template.ParseFiles("systemViewer.html")
    if err != nil {
        log.Fatal("Error loading systemViewer.html")
    }
    tmpl.Execute(w, systemData)
}
