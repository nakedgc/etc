
package main

import "fmt" // unused

import (
	"reflect"       // unused
	"math"          // used
	tmplt "text/template" // unused
)

import "flag" // unused

var unusedvar int = 123

var usedvar int = 456

func main() {
	a := 5
	// a unused
	b := 6
	if (b==6) {
		fmt.Println("Hello world!")
		if c:=0; b==6 {
			fmt.Println("Hello galaxy!")
			// c unused
		}
	}
	
	usedvar = 789
	
	fmt.Println(math.Sqrt(float64(usedvar)))
}
