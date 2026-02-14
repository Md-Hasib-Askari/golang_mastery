package main

import "fmt"

func iteration() {
	// 1️⃣ Traditional For Loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	/*
			for initialization; condition; post {
		    	// code
			}
	*/

	// 2️⃣ For Loop as a While Loop
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// 3️⃣ Alternative Condition-Based Loop
	j := 0
	for j != 10 {
		fmt.Print(j*j, " ")
		j++
	}

	/*
		for condition {} 		→ while
		for {}					→ infinite loop
		for init; cond; post {} → classic loop

		One keyword. Multiple behaviors.
	*/

	// 4️⃣ continue
	/*
		break 	 → exit loop completely
		continue → skip current iteration
	*/

	// 5️⃣ The range Keyword
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value:", v)
	}

	// 6️⃣ Ignoring Values
	for _, v := range aSlice {
		fmt.Println("value:", v)
	}

	// Important Rule Reminder
	/*
		Even if your loop has one line:
		You MUST use curly braces {}.
	*/
}
