package main

import "fmt"

func main() {

	a := -1

	if a >= 0 {
		fmt.Println(a, "is non negative")
	} else {
		fmt.Println(a, "is negative")
	}

	// same with switch
	switch {
	case a >= 0:
		fmt.Println(a, "is non negative")
	default:
		fmt.Println(a, "is negative")
	}

	rank := 11
	switch rank {
	case 1:
		fmt.Println("Rank 1")
	case 2:
		fmt.Println("Rank 2")
	case 3:
		fmt.Println("Rank 3")
	default:
		fmt.Println("Not top 3, but ranks will not indicate overall hardwork or talent!")
	}

	rank = 2
	switch rank {
	case 1, 2, 3:
		fmt.Println("Top Ranker,Secured ", rank)
	default:
		fmt.Println("Not top 3, but ranks will not indicate overall hardwork or talent!")
	}

	rank = 3
	switch rank {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("Top 3")
	default:
		fmt.Println("Not top 3, but ranks will not indicate overall hardwork or talent!")
	}

	switch rank := 4; rank {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		if rank == 4 {
			fmt.Println("just missed Top 3")
			break
		}
		fmt.Println("Top 3")
	default:
		fmt.Println("Not top 3, but ranks will not indicate overall hardwork or talent!")
	}

	var value interface{} = true
	switch value.(type) {
	case string:
		fmt.Println(value, "is string type")
	case int:
		fmt.Println(value, "is int type")
	default:
		fmt.Println(value, "is neither string nor int type")
	}

}

/*OUTPUT
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run switch.go
-1 is negative
-1 is negative
Not top 3, but ranks will not indicate overall hardwork or talent!
Top Ranker,Secured  2
Top 3
just missed Top 3
true is neither string nor int type
*/
