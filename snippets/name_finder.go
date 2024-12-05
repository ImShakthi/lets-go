package snippets

import "fmt"

var letterNumMap = map[string]int{
	"a": 1,
	"i": 1,
	"j": 1,
	"q": 1,
	"y": 1,

	"b": 2,
	"k": 2,
	"r": 2,

	"c": 3,
	"g": 3,
	"l": 3,
	"s": 3,

	"d": 4,
	"m": 4,
	"t": 4,

	"e": 5,
	"h": 5,
	"n": 5,
	"x": 5,

	"u": 6,
	"v": 6,
	"w": 6,

	"o": 7,
	"z": 7,

	"f": 8,
	"p": 8,

	"A": 1,
	"I": 1,
	"J": 1,
	"Q": 1,
	"Y": 1,

	"B": 2,
	"K": 2,
	"R": 2,

	"C": 3,
	"G": 3,
	"L": 3,
	"S": 3,

	"D": 4,
	"M": 4,
	"T": 4,

	"E": 5,
	"H": 5,
	"N": 5,
	"X": 5,

	"U": 6,
	"V": 6,
	"W": 6,

	"O": 7,
	"Z": 7,

	"F": 8,
	"P": 8,
}

type ATC []int

var allowedTotalCounts = ATC{5, 6}

var suffixes = []string{"", "raj", "kumar"}

func (atc ATC) In(num int) bool {
	for _, i := range atc {
		if i == num {
			return true
		}
	}
	return false
}

func PrintMatchingNames() {

	println("Get matching names!!!")

}

// M and T
// 1 and 5
// 14,21,23,39,41,50,59
// 13,20,22,38,40,49,58

func PrintValidNames(names []string) {
	for _, name := range names {
		for _, suffix := range suffixes {
			n := name + suffix
			IsValidName("k" + n)
			IsValidName("ks" + n)
		}
		println()
	}
}

func IsValidName(name string) bool {
	//fmt.Printf("\nGiven name is %s", name)
	count := 0
	for _, ch := range name {
		//fmt.Printf("%s->%d, ", string(ch), letterNumMap[string(ch)])
		count += letterNumMap[string(ch)]
	}
	//fmt.Printf("\ncount for %s is %d ", name, count)
	//in := allowedTotalCounts.In(count)
	//if in {
	fmt.Printf("\n%s -> %d -> %v", name, count, true)
	//}
	return true
}

func sumOfNumber() int {

	return 0
}
