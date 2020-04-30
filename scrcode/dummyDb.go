package main

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	mapofE      map[string]event
}

type allEvents []event

type mapOfEvents map[string]event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

// var me = mapOfEvents{
// 	"1": event{
// 		ID:          "1",
// 		Title:       "Introduction to Golang1",
// 		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
// 	},

// 	"2": event{
// 		ID:          "2",
// 		Title:       "Introduction to Golang2",
// 		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
// 	},
// }
