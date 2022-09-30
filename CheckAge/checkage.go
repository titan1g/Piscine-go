package main

import (
	"fmt"
)

func main() {
	fmt.Checkage(nb int)
}

func Checkage (nb int) {
	if nb <= 18 {
		fmt.string("Vous pouvez entrer")
	} else if nb > 18 {
		fmt.string("passez votre chemin")
	}

}




DEUXIEME QUESTION 


package main

import (
	"fmt"
	"github.com/01-edu/z01"
)

func main() {
	fmt.Checkage(nb int)
}

func Isnegative(nb int) {
	if nb <= 18 {
		z01.PrintRune('T')
	} else if nb >18 {
		z01.PrintRune('F')
	}
}
