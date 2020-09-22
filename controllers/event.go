package controllers

import (
	"fmt"
	"github.com/UoYMathSoc/2020-site/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
)

type EventController struct {
	Controller
}

func NewEventController(c *structs.Config, q *database.Queries) *EventController {
	return &EventController{Controller{config: c, querier: q}}
}

func (eventC *EventController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	eventM := models.NewEventModel(eventC.querier)
	event, err := eventM.Get(int32(id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !event.EndTime.Valid {
		event.EndTime.Time = event.StartTime.Add(time.Hour)
		event.EndTime.Valid = true
	}

	fmt.Fprintf(w, "%s\n", event.Name)
	fmt.Fprintf(w, "We are going to %s on %s, so get ready!\n", event.Location.String, event.Date.Weekday())
	fmt.Fprintf(w, "%s\n", event.Description.String)
	fmt.Fprintf(w, "%d:%d-%d:%d", event.StartTime.Hour(), event.StartTime.Minute(), event.EndTime.Time.Hour(), event.EndTime.Time.Minute())
}
