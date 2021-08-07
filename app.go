package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	//"fmt"
	"github.com/gorilla/mux"
)

// App method is the main struct for the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init method is initialized the configs,routes,etc.
func (a *App) Init(database string) error {
	log.Println("Creating Router...")
	a.Router = mux.NewRouter()
	log.Println("Initializing Routes...")
	a.InitRoutes()
	if _, err := os.Stat(database); os.IsNotExist(err) {
		log.Println("Database can not be found.Creating new...")
		file, err := os.Create(database) // Create SQLite file
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		a.DB, err = sql.Open("sqlite3", database)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Created database file: " + database)

		query := "CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);"
		statement, err := a.DB.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Users table")

		query = "CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL DEFAULT 0,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER DEFAULT 0,Player2Number INTEGER DEFAULT 0,Predictor INTEGER DEFAULT 0,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0,Password TEXT);"
		statement, err = a.DB.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Sessions table")

		query = "CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT,Predictor INTEGER,Start INTEGER,End INTEGER,Winner INTEGER,Turn INTEGER);"
		statement, err = a.DB.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Moves table")

		enableFK := "PRAGMA foreign_keys=ON;"
		statement, err = a.DB.Prepare(enableFK)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Foreign Key enabled for sqlite.")
		log.Println("Created tables")
		return a.DB.Ping()
	} else {
		log.Println("Database file found.")
		var err error
		a.DB, err = sql.Open("sqlite3", database)
		if err != nil {
			log.Fatal(err)
		}
		return a.DB.Ping()
	}
}

// Run method runs the application with specified parameters
func (a *App) Run(addr string) {
	log.Println("Serving on: ", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// InitRoutes initializes the backend routes.
func (a *App) InitRoutes() {
	a.Router.HandleFunc("/backend/CreateGame", Inbound(a.CreateGameHandler))
	a.Router.HandleFunc("/backend/JoinGame", Inbound(a.JoinGameHandler))
	a.Router.HandleFunc("/backend/AbandonGame", Inbound(Auth(a.AbandonGameHandler)))
	a.Router.HandleFunc("/backend/GetReady", Inbound(Auth(a.GetReadyHandler)))
	a.Router.HandleFunc("/backend/MakePrediction", Inbound(Auth(a.MakePredictionHandler)))
	a.Router.HandleFunc("/backend/Connect", a.ConnectHandler)
}
