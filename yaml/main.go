package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

func structToYAML() {
	type Movie struct {
		Title       string
		Genres      []string
		Year        int
		movieLength int32 // in minutes
		Rating      float32
		CreatedAt   time.Time
	}

	movie := Movie{
		Title:       "Titanic",
		Genres:      []string{"drama", "romance"},
		Year:        1997,
		movieLength: 197,
		Rating:      7.9,
		CreatedAt:   time.Now(),
	}
	// Serializing the `movie` struct into a YAML object
	// `yaml.Marshal()` returns a slice of bytes and an error:
	movieYAML, err := yaml.Marshal(movie)
	if err != nil {
		log.Fatal(err)
	}
	// Remember to "cast" the returned slice of bytes as a 'string' to properly print it:
	fmt.Println(string(movieYAML))
}

func structToYAMLCustom() {
	// Observe the inclusion of yaml:"<key-name>" tags next to each field in the following struct. These tags are utilized to instruct the YAML package to customize the keys in the resulting YAML object.
	type Movie struct {
		Title     string    `yaml:"title"`
		Genres    []string  `yaml:"genres"`
		Year      int       `yaml:"year"`
		Runtime   int32     `yaml:"movie_length_in_minutes"`
		Rating    float32   `yaml:"rating"`
		CreatedAt time.Time `yaml:"created_at"`
	}

	movie := Movie{
		Title:     "Titanic",
		Genres:    []string{"drama", "romance"},
		Year:      1997,
		Runtime:   197,
		Rating:    7.9,
		CreatedAt: time.Now(),
	}

	movieYAML, err := yaml.Marshal(movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(movieYAML))
}

func structToYAMLWithOptionalTags() {
	type User struct {
		ID        int64     `yaml:"id"`
		CreatedAt time.Time `yaml:"created_at"`
		Name      string    `yaml:"name"`
		Email     string    `yaml:"email"`
		Password  string    `yaml:"-"` // Note the usage of the `-` tag
	}

	user := User{
		ID:        6001,
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Password:  "very secret",
		CreatedAt: time.Now(),
	}

	userYAML, err := yaml.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(userYAML))
}

func structToYAMLWithInlineDirective() {
	type Address struct {
		City    string `yaml:"city"`
		Country string `yaml:"country"`
	}

	// Observe how we use the 'inline' directive for the 'Address' field
	type Person1 struct {
		Name    string `yaml:"name"`
		Age     int    `yaml:"age"`
		Address `yaml:",inline"`
	}

	// Observe that for the following struct we do not use the 'inline' directive
	type Person2 struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
		Address
	}

	addr := Address{
		City:    "New York",
		Country: "USA",
	}

	person1 := Person1{
		Name:    "John Doe",
		Age:     39,
		Address: addr,
	}

	person2 := Person2{
		Name:    "John Doe",
		Age:     39,
		Address: addr,
	}

	// encoding `person1`:
	person1YAML, err := yaml.Marshal(person1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(person1YAML))

	// encoding `person2`:
	person2YAML, err := yaml.Marshal(person2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(person2YAML))
}

func structToYAMLWithOmitemptyDirective() {
	type Movie1 struct {
		Title     string    `yaml:"title"`
		Genres    []string  `yaml:"genres"`
		Year      int       `yaml:"year"`
		Runtime   int32     `yaml:"movie_length_in_minutes"`
		Rating    float32   `yaml:"rating"`
		CreatedAt time.Time `yaml:"created_at,omitempty"` // Note the `omitempty` tag
	}

	type Movie2 struct {
		Title     string    `yaml:"title"`
		Genres    []string  `yaml:"genres"`
		Year      int       `yaml:"year"`
		Runtime   int32     `yaml:"movie_length_in_minutes"`
		Rating    float32   `yaml:"rating"`
		CreatedAt time.Time `yaml:"created_at"` // Note that no `omitempty` tag is used
	}

	movie1 := Movie1{
		Title:   "Titanic",
		Genres:  []string{"drama", "romance"},
		Year:    1997,
		Runtime: 197,
		Rating:  7.9,
	}

	movie2 := Movie2{
		Title:   "Titanic",
		Genres:  []string{"drama", "romance"},
		Year:    1997,
		Runtime: 197,
		Rating:  7.9,
	}

	// serializing the `movie1` struct into a YAML object:
	movie1YAML, err := yaml.Marshal(movie1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(movie1YAML))

	// serializing the `movie2` struct into a YAML object:
	movie2YAML, err := yaml.Marshal(movie2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(movie2YAML))
}

func YAMLToStruct() {
	type EmailServer struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	type EmailClient struct {
		Email    string `yaml:"email"`
		Password string `yaml:"password"`
		SendTo   string `yaml:"send_to"`
	}

	type Config struct {
		Server EmailServer `yaml:"server"`
		Client EmailClient `yaml:"client"`
	}

	yamlData := `
server:
  host: "127.0.0.1"
  port: 2500
client:
  email: "sender@example.org"
  password: "123456"
  send_to: "recipient@example.net"
`
	var config Config
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Server: %+v\n", config.Server)
	fmt.Printf("Client: %+v\n", config.Client)

	// Output:
	// Server: {Host:127.0.0.1 Port:2500}
	// Client: {Email:sender@example.org Password:123456 SendTo:recipient@example.net}
}

func main() {
	fmt.Println("Serializing structured YAML")

	// structToYAML()
	// structToYAMLCustom()
	// structToYAMLWithOptionalTags()
	// structToYAMLWithInlineDirective()
	// structToYAMLWithOmitemptyDirective()
	YAMLToStruct()
}
