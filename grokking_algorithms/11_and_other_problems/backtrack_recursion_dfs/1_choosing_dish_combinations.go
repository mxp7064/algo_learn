package main

import "fmt"

/*
Backtracking Intro: Choosing combinations of dishes from a menu

🧠 Problem statement:

We are given a menu with a list of dishes: ["Salad", "Chicken", "Rice"]
We want to generate all possible dish combinations of up to 2 dishes (no duplicates, no repeats).
We basically want to generate all subsets with length up to 2:

[]
[Salad]
[Salad, Chicken]
[Salad, Rice]
[Chicken]
[Chicken, Rice]
[Rice]

Other variation of this problem would be to only pick combinations of certain length (ex. 2):

[Salad, Chicken]
[Salad, Rice]
[Chicken, Rice]

But in the recursive logic we would still need to explore those combinations which have length < 2 until we come
to those which have length = 2 so it's basically the same problem/solution.

And we can have a variation of this problem which will explore all subsets (power set generation - 2_subset_generation.go)

🔑 This example shows how backtracking works:

- At each step, we make a choice: include the dish or not
- We explore further with that choice
- Then we "backtrack": undo that choice so we can try something else

Backtracking is like exploring all branches of a decision tree:

1. Make a decision (e.g., add "Chicken")
2. Recurse (check is decision ok or continue exploring picking other elements)
3. Undo the decision (remove "Chicken") → so we can try adding "Rice" instead

Every backtracking algorithm comes down to:

	func dfs(start int, path []string) {
		if len(path) == k { // in this example we only want subsets of length k
			result = append(result, copy(path))
			return
		}

		for i := start; i < len(input); i++ {
			path = append(path, input[i])   // choose element
			dfs(i+1, path)                  // explore with this element(s) by adding remaining elements onto path
			path = path[:len(path)-1]       // unchoose/backtrack so we can try adding the next element (in next iteration)
		}
	}

- `start` ensures we don’t revisit earlier elements
  - we are always moving forward, if B was picked, then the next one will be C, not A (previous one) or B (same one)
  - so we never pick the same element twice in the same path (like [A,A])
  - and we never revisit earlier elements, so permutations like [B,A] are avoided if [A,B] exists.

- `path` tracks current combination (depth wise)

- base case: when len(path) == k, we found one
  - or some other condition based on assigment
  - or without condition if we want to explore all combinations (subsets)

Note: We never revisit earlier elements, each recursive call only explores forward from the current index (i+1). This ensures:
- No duplicates
- No permutations (i.e., [A, B] and [B, A] are treated as the same combination)

🌳 Visualizations for input: ["A", "B", "C"], limit = 2):

Let's say we want to generate all subsets of length = 2 - we want this output:

[A,B]
[A,C]
[B,C]

Decision tree:

	          []
		    / \  \
		   A   B  C
		  / \   \
		AB  AC   BC

Each level represents a choice of adding an element or skipping to the next one.
The recursive tree explores all paths of size 0 to `limit`, using depth-first search.

So we first start from empty set, then we pick the first element (A) and then the logic is as follows:

- [A] -> len([A]) != 2 -> try adding next element (B)
- [A,B] -> len([A,B]) == 2 -> ✅ add to result, remove the last element (B) and try adding the next one (C)
- [A,C] -> len([A,C]) == 2 -> ✅ add to result, there are no more elements after C, now start picking from B
- [B] -> len([B]) != 2 -> try adding next element (C)
- [B,C] -> len([B,C]) == 2 -> ✅ add to result, there are no more elements after C, now start picking from C
- [C] -> len([C]) != 2 -> no other elements to add, end

Recursive call stack:

dfs(0, [])                   // path = []
│
├── dfs(1, [A])              // pick A
│   ├── dfs(2, [A B])        // pick B → ✅ result: [A B]
│   └── dfs(3, [A C])        // pick C → ✅ result: [A C]
│
├── dfs(2, [B])              // pick B
│   └── dfs(3, [B C])        // pick C → ✅ result: [B C]
│
└── dfs(3, [C])              // pick C → can't go deeper

🌳 Why do we talk about "trees" in backtracking / DFS problems?

This has nothing to do with explicit tree data structures (like binary trees).
The tree here is a decision tree that naturally emerges from recursive DFS/backtracking behaviour.

Why this decision tree forms:
1. Each recursive call = one decision
  - Whenever we pick an item (e.g. a dish), we go one level deeper in the tree.
  - Each level = a decision layer.
  - For example: "pick A" → [A] (level 1), then "pick B" → [A,B] (level 2)

2. Each decision leads to multiple choices
  - From any state, we can branch into several possible next steps.
  - These branches split further = new “children” in the tree.
  - For example: from [A,B], we unchoose B and pick C → new branch [A,C]

3. No cycles occur
  - We avoid revisiting earlier choices (e.g. by using `start`).
  - Ensures we never revisit same element - like [A,A] - instead we only progress forward → [A,B], [A,C]
  - So it behaves like a tree (connected, acyclic graph).

4. DFS traversal
  - We go deep into one path first before backtracking → classic depth-first style.
  - Example traversal: [], [A], [A,B], backtrack, [A,C], backtrack, [B], [B,C], backtrack, [C]
  - So we have: [], [A], [A,B], [A,C], [B], [B,C], [C]

What’s the root?
- The root is just the starting state (e.g., an empty subset or an empty plate).
- We don't explicitly refer to a “rooted” tree, but visually it starts from this point.

Analogy:
Start with an empty plate (root). Try adding "Salad", go deeper. Backtrack, try "Chicken" instead.
Every “add” creates a new branch, every “backtrack” returns to the parent decision from which we do further branching
on remaining choices under that parent.

This structure is what makes backtracking problems look and behave like trees.

💡 Why we need to copy `plate` before appending to result?

A slice in go is a small struct (called slice header) which contains fields: length, capacity and a pointer reference to the underlying array.
Imagine that we didn't copy the `plate`, we would have the following code:

	for i := start; i < len(menu); i++ {
		plate = append(plate, menu[i])
		backtrack(i + 1) // calls `result = append(result, plate) without copying`
		plate = plate[:len(plate)-1]
	}

Let's analyze in detail what happens here. We will have the following successive calls (in loop):

- backtrack(i + 1)
  - which calls: `result = append(result, plate)`
  - this appends a copy of the `plate` slice header (plate is passed by value to append) to the result
  - this slice header copy holds a reference to our underlying array

- plate = plate[:len(plate)-1]
  - trims/removes last element
  - reduces length by 1 (changes the "view window" of the array)
  - capacity stays the same
  - underlying array remains unchanged
  - plate still holds the reference to the underlying array

Then in next loop iteration we have:

- plate = append(plate, menu[i])
  - if length < capacity, it won't allocate a new underlying array - this is exactly what will happen because of the trim previously
  - it will mutate the original underlying array by overwriting the value of the last element in array
  - now all slice headers which point to this underlying array are affected, including the one we stored in result

Slices stored in the `result` now no longer reflect the values which they had when we appended them to the result - instead
they reflect the last mutation that occurred onto their underlying arrays.

Solution is to create a deep copy of the `plate` slice before appending to the `result`.

Correct ways to copy:

1) Using append:

plateCopy := append([]string{}, plate...)
result = append(result, plateCopy)

2) Using Go's builtin copy function:

plateCopy := make([]string, len(plate))
copy(plateCopy, plate)
result = append(result, plateCopy)

⏱ Time Complexity:
- there are 2^n subsets in the worst case.
- but here we restrict to subsets of size <= limit (e.g. 2).
- each level of recursion can spawn O(n) branches.
  - because we never revisit earlier element
  - more precisely, each level spawns (n - depth) branches -> actual number of branches decreases as depth increases
  - but for big O estimates, we often use the maximum per level to approximate the total

- you go up to limit levels.
- so total work is O(n^limit) — i.e., polynomial in n, exponential in limit.

📦 Space Complexity:
- O(limit) stack depth at most, plus result storage O(n^limit) for all subsets.

🔷 Other notes:
- Closures are a clean way for holding shared state like `result` without needing to pass pointer parameters or
using tricky logic to accumulate the result via recursive function return.
*/
func main() {
	menu := []string{"Salad", "Chicken", "Rice"}
	limit := 2 // max number of dishes per combination

	result := getDishCombos(menu, limit)
	// Print final results
	fmt.Printf("\nResult:\n\n")
	for _, combo := range result {
		fmt.Println(combo)
	}
}

func getDishCombos(menu []string, limit int) [][]string {
	var result [][]string

	// Recursive backtracking function
	var backtrack func(start int, plate []string)
	backtrack = func(start int, plate []string) {
		// Print current state at this level
		fmt.Printf("Plate now: %v\n", plate)

		// If we've reached the size limit, stop recursion and save the current plate
		if len(plate) <= limit {
			// Make a copy before storing
			plateCopy := make([]string, len(plate))
			copy(plateCopy, plate)
			result = append(result, plateCopy)
		}

		if len(plate) == limit { // Don't go deeper if we hit the dish limit (prune this path)
			return
		}

		// Try adding each dish starting from 'start' index
		for i := start; i < len(menu); i++ {
			fmt.Printf("  Adding %s\n", menu[i])
			plate = append(plate, menu[i]) // Make a decision
			backtrack(i+1, plate)          // Recurse (go deeper), try adding the next element to the plate
			fmt.Printf("  Removing %s (backtrack)\n", plate[len(plate)-1])
			plate = plate[:len(plate)-1] // Undo the decision
		}
	}

	backtrack(0, nil)

	return result
}
