package main

import "fmt"

func main() {
	//elements := map[string]int{"one": 1, "two": 2, "three": 3}
	var elements = make(map[string]string, 2)

	fmt.Println("element length is:", len(elements)) // element length is: 3

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println("element length is:", len(elements)) // element length is: 3

	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"

	fmt.Println("element length is:", len(elements)) // element length is: 5
	fmt.Println(elements)                            // map[B:Boron Be:Beryllium He:Helium Li:Lithium]

	delete(elements, "H")
	delete(elements, "Heeeee")
	fmt.Println("element length is:", len(elements)) // element length is: 5
	fmt.Println(elements)                            // map[B:Boron Be:Beryllium He:Helium Li:Lithium]

	for key, value := range elements {
		fmt.Println(key, value)
		//fmt.Printf("%p\n", &elements)
	}

	fmt.Println("")
	fmt.Println("=================== Movies ===================")
	movieRatings := map[string]int{
		"The Matrix":          88,
		"Speed":               94,
		"The Matrix Reloaded": 73,
		"John Wick":           86,
	}

	// Option #1 - create the 'val' variable to print the values of the map
	for key, val := range movieRatings {
		fmt.Println(key, val)
	}

	// Output:
	// The Matrix 88
	// Speed 94
	// The Matrix Reloaded 73
	// John Wick 86

	fmt.Println("")
	// Option #2 - pass the 'key' variable within [] square brackets after the map's name
	for key := range movieRatings {
		fmt.Println(key, movieRatings[key])
	}

	fmt.Println("")
	// Increase the ratings in the 'movieRatings' map by 5
	for key, val := range movieRatings {
		movieRatings[key] = val + 5
	}

	for key, val := range movieRatings {
		fmt.Println(key, val)
	}

	// Output:
	// The Matrix 93
	// Speed 99
	// The Matrix Reloaded 78
	// John Wick 91

	fmt.Println("")
	fmt.Println("=================== movieCharacters ===================")
	movieCharacters := map[string][]string{
		"Neo":       {"Noodles", "Sushi"},
		"John Wick": {"Steak", "Bacon"},
	}

	// Attempting to modify the slices directly during iteration won't work
	for _, foods := range movieCharacters {
		foods = append(foods, "Pizza") // This modifies a copy of the slice, not the original
	}

	fmt.Println(movieCharacters) // map[Neo:[Noodles Sushi] John Wick:[Steak Bacon]]

	// To modify the original slices, update them through the map key
	for key := range movieCharacters {
		movieCharacters[key] = append(movieCharacters[key], "Pizza")
	}
	//for key, value := range movieCharacters {
	//	movieCharacters[key] = append(value, "Pizza")
	//}

	fmt.Println(movieCharacters) // map[Neo:[Noodles Sushi Pizza] John Wick:[Steak Bacon Pizza]]

	fmt.Println("")
	fmt.Println("=================== Sets ===================")

	vegetables := map[string]bool{
		"🥕": true,
		"🧅": true,
		"🥦": true,
	}

	if _, ok := vegetables["🥕"]; ok {
		fmt.Println("🥕 is in the set.")
	}

	if _, ok := vegetables["🍇"]; ok {
		fmt.Println("🍇 is in the set.")
	}

	// Output:
	// 🥕 is in the set

	fruits := map[string]struct{}{
		"🍎": struct{}{},
		"🍊": struct{}{},
		"🥝": struct{}{},
	}
	fmt.Println(fruits)
}
