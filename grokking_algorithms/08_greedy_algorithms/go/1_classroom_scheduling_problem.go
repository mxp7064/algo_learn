/*
CLASSROOM SCHEDULING PROBLEM – GREEDY ALGORITHM

PROBLEM STATEMENT:
Given a set of classes, each with a start and end time, schedule the maximum number of non-overlapping classes.

You are NOT trying to maximize the total duration, only the number of non-overlapping classes that can fit in a single day (i.e., greedy selection by compatibility).

GOAL:
Return a list of classes such that:
- No classes overlap
- The list contains the maximum possible number of classes

GREEDY STRATEGY:
Always select the class that ends the earliest (among those that do not overlap with already selected classes).
This maximizes the remaining time for scheduling additional classes.

WHY EARLIEST-ENDING?
Classes that end sooner leave more room for other classes later. Choosing long-running or late-ending classes would block more time and prevent optimal schedules.

APPROACHES:
1. Sorted slice approach (optimal)
2. Map scan approach (more direct, less efficient)

PATTERN USED:
This uses the "nil or better" pattern:
if best == nil || current is better → pick it. It picks the first or the best element.
This avoids using dummy/extreme initial values and handles edge cases gracefully.

TIME COMPLEXITY:

1. Sorted slice:
  - Sort classes by end time → O(n log n)
  - Iterate once and greedily select → O(n)
	- we ignore this because it's small compared to O(n log n)
  - Total: O(n log n)

2. Map scan:
  - For each selected class -> O(n), scan the map → O(n)
  - Total: O(n²) in the worst case

Sorted slice is preferred in interviews due to better performance and clarity.

*/

package main

import (
	"fmt"
	"sort"
	"time"
)

type Class struct {
	Name  string
	Start time.Time
	End   time.Time
}

// createClassTime creates a time.Time value with only hour and minute set, using a fixed base date
func createClassTime(hour int, minutes int) time.Time {
	now := time.Unix(0, 0) // fixed reference point
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minutes, 0, 0, time.UTC)
}

// getClassScheduleSortedSliceApproach implements the optimal O(n log n) solution
// by sorting classes by end time and greedily selecting non-overlapping ones
func getClassScheduleSortedSliceApproach(classes []Class) []Class {
	var result []Class

	// Step 1: sort all classes by their end time (greedy criterion)
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].End.Before(classes[j].End)
	})

	// Step 2: iterate through sorted list and add to result if class doesn't overlap with the last scheduled one
	for _, class := range classes {
		// If result is empty OR class starts after the last scheduled class ends
		if len(result) == 0 || class.Start.After(result[len(result)-1].End) { // variation of "nil or better" pattern
			result = append(result, class)
		}
	}

	return result
}

// getClassScheduleMapApproach is the original Grokking-like version using a map
// It finds the soonest-ending non-overlapping class in each iteration via full scan
func getClassScheduleMapApproach(classes map[string]Class) []Class {
	classesCopy := make(map[string]Class) // Make a fresh copy of the map since the function deletes entries
	for k, v := range classes {
		classesCopy[k] = v
	}

	var result []Class
	var lastScheduled *Class

	for len(classesCopy) > 0 {
		// Find the class which starts after the last scheduled and which ends earliest
		lastScheduled = findEarliestNonOverlapping(classesCopy, lastScheduled)
		if lastScheduled == nil { // no such classes exist
			break
		}
		result = append(result, *lastScheduled)
		delete(classesCopy, lastScheduled.Name) // delete selected class to reduce future work
	}

	return result
}

// findEarliestNonOverlapping scans the map and selects the earliest-ending class that doesn't overlap
func findEarliestNonOverlapping(classes map[string]Class, lastScheduledClass *Class) (result *Class) {
	for _, class := range classes {
		isNonOverlapping := lastScheduledClass == nil || class.Start.After(lastScheduledClass.End) // Nil-or-better pattern: pick first or non-overlapping one
		isEarlier := result == nil || class.End.Before(result.End)                                 // Nil-or-better pattern: pick first or earlier
		if isNonOverlapping && isEarlier {                                                         // conditions are independent so we can merge them with '&&'
			result = &class
		}
	}
	return result
}

func main() {
	// Example dataset (map-based for Grokking version)
	classes := map[string]Class{
		"A": {"A", createClassTime(9, 0), createClassTime(10, 0)},
		"B": {"B", createClassTime(9, 30), createClassTime(12, 0)},
		"C": {"C", createClassTime(10, 15), createClassTime(11, 15)},
		"D": {"D", createClassTime(11, 30), createClassTime(13, 0)},
		"E": {"E", createClassTime(12, 0), createClassTime(13, 0)},
	}

	// For sorted slice approach, convert map to slice
	var classList []Class
	for _, class := range classes {
		classList = append(classList, class)
	}

	// Run sorted slice approach
	fmt.Println("Sorted slice approach:")
	sortedSchedule := getClassScheduleSortedSliceApproach(classList)
	for _, class := range sortedSchedule {
		fmt.Printf("%s: %s - %s\n", class.Name, class.Start.Format("15:04"), class.End.Format("15:04"))
	}

	// Run map scan approach
	fmt.Println("\nMap scan approach (Grokking style):")
	mapSchedule := getClassScheduleMapApproach(classes)
	for _, class := range mapSchedule {
		fmt.Printf("%s: %s - %s\n", class.Name, class.Start.Format("15:04"), class.End.Format("15:04"))
	}
}
