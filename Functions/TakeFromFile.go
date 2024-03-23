package Web_Art

import (
	"bufio"
	"os"
)

var PrintStr string

func TakeFromFile(LineNumber int, BannaerName string) (string, string) {

	file, err := os.Open(TakeBanner(BannaerName))

	if err != nil {
		return result, "error bannaer not found"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0

	for scanner.Scan() {
		counter++

		if counter == LineNumber {
			PrintStr = scanner.Text()
			break
		}
	}
	result = result + PrintStr
	return result, ""
}
