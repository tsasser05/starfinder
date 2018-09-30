package dieroll

import (
	"time"
	"fmt"
	"math/rand"
)

/*
* roll (min, max, num)
*
* Rolls a number of times equal to num.  Each result is
* between min and max.
*
* Returns either the number if only one die was rolled
* or the sum of all dice
*
*/

// min, max, number of dice

func Roll (min int, max int, num int) int {
	sum := 0
	rand.Seed(time.Now().Unix())
	
	if num == 1 {
		sum := rand.Intn(max - min) + 1		
	} else if num > 1 {
		count := 0
		
		for count < num {
			new_num = rand.Intn(max - min) + 1
			sum += new_num
			count += 1
		} // for
	} else {
		fmt.Println("dieroll.roll: Incorrect number of dice passed to function.")
	} // if

	return sum
	
} // roll()
