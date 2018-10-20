package class

/*******************************************************
*
* package class
*
* Design will include all modifiers that will be applied
* to the character's stats.  For now, this package will 
* merely include the character's class and description.
*
*******************************************************/

import (
	"github.com/tsasser05/starfinder/dieroll"
	"fmt"
) 

type Class struct {
	Class string
	Description string
} // Class


func (a *Class) Roll() {
	classes := map[string]string{
		"Envoy": "Charismatic people person good at a wide range of skills who inspires allies to accomplish great heroic feats.",
		"Mechanic": "Master of machines and technology whose tinkering produces a drone companion or a powerful brain implant.",
		"Mystic": "Magic user whose mysterious connection to a powerful force grants abilities that break the laws of the universe.",
		"Operative": "Stealthy combatant with wide-ranging know-how who is adept at taking advantage of unprepared foes.",
		"Solarian": "Disciplined warrior whose mastery of the stars grants either a weapon or armor made of stellar power.",
		"Soldier": "Expert with a huge range of armor, guns, and melee weapons who specializes in certain types of gear.",
		"Technomancer": "Magic user who is preternaturally attuned to technology and can use it to unlock powerful effects.",
	} // classes map

	a.Class = dieroll.RandomKey(classes)
	a.Description = classes[a.Class]
	
} // 


func (a *Class) Display() {
	fmt.Println("Class")
	fmt.Println("    Class:  ", a.Class)
	fmt.Println("    Description:  ", a.Description)

} // display()
