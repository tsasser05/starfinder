package dieroll

import (
	"time"
	"fmt"
	"math/rand"
	"os"
)


/*******************************************************
*
* contains(valid, max)
*
* Search the valid array for the max.
*
*
*******************************************************/

func contains(valid [7]int, max int) bool {
	for _, element := range valid {
		if element == max {
			return true
		} // if
	} // for
	
	return false

} // contains()


/*******************************************************
*
* init()
*
*
*******************************************************/

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

} // init


/*******************************************************
*
* randomInt(max)
*
*******************************************************/

func randomInt(max int) int {
	min := 1
	return min + rand.Intn(max - min) 

} // randomInt()


/*******************************************************
* roll (max, num)
*
* Rolls a number of times equal to num.  Each result is
* between 1 and max.
*
* Returns either the number if only one die was rolled
* or the sum of all dice
*
* Max must be equal to 4, 6, 8, 10, 12, 20 or 100.  The 
* function will handle the details for rand.Intn itself.
*
* TBD:
*
* Figure out better error handling
*
*******************************************************/

func Roll (max int, num int) int {
	sum := 0
	debug := 1
	
	// max must be 5 for d4, 7 for d6, 9 for d8, 11 for d10, 13 for d12 and 21 for d20
        // max += 1
	// Limited to the following:
	var valid = [7]int{4, 6, 8, 10, 12, 20, 100} 

	valid_check := contains(valid, max)

	if valid_check == false {
		fmt.Println("dieroll().Roll(): max value = ", max, "is not correct value.")
		os.Exit(3)
	} // if

	// Increment max by one for randomInt calculation.
	max += 1
	
	if num == 1 {
		sum = randomInt(max) 

		if debug > 0 {
			fmt.Println("dieroll.Roll() number rolled is ", sum, "\n")
		} // if

	} else if num > 1 {
		count := 0
		
		for count < num {
			new_num := randomInt(max)

			if debug > 0 {
				fmt.Println("dieroll.Roll() number rolled is ", new_num, "\n")
			} // if
			
			sum += new_num
			count += 1

			if debug > 0 {
				fmt.Println("dieroll.Roll() sum is ", sum, "\n")
			} // debug
			
		} // for

		if debug > 0 {
			fmt.Println("dieroll.Roll() final sum is ", sum, "\n")
		} // debug
		
	} else {
		fmt.Println("dieroll.Roll: Incorrect number of dice passed to function: ", num, "\n")
		os.Exit(3)
	} // if

	return sum
	
} // roll()


/*******************************************************
*
* RollStat()
*
* Rolls 4d6 for a stat and takes the top 3
* 
*******************************************************/

//func RollStat() int {
//	rolls := [...]int{}

//} // RollStat()
	

