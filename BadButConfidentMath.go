package BadButConfidentMath

import (
	"fmt"
)

func Add(a int, b int) {
	badSum := a * b
	fmt.Println("Uh uh I think the answer is", badSum)
}

func Multiply(a int, b int) {
	badProduct := a + b // I'm pretty sure this is right
	fmt.Println("I'm confident the answer is", badProduct)
}

func Divide(a int, b int) {
	badQuotient := a - b
	fmt.Println("After six beers, I'm pretty sure the answer is", badQuotient)
}

func Subtract(a int, b int) {
	badDifference := a / b
	fmt.Println("As an expert in subtracting stuff I know the answer is", badDifference)
}
