package snippets

import (
	"strings"
)

var letterToCodeMap = map[string]string{
	"A": "._",
	"B": "_...",
	"C": "_._.",
	"D": "_..",
	"E": ".",
	"F": ".._.",
	"G": "__.",
	"H": "....",
	"I": "..",
	"J": ".___",
	"K": "_._",
	"L": "._..",
	"M": "__",
	"N": "_.",
	"O": "___",
	"P": ".__.",
	"Q": "__._",
	"R": "._.",
	"S": "...",
	"T": "_",
	"U": ".._",
	"V": "..._",
	"W": ".__",
	"X": "_.._",
	"Y": "_.__",
	"Z": "__..",
}

var codeToLetterMap = map[string]string{
	"._":   "A",
	"_...": "B",
	"_._.": "C",
	"_..":  "D",
	".":    "E",
	".._.": "F",
	"__.":  "G",
	"....": "H",
	"..":   "I",
	".___": "J",
	"_._":  "K",
	"._..": "L",
	"__":   "M",
	"_.":   "N",
	"___":  "O",
	".__.": "P",
	"__._": "Q",
	"._.":  "R",
	"...":  "S",
	"_":    "T",
	".._":  "U",
	"..._": "V",
	".__":  "W",
	"_.._": "X",
	"_.__": "Y",
	"__..": "Z",
}

func encodeToMorseCode(input string) string {
	var mc []string
	for i := 0; i < len(input); i++ {
		mc = append(mc, letterToCodeMap[string(input[i])])
	}
	return strings.Join(mc, " ")
}

func decodeMorseCode(input string) string {
	var output []string
	codes := strings.Split(input, " ")
	for i := 0; i < len(codes); i++ {
		output = append(output, codeToLetterMap[codes[i]])
	}

	return strings.Join(output, "")
}
