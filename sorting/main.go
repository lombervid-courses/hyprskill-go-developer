package main

import (
	"fmt"
	"math"
	"sort"
)

func sliceOfIntegers() {
	fmt.Println("\n\nSorting a slice of integers\n")

	nums := []int{7, 2, 8, -9, 4, 0, -1}
	fmt.Printf("%v <- Unsorted slice\n", nums)

	sort.Ints(nums)
	fmt.Printf("%v <- Sorted in ascending order\n", nums)

	// Output:
	// [7 2 8 -9 4 0 -1] <- Unsorted slice
	// [-9 -1 0 2 4 7 8] <- Sorted in ascending order
}

func sliceOfFloats() {
	fmt.Println("\n\nSorting a slice of floats\n")

	nums := []float64{3.1416, 2.7182, 1.6180, -273.15, math.NaN(), math.Inf(1), math.Inf(-1)}
	fmt.Printf("%v <- Unsorted slice\n", nums)

	sort.Float64s(nums)
	fmt.Printf("%v <- Sorted in ascending order\n", nums)

	// Output:
	// [3.1416 2.7182 1.618 -273.15 NaN +Inf -Inf] <- Unsorted slice
	// [NaN -Inf -273.15 1.618 2.7182 3.1416 +Inf] <- Sorted in ascending order
}

func sliceOfStrings() {
	fmt.Println("\n\nSorting a slice of strings\n")

	lastNames := []string{"Messi", "Salah", "Ronaldo", "Haaland", "Mbappé"}
	fmt.Printf("%v <- Unsorted slice\n", lastNames)

	sort.Strings(lastNames)
	fmt.Printf("%v <- Sorted in alphabetical order\n", lastNames)

	// Output:
	// [Messi Salah Ronaldo Haaland Mbappé] <- Unsorted slice
	// [Haaland Mbappé Messi Ronaldo Salah] <- Sorted in alphabetical order

	/*
	 * https://webdevstation.com/posts/how-to-sort-strings-with-go-alphabetically-in-any-language/
	 * https://pkg.go.dev/golang.org/x/text/collate
	 */
	fmt.Println("\nAccented words\n")

	cities := []string{"Ürkmez", "İstanbul", "München", "Muğla", "Ulm", "Zürich"}
	fmt.Printf("%v <- Unsorted slice\n", cities)

	sort.Strings(cities)
	fmt.Printf("%v <- Sorted in alphabetical order\n", cities)

	// Output:
	// [Ürkmez İstanbul München Muğla Ulm Zürich] <- Unsorted slice
	// [Muğla München Ulm Zürich Ürkmez İstanbul] <- Sorted in alphabetical order
}

func sortSliceFunction() {
	fmt.Println("\n\nThe sort.Slice() function\n")

	words := []string{"Epsilon", "Zeta", "Gamma", "Alpha", "Beta"}
	fmt.Printf("%v <- Unsorted 'words' slice\n", words)

	// Here we use the '>' operator to return the words in descending order (from z to a)
	sort.Slice(words, func(i, j int) bool { return words[i] > words[j] })

	fmt.Printf("%v <- Sorted 'words' in descending order\n", words)

	// Output:
	// [Epsilon Zeta Gamma Alpha Beta] <- Unsorted 'words' slice
	// [Zeta Gamma Epsilon Beta Alpha] <- Sorted 'words' in descending order
}

func sortStructs() {
	fmt.Println("\n\nSorting structs with multiple fields\n")

	type Student struct {
		fullName string
		score    float64
	}

	students := []Student{
		{"Harry Potter", 100},
		{"Hermione Granger", 100},
		{"Ron Weasley", 80},
		{"Draco Malfoy", 95},
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].score != students[j].score {
			// here, we sort the students by the highest score first
			return students[i].score > students[j].score
		}

		// then, if any students have the same score,
		//we sort them by their 'fullName' in alphabetical order
		return students[i].fullName < students[j].fullName
	})

	for _, s := range students {
		fmt.Println(s.fullName, s.score)
	}

	// Output:
	// Harry Potter 100
	// Hermione Granger 100
	// Draco Malfoy 95
	// Ron Weasley 80
}

func stableSortForSlices() {
	fmt.Println("\n\nStable sort for slices\n")

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{"Amy", 25},
		{"Eli", 75},
		{"Amy", 75},
		{"Bob", 75},
		{"Bob", 25},
		{"Joe", 25},
		{"Eli", 25},
	}

	// sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].name < people[j].name })
	fmt.Printf("%v <- Stable sort - by name\n", people)

	// sort by age, preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].age < people[j].age })
	fmt.Printf("%v <- Stable sort - by age & name", people)

	// Output:
	// [{Amy 25} {Amy 75} {Bob 75} {Bob 25} {Eli 75} {Eli 25} {Joe 25}] <- Stable sort - by name
	// [{Amy 25} {Bob 25} {Eli 25} {Joe 25} {Amy 75} {Bob 75} {Eli 75}] <- Stable sort - by age & name
}

func main() {
	fmt.Println("Sorting slices")

	sliceOfIntegers()
	sliceOfFloats()
	sliceOfStrings()
	sortSliceFunction()
	sortStructs()
	stableSortForSlices()
}
