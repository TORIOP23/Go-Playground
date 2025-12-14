package main

import "sort"

func main() {
	println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})) // 3
}

func carFleet(target int, position []int, speed []int) int {
	type car struct {
		position int
		speed    int
	}

	cars := make([]car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = car{position[i], speed[i]}
	}

	// Sort cars by position in descending order using sort.Slice (O(n log n))
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleetCount := 0
	var lastArrivalTime = -1.0

	// Cars are processed from closest to farthest from the target.
	// Compute each car’s arrival time. If it’s greater than the last
	// recorded arrival time, it can’t catch the fleet ahead and forms
	// a new fleet; otherwise it joins the existing fleet.
	// Sorting costs O(n log n); this pass is O(n).
	for _, c := range cars {
		arrivalTime := float64(target-c.position) / float64(c.speed)
		if arrivalTime > lastArrivalTime {
			fleetCount++
			lastArrivalTime = arrivalTime
		}
	}

	return fleetCount
}
