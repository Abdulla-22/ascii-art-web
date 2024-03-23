package Web_Art

import "strings"

func TakeBanner(str string) string {

	// The ( ../ ) has been used to go bace to the previce folder
	// then use ( FontTypes/Font-Name ) to navegait to the file.

	BannerMap := map[string]string{

		"standard":   "../FontTypes/standard.txt",
		"shadow":     "../FontTypes/shadow.txt",
		"thinkertoy": "../FontTypes/thinkertoy.txt",
	}

	if str == "" {
		return BannerMap["standard"]

	}

	return BannerMap[strings.TrimSuffix(str, ".txt")]

}
