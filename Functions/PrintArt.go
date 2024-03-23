package Web_Art

import (
	"strings"
)

func PrintArt(str, BannaerName string) (string, string) {
	err := ""

	result = ""
	if str == "\\n" {
		result = result + "\n"

	} else {
		Sentence := strings.Split(str, "\n")

		for i := 0; i < len(Sentence); i++ {

			if Sentence[i] == "" {

				if !(i+1 == len(Sentence)) || !(Sentence[(len(Sentence)-1)] == "") {
					result = result + "\n"
				}

			} else {
				
				result, err = PrintLintByLine(Sentence[i], BannaerName)

			}

		}
	}
	
	
	if Output != "" {
	

		if strings.HasSuffix(Output, ".txt") {
		} else {
			Output = Output + ".txt"
		}

	}
	return result, err
}
