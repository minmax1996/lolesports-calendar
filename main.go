package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/minmax1996/lolesports-calendar/app"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/clients"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/types"
)

func main() {
	lolClient := clients.NewLolEsportsClient(os.Getenv("ESPORTS_TOKEN"))
	app := app.NewApp(lolClient)

	router := router{
		app: app,
	}

	r := mux.NewRouter()
	r.HandleFunc("/calendar", router.calendarHandler)
	log.Println("start")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

type router struct {
	app *app.App
}

func (rr router) calendarHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CalendarRequest
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	params.Normalize()
	cal, err := rr.app.CalendarHandler(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/ics")
	w.Header().Set("Content-Disposition", "attachment;filename=calendar.ics")
	w.Write(cal)
}
