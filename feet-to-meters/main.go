package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fpdevil/learngo/feet-to-meters/helper"
)

func main() {

	if os.Args != nil && len(os.Args) > 1 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if arg1 == "meter" {
			// convert from string argument to 64 bit float
			feet, _ := strconv.ParseFloat(arg2, 64)
			// meters := feet * 0.3048
			meters := helper.Feet2Meter(feet)

			fmt.Printf("%g feet is %f meters\n", feet, meters)
			fmt.Printf("%g feet is %g meters\n", feet, meters)
		} else if arg1 == "farenheit" {
			celsius, _ := strconv.ParseFloat(arg2, 64)
			// celsius to farenheit
			farenheit := helper.Celsius2Farenheit(celsius)

			fmt.Printf("%g°C celsius is %f°F farenheit\n", celsius, farenheit)
			fmt.Printf("%g°C celsius is %g°F farenheit\n", celsius, farenheit)
		} else {
			arg1 = ""
			fmt.Println("unable to decrypt...")
		}
	} else {
		fmt.Printf("go run %q [feet | celcius]", os.Args[0])
	}
}
