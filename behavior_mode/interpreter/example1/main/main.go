package main

import (
	"fmt"
	"go-patterns/behavior_mode/interpreter/example1"
)

func getMaleExpression() *example1.OrExpression {
	robert := example1.NewTe("Robert")
	john := example1.NewTe("John")
	return example1.NewOr(robert, john)
}
func getMarriedWomanExpression() *example1.AndExpression {
	julie := example1.NewTe("Julie")
	married := example1.NewTe("Married")
	return example1.NewAnd(julie, married)
}

func main() {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()

	fmt.Println("John is male?", isMale.Interpret("John"))
	fmt.Println("Julie is married?", isMarriedWoman.Interpret("Julie"))
}
