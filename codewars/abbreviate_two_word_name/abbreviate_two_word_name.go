package abbreviatetwowordname

import "strings"

func AbbrevName(name string) string {
	// example: "Sam Harris" => "S.H"
	var abbrev string
	for i, char := range name {
		if i == 0 || (i > 0 && name[i-1] == ' ') {
			abbrev += strings.ToUpper(string(char)) + "."
		}
	}
	return abbrev[:len(abbrev)-1] // Remove the trailing dot
}
