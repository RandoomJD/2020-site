package main

import (
	"database/sql"
	"fmt"
	"github.com/UoYMathSoc/2020-site/controllers"
	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Server is the type of the main 2020site web application.
type Server struct {
	*negroni.Negroni
}

// NewServer creates a 2020site server based on the config c.
func NewServer(c *structs.Config) (*Server, error) {

	s := Server{negroni.Classic()}

	router := mux.NewRouter().StrictSlash(true)

	getRouter := router.Methods("GET").Subrouter()
	postRouter := router.Methods("POST").Subrouter()
	//headRouter := router.Methods("HEAD").Subrouter()

	db := c.Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	q := database.New(conn)

	// Routes go in here
	loginC := controllers.NewLoginController(c, q)
	postRouter.HandleFunc("/login/", loginC.Post)

	userC := controllers.NewUserController(c, q)
	getRouter.HandleFunc("/user/{id}", userC.Get)

	calendarC := controllers.NewCalendarController(c, q)
	getRouter.HandleFunc("/calendar/ical/MathSoc.ics", calendarC.GetICal)

	staticC := controllers.NewStaticController(c)
	getRouter.HandleFunc("/", staticC.GetIndex)
	getRouter.HandleFunc("/login/", staticC.GetLogin)
	// End routes

	s.UseHandler(router)

	return &s, nil

}
