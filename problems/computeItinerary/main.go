package main

import (
	"fmt"
	"strings"
)

//Given an unordered list of flights taken by someone, each represented as (origin, destination) pairs,
//and a starting airport, compute the person's itinerary. If no such itinerary exists, return null.
//If there are multiple possible itineraries, return the lexicographically smallest one. All flights must be used in the itinerary.

func main() {
	flights := [][]string{{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"}}
	fmt.Println(itinerary(flights, "YUL"))

	flights = [][]string{{"SFO", "COM"}, {"COM", "YYZ"}}
	fmt.Println(itinerary(flights, "COM"))

	flights = [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(itinerary(flights, "JFK"))
}

func itinerary(flights [][]string, start string) []string {
	flightMap := make(map[string][]string)
	for _, f := range flights {
		flightMap[f[0]] = append(flightMap[f[0]], f[1])
	}
	res := []string{start}
	return traverse(flightMap, start, len(flights), res)

}

func traverse(flightMap map[string][]string, loc string, length int, res []string) []string {
	if len(res) == length {
		if len(flightMap[loc]) == 0 {
			return nil
		}
		return append(res, flightMap[loc][0])
	}

	options := flightMap[loc]
	var bestOption []string

	for i, o := range options {
		optionTemp := make([]string, len(options)-1)
		copy(optionTemp[:i], options[:i])
		copy(optionTemp[i:], options[i+1:])
		flightMap[loc] = optionTemp
		res = append(res, o)
		temp := traverse(flightMap, o, length, res)

		if bestOption == nil || (temp != nil && strings.Join(temp, "") < strings.Join(bestOption, "")) {
			bestOption = temp
		}
		res = res[:len(res)-1]
		flightMap[loc] = options
	}

	if len(bestOption) != length+1 {
		return nil
	}

	return bestOption
}
