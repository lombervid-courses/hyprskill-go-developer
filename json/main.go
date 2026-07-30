package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func serialize() {
	actor := map[string]interface{}{"name": "Keanu Reeves", "age": 57}
	// actor := map[string]interface{}{"name": "Keanu Reeves", "attr": map[string]interface{}{"one": 1, "another": "yep"}}

	// The json.Marshal() function encodes/marshals the 'actor' map
	actorJson, err := json.Marshal(actor)
	if err != nil {
		log.Fatal(err)
	}

	// We need to "cast" the returned slice of bytes as a string to properly print it:
	fmt.Println(string(actorJson))
}

func serializeIndent() {
	actor := map[string]interface{}{"name": "Keanu Reeves", "age": 57}
	actorJson, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(actorJson))
}

func serializeSlice() {
	carBrands := []string{"Tesla", "BMW", "Toyota", "Ford"}

	carBrandsJson, err := json.Marshal(carBrands)
	// carBrandsJson, err := json.MarshalIndent(carBrands, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(carBrandsJson))
}

func serializeNestedObjects() {
	books := map[string]interface{}{
		"books": []interface{}{
			map[string]interface{}{
				"isbn":   "9781491941959",
				"title":  "Introducing Go",
				"author": "Caleb Doxsey",
				"pages":  124,
			},
			map[string]interface{}{
				"isbn":   "9781491941960",
				"title":  "Eragon",
				"author": "Christopher Paolini",
				"pages":  570,
			},
		},
	}

	booksJson, err := json.MarshalIndent(books, "", " ")
	// booksJson, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(booksJson))
}

func deserialize() {
	actorJson := `{"name": "Will Smith", "age": 53}`

	// Create the 'actor' map that will contain the decoded/unmarshaled JSON data:
	var actor map[string]interface{}

	// We need to "cast" the JSON object as slice of bytes '[]byte()' to properly decode it,
	// and pass a pointer '&actor' to decode the JSON object into the 'actor' variable:
	err := json.Unmarshal([]byte(actorJson), &actor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(actor)
}

func main() {
	fmt.Println("Working with JSON")
	// serialize()
	// serializeIndent()
	// serializeSlice()
	// serializeNestedObjects()
	deserialize()
}
