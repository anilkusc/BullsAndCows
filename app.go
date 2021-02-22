package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// App method is the main struct for the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init method is initialized the configs,routes,etc.
func (a *App) Init(database string, username string, password string) error {
	log.Println("Creating Router...")
	a.Router = mux.NewRouter()
	log.Println("Initializing Routes...")
	a.initRoutes()
	if _, err := os.Stat(database); os.IsNotExist(err) {
		log.Println("Database can not be found.Creating new...")
		file, err := os.Create(database) // Create SQLite file
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		db, err := sql.Open("sqlite3", database)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		log.Println("Created database file: " + database)

		query := "CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);"
		statement, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Users table")

		query = "CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOY NULL DEFAULT 0);"
		statement, err = db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Sessions table")

		query = "CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT -1,Negative INTEGER DEFAULT -1,Turn INTEGER NOT NULL,Player1 TEXT,Player2 TEXT,Player1Number INTEGER,Player2Number INTEGER,Predictor INTEGER,Prediction INTEGER,Action TEXT,FOREIGN KEY (SessionId) REFERENCES Sessions (Id) ON DELETE CASCADE);"
		statement, err = db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Created Moves table")

		enableFK := "PRAGMA foreign_keys=ON;"
		statement, err = db.Prepare(enableFK)
		if err != nil {
			log.Fatal(err)
		}
		statement.Exec()
		log.Println("Foreign Key enabled for sqlite.")
		log.Println("Created tables")
		return db.Ping()
	} else {
		log.Println("Database file found.")
		var err error
		db, err := sql.Open("sqlite3", database)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		return db.Ping()
	}
}

// Run method runs the application with specified parameters
func (a *App) Run(addr string) {
	log.Println("Serving on: ", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
func (a *App) initRoutes() {
	a.Router.HandleFunc("/backend/ReadUser", a.ReadUser).Methods("POST")
}
