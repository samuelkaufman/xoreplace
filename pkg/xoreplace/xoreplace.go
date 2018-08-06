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

//The big O notation of this algorithm is O(2^n) where n is the number of Xs in the pattern, ie each additional X roughly doubles the max runtime.
package xoreplace

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type XoReplace struct {
	baseString     string
	numOccurrences int
	parts          []string
	outFh          io.Writer
}

func New(baseString string, outFh io.Writer) *XoReplace {
	if outFh == nil {
		outFh = os.Stdout
	}
	parts := strings.Split(baseString, "X")
	return &XoReplace{
		baseString:     baseString,
		numOccurrences: len(parts) - 1,
		outFh:          outFh,
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
		fmt.Fprint(x.outFh, x.parts[i])
		if i < x.numOccurrences {
			fmt.Fprintf(x.outFh, "%s", []byte{binString[i]})
		}
	}
	fmt.Fprint(x.outFh, "\n")
}

func (x *XoReplace) Run() {
	max := int(math.Pow(2, float64(x.numOccurrences)))
	for i := 0; i < max; i++ {
		x.PrintFor(i)
	}
}
