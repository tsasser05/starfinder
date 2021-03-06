package abilities

import (
	"github.com/tsasser05/starfinder/dieroll"
	"fmt"
) 

type Abilities struct {
	Strength int
	Dexterity int
	Constitution int
	Intelligence int
	Wisdom int
	Charisma int
} // Abilities


func (a *Abilities) Roll() {
	a.Strength = dieroll.RollStat()
	a.Dexterity = dieroll.RollStat()
	a.Constitution = dieroll.RollStat()
	a.Intelligence = dieroll.RollStat()
	a.Wisdom = dieroll.RollStat()
	a.Charisma = dieroll.RollStat()
	
} // Roll()


func (a *Abilities) Display() {
	fmt.Println("Abilities")
	fmt.Printf("    Strength:  %d\n", a.Strength)
	fmt.Printf("    Dexterity:  %d\n", a.Dexterity)
	fmt.Printf("    Constitution:  %d\n", a.Constitution)
	fmt.Printf("    Intelligence:  %d\n", a.Intelligence)
	fmt.Printf("    Wisdom:  %d\n", a.Wisdom)
	fmt.Printf("    Charisma:  %d\n", a.Charisma)

} // display()
