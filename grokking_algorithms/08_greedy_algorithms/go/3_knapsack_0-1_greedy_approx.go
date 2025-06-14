package main

import (
	"fmt"
	"time"
)

// Example of approximate greedy solution for 0/1 knapsack style problem - exercise 8.2

// You’re going to Europe, and you have seven days to see everything
// you can. You assign a point value to each item (how much you
// want to see it) and estimate how long it takes. How can you
// maximize the point total (seeing all the things you really want to
// see) during your stay? Come up with a greedy strategy. Will that
// give you the optimal solution?
// Answer: Keep picking the activity with the highest point value that
// you can still do in the time you have left. Stop when you can’t do
// anything else. No, this won’t give you the optimal solution.
type Activity struct {
	name     string
	duration time.Duration // aka weight
	points   int           // aka value
}

func main() {
	timeLeft := 7 * 24 * time.Hour // aka capacity
	activities := map[string]Activity{
		"vienna museums":   {duration: 24 * time.Hour, name: "vienna museums", points: 45},
		"rome coliseums":   {duration: 15 * time.Hour, name: "rome coliseums", points: 23},
		"eifel tower":      {duration: 13 * time.Hour, name: "eifel tower", points: 67},
		"croatian beaches": {duration: 3 * 24 * time.Hour, name: "croatian beaches", points: 89},
		"london":           {duration: 2 * 24 * time.Hour, name: "london", points: 5},
	}

	for _, a := range pickActivities(activities, timeLeft) {
		fmt.Printf("%s - points: %d, duration: %v\n", a.name, a.points, a.duration)
	}
}

// pick the activity with the highest point value that you can still do in the time you have left
func pickActivity(activities map[string]Activity, timeLeft time.Duration) *Activity {
	var best *Activity
	for _, a := range activities {
		if a.duration > timeLeft { // skip activities which we can't do in the time we have left
			continue
		}
		// we could implement a bit smarter greedy strategy which would look at points per hour, ex:
		// score := float64(a.points) / a.duration.Hours()
		if best == nil || a.points > best.points {
			tmp := a
			best = &tmp
		}
	}
	return best
}

func pickActivities(activities map[string]Activity, timeLeft time.Duration) []Activity {
	var pickedActivities []Activity
	for {
		activity := pickActivity(activities, timeLeft)
		if activity == nil {
			break // stop when we can't find an activity to do in the time left anymore
		}
		pickedActivities = append(pickedActivities, *activity)
		delete(activities, activity.name) // delete the activity from the map to reduce future work
		timeLeft -= activity.duration
	}
	fmt.Println("time left:", timeLeft)
	return pickedActivities
}
