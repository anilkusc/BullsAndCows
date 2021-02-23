package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

var a App

// Init method is initialized the configs,routes,etc.
func TestInit(t *testing.T) {
	a.Init("test.db")
	e := os.Remove("test.db")
	if e != nil {
		log.Fatal(e)
	}
}
func TestInitRoutes(t *testing.T) {
	a.InitRoutes()
}

func TestRun(t *testing.T) {
	//a.Run(":8080")

}
