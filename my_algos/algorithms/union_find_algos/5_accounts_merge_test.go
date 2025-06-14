/*
ACCOUNTS MERGE – LeetCode 721 (Union-Find)

🧠 PROBLEM:
You're given a list of accounts where each account is a list:
- The first element is the person's name
- The rest are email addresses belonging to that person

Merge accounts that belong to the same person. Two accounts belong to the same person if they share at least one email.
Return the merged accounts — name followed by all associated emails (sorted).

✅ IDENTITY IS BASED ON EMAILS, NOT NAMES!
Even if two names differ (e.g. "John" vs "Jon"), we merge their accounts if they share an email(s).
Account/person name doesn't play any role in the logic.

🔑 INSIGHT:
This is a classic use case for Union-Find:
- Each email is a node
- If two emails appear in the same account, we union them
- After processing all accounts, connected components represent the same person

🧱 DATA STRUCTURES:
- UnionFind<string>: Email is used directly as the key (avoids mapping to integer IDs)
- emailToName[email] = name: Tracks the name associated with each email
- rootToEmails[rootEmail] = []emails: Groups all emails by their root

📌 WHY GENERIC UNION-FIND IS BETTER:
- Avoids email <-> ID translation (less code, fewer bugs)
- Emails can be unioned directly
- Cleaner, more readable

🌀 DUPLICATES:
- Same email may appear in multiple accounts — that's expected
- emailToName just needs to record one mapping per email

🎯 REPRESENTATIVE:
- The root email (from Find) is effectively used as the group ID
- It's not important which email becomes the root — we only care that all emails in the group are connected

📊 TIME COMPLEXITY:
- O(N * α(N)) for Union-Find operations, where N = total number of emails
- O(N log N) to sort email groups

📦 SPACE COMPLEXITY:
- O(N) for maps in Union-Find structure

👀 VISUAL:
Account 1: ["John", "a@mail", "b@mail"]
Account 2: ["John", "b@mail", "c@mail"]
Account 3: ["John", "d@mail"]

UF Connections:

"a"--"b"--"c"         "d"
⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯    ⎯⎯⎯⎯⎯⎯⎯⎯
Group 1: [a, b, c]    Group 2: [d]

Merged output: ["John", a, b, c] and ["John", d]
*/

package union_find_algos

import (
	"fmt"
	"panca.com/algo/union_find"
	"sort"
	"testing"
)

func mergeAccounts(accounts [][]string) [][]string {
	uf := union_find.NewUnionFind[string]()
	emailToName := make(map[string]string) // maps every email to account/person name

	// First pass: union emails and assign names
	for _, account := range accounts {
		name := account[0]
		emails := account[1:]

		emailToName[emails[0]] = name

		// union all emails to the first one
		for _, email := range emails[1:] {
			uf.Union(emails[0], email)
			emailToName[email] = name
		}
	}

	// Second pass: group emails by their representative/root
	rootToEmails := make(map[string][]string) // root -> [root, other emails...]
	for email := range emailToName {
		root := uf.Find(email)
		rootToEmails[root] = append(rootToEmails[root], email)
	}

	// Build result
	var res [][]string
	for root, emails := range rootToEmails {
		sort.Strings(emails)
		name := emailToName[root] // or we could do: emailToName[emails[0]] - all emails in group belong to same person so it doesn't matter
		res = append(res, append([]string{name}, emails...))
	}

	return res
}

func Test_mergeAccounts(t *testing.T) {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"}, // johnsmith@mail.com from here and johnsmith@mail.com from third John will end up in the same group
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, // john_newyork@mail.com will be grouped with johnsmith@mail.com
		{"Mary", "mary@mail.com"},
	} // thus we will form a group of [johnsmith@mail.com, john00@mail.com, john_newyork@mail.com]

	result := mergeAccounts(accounts)
	for _, account := range result {
		fmt.Println(account)
	}

	/*
		Expected groups:
		- {"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}
		- {"John", "johnnybravo@mail.com"}
		- {"Mary", "mary@mail.com"}
	*/

	// Additional test: conflicting names (still works)
	accounts2 := [][]string{
		{"Jon", "x@mail.com", "y@mail.com"},
		{"Jonathan", "y@mail.com", "z@mail.com"},
	}
	result2 := mergeAccounts(accounts2)
	fmt.Println("\nConflicting name input:")
	for _, account := range result2 {
		fmt.Println(account)
	}
	// The output will use the name attached to the root email
}
