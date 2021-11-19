package main

import (
	"fmt"

	"github.com/rifannurmuhammad/design-pattern/creational_pattern/builder"
	"github.com/rifannurmuhammad/design-pattern/creational_pattern/factory"
)

func main() {
	exampleFactory()
	fmt.Println("==================")
	exampleBuilder()
}

func exampleFactory() {
	var resto factory.Resto
	resto = &factory.Lemonilo{}

	food := resto.Cook()
	food.Order()
}

func exampleBuilder() {
	w := builder.WorkspaceBuilder{}
	myworkspace, err := w.BuildMonitor(1).BuildKeyboard(1).BuildTable().CreateWorkspace()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myworkspace)
	fmt.Println("done")
}
