//
//You are given a string composed of only 1s, 0s, and Xs.
//Write a program that will print out every possible combination where you replace the X with  both 0 and 1.
//Example
//$ myprogram X0 00
//10
//$ myprogram 10X10X0 1001000
//1001010
//1011000
//1011010

//The big O notation of this algorithm is O(2^n) where n is the number of Xs in the pattern.
package xoreplace

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type XoReplace struct {
	baseString     string
	numOccurrences int
	parts          []string
}

func New(baseString string) *XoReplace {
	parts := strings.Split(baseString, "X")
	return &XoReplace{
		baseString:     baseString,
		numOccurrences: len(parts) - 1,
		//		numOccurrences: strings.Count(baseString, "X"),
		parts: parts,
	}
}

func (x *XoReplace) PrintFor(i int) {
	formatString := fmt.Sprintf("%%0%db", x.numOccurrences)
	binString := fmt.Sprintf(formatString, i)
	//	log.Println("formatString", formatString, "binString", binString)
	//	letters := strings.Split(binString, "")
	//	var b strings.Builder

	for i := range x.parts {
		fmt.Print(x.parts[i])
		if i < x.numOccurrences {
			fmt.Printf("%s", []byte{binString[i]})
		}
	}
	fmt.Print("\n")
}

func (x *XoReplace) Run() {
	log.Println("Printing out combinations for", x.baseString)
	max := int(math.Pow(2, float64(x.numOccurrences)))
	for i := 0; i < max; i++ {
		x.PrintFor(i)
	}
}

//running time,
//X0 = 2
//0X = 2
// XX0 = 3
// X0X = 3
//for a string with n Xs it writes 2n+1 times.
//So the whole operation should take (2^n+1)*2n+1
