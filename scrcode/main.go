package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Welcomes you home! Daniel")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var newEvent event

	for i := 0; i < len(events); i++ {
		if events[i].ID == eventID {
			newEvent = events[i]
			break
		}

	}

	json.NewEncoder(w).Encode(newEvent)

	// for _, singleEvent := range events {
	// 	if singleEvent.ID == eventID {

	// 		json.NewEncoder(w).Encode(singleEvent)
	// 	}
	// }
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	var index int = -1
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			index = i
		}
	}
	var singleEvent event
	if index >= 0 {
		singleEvent = events[index]
		singleEvent.ID = eventID
		singleEvent.Title = updatedEvent.Title
		singleEvent.Description = updatedEvent.Description

		events[index] = singleEvent
	}
	json.NewEncoder(w).Encode(singleEvent)

}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/welcome", homeLink)
	router.HandleFunc("/events", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
