package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultFounders, err := gocompany.GetFounders("7030")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultFounders)
	os.Exit(0)
}
