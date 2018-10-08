package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tsasser05/starfinder/dieroll"
)

func main() {
	max_str := os.Args[1]
	num_str := os.Args[2]

	max, _ := strconv.Atoi(max_str)
	num, _ := strconv.Atoi(num_str)

	argCount := len(os.Args[1:])

	if argCount != 2 {
		fmt.Println("Incorrect number of args.  Call this program with max, num")
		os.Exit(3)
	} // if

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 1: dieroll.Roll()")
	fmt.Println("-------------------------------------------------------")
	roll1 := dieroll.Roll(max, 1)
	fmt.Println("The roll1  was:  ", roll1, "\n")

	roll2 := dieroll.Roll(max, 2)
	fmt.Println("The roll2 total for two rolls with max =", max, " was:  ", roll2, "\n")

	roll3 := dieroll.Roll(max, num)
	fmt.Println("max = ", max, "num = ", num, "\tSum of rolls:  ", roll3)

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 2: dieroll.RollStat()")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("The stat rolled was:  ", dieroll.RollStat())

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 3: dieroll.RandomMap()")
	fmt.Println("-------------------------------------------------------")

	races := map[string]string{
		"android":  "Artificial people with mechanical components, formerly built as servants but now recognized as citizens.",
		"human":    "Extremely versatile and adaptable race that’s constantly expanding and exploring.",
		"kasatha":  "Four-armed race from a distant desert world with a highly traditional culture.",
		"lashunta": "Charismatic and telepathic race of scholars with two subspecies: one tall and lean, the other short and muscular.",
		"shirren":  "Insectile race that broke away from a locustlike hive, community-minded but addicted to individual choice.",
		"vesk":     "Warlike reptilian race that recently declared a truce with the others—for now.",
		"ysoki":    "Also called “ratfolk,” these short, furry scavengers make up for their short size with big personalities.",
	} // concept

	fmt.Println("The random key is:  ", dieroll.RandomKey(races))

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 4: dieroll.RandomMap() -> Verification of randomness of keys")
	fmt.Println("-------------------------------------------------------")

	counts_by_race := make(map[string]int)

	for cbr := range races {
		counts_by_race[cbr] = 0

	} // for

	fmt.Println("Verify that all races are initialized to zero.")

	for key, val := range counts_by_race {
		fmt.Printf("\t%s = %d\n", key, val)
	} // for

	fmt.Println("\nVerify that race distribution:")

	race_count := 0

	for race_count < 100 {
		element := dieroll.RandomKey(races)
		counts_by_race[element]++
		race_count++
	} // for

	for key, val := range counts_by_race {
		fmt.Printf("\t%s = %d\n", key, val)
	} // for

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("End testing")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("\n")

} // main()
