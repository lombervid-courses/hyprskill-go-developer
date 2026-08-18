package main

import (
	"fmt"
	"time"
)

func timeInitialization() {
	fmt.Println("\n\nTime initialization\n")

	//                    YEAR   MONTH    D  H  M  S  nS Location
	starWars := time.Date(1977, time.May, 4, 0, 0, 0, 0, time.UTC)

	fmt.Println(starWars)
	fmt.Println(starWars.Date())
	fmt.Println(time.Now())

	// Output:
	// 1977-05-04 00:00:00 +0000 UTC
	// 1977 May 4
	// 2026-08-18 16:26:31.830445465 -0600 CST m=+0.000057341
}

func timeComparison() {
	fmt.Println("\n\nTime comparison\n")

	darwin, _ := time.LoadLocation("Australia/Darwin")

	presentMcFly := time.Date(1985, time.October, 26, 1, 20, 0, 0, time.UTC)
	futureMcFly := time.Date(1955, time.November, 12, 22, 4, 0, 0, time.UTC)

	fmt.Println(presentMcFly.Before(futureMcFly))
	fmt.Println(presentMcFly.After(futureMcFly))
	fmt.Println(presentMcFly.Equal(futureMcFly))

	futureMcFlyInDarwin := time.Date(1955, time.November, 13, 7, 34, 0, 0, darwin)

	fmt.Println(futureMcFly.Equal(futureMcFlyInDarwin))
	fmt.Println(futureMcFly == futureMcFlyInDarwin)

	// Output:
	// false
	// true
	// false
	// true
	// false
}

func operationsWithTime() {
	fmt.Println("\n\nOperations with time\n")

	twoSecond := time.Millisecond*1000 + time.Second
	sevenDays := time.Hour * 24 * 7
	oneYear := time.Hour * 24 * 365
	threeYearsAndTwoMonths := oneYear*3 + time.Hour*24*30

	fmt.Println(twoSecond)
	fmt.Println(sevenDays)
	fmt.Println(oneYear)
	fmt.Println(threeYearsAndTwoMonths)

	// Output:
	// 2s
	// 168h0m0s
	// 8760h0m0s
	// 27000h0m0s

	CopernicusBDay := time.Date(1473, time.February, 19, 0, 0, 0, 0, time.UTC)
	NewtonBDay := CopernicusBDay.Add(time.Hour * 1489080)
	EinsteinBDay := time.Date(1879, time.March, 14, 0, 0, 0, 0, time.UTC)

	fmt.Println("\nTimeline:")
	fmt.Printf("Nicolaus Copernicus: %v\n", CopernicusBDay)
	fmt.Printf(" | %v\n", NewtonBDay.Sub(CopernicusBDay))
	fmt.Printf("Isaac Newton: %v\n", NewtonBDay)
	fmt.Printf(" | %v\n", EinsteinBDay.Sub(NewtonBDay))
	fmt.Printf("Albert Einstein: %v\n", EinsteinBDay)

	// Output:
	// Timeline:
	// Nicolaus Copernicus: 1473-02-19 00:00:00 +0000 UTC
	//  | 1489080h0m0s
	// Isaac Newton: 1643-01-04 00:00:00 +0000 UTC
	//  | 2070384h0m0s
	// Albert Einstein: 1879-03-14 00:00:00 +0000 UTC
}

func main() {
	fmt.Println("Working with time")

	timeInitialization()
	timeComparison()
	operationsWithTime()
}
