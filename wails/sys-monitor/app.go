package main

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// LogConsoleInfo is a method that can be called from the frontend to log information
func (a *App) LogConsoleInfo(info string) {
	log.Infof("LogConsoleInfo called with info: %s", info)
}

// Person struct represents a person with a name, age, address, and number of times logged in
type Person struct {
	Name             string   `json:"name"`
	Age              uint8    `json:"age"`
	Address          *Address `json:"address"`
	NumTimesLoggedIn uint8    `json:"numTimesLoggedIn"`
}

// Address struct represents an address with a city and country
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

// GetPerson returns a Person struct with default values
func (a *App) GetPerson() Person {

	return Person{
		Name:             "John Doe",
		Age:              30,
		Address:          &Address{City: "New York", Country: "USA"},
		NumTimesLoggedIn: 0,
	}
}

// UpdateLoginCount increments the number of times a person has logged in and returns the updated person
func (a *App) UpdateLoginCount(p Person) Person {
	p.NumTimesLoggedIn++
	log.Infof("Updated login count for %s: %d", p.Name, p.NumTimesLoggedIn)

	return p
}
