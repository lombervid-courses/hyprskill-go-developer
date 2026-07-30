package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func serializeAndPrint(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

func structToJson() {
	type User struct {
		ID        int
		IsActive  bool
		LastLogin time.Time
		email     string
	}

	usr := User{ID: 343, IsActive: true, LastLogin: time.Now(), email: "sus@among.us"}
	serializeAndPrint(usr)
}

func structToJsonCustom() {
	type User struct {
		ID        int       `json:"id"`
		IsActive  bool      `json:"isActive"`
		LastLogin time.Time `json:"lastLogin"`
		Email     string    `json:"email"`
	}

	usr := User{ID: 343, IsActive: true, LastLogin: time.Now(), Email: "sus@among.us"}
	serializeAndPrint(usr)

}

func structToJsonWithOptionalTags() {
	// `omitempty` omits a struct field when it contains a zero/default value;
	// we can use the hyphen `-` directive when we require a field in a struct to be public and accessible to other packages but prevent it from being encoded into JSON;
	// `string` forces the data in an individual field to be encoded as a string type in the resulting JSON.
	type User struct {
		ID        int       `json:"id,string"`
		IsActive  bool      `json:"isActive,omitempty"`
		LastLogin time.Time `json:"lastLogin,omitempty"`
		Email     string    `json:"-"`
	}

	usr := User{ID: 343, Email: "sus@among.us"}
	serializeAndPrint(usr)

}

func jsonToStruct() {
	// Create the ErrJSONObject struct that will contain the decoded JSON data:
	type ErrJSONObject struct {
		Errors []struct {
			Source struct {
				Pointer string `json:"pointer"`
			} `json:"source"`
			Detail string `json:"detail"`
		} `json:"errors"`
	}

	errJSON := `{
      "errors": [
        {
          "source": { "pointer": "" },
          "detail":  "Missing 'data' Member at document's top level."
        }
      ]
    }`

	var errObj ErrJSONObject // create an instance of the 'ErrJSONObject' struct type

	// Cast the JSON object as a slice of bytes to properly decode it,
	// and pass a pointer to the '&errObj' struct we plan to decode the JSON object to:
	err := json.Unmarshal([]byte(errJSON), &errObj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", errObj)
}

func main() {
	fmt.Println("Serializing structured JSON")

	structToJson()
	structToJsonCustom()
	structToJsonWithOptionalTags()
	jsonToStruct()
}
