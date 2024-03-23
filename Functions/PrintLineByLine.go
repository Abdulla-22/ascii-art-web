package Web_Art

func PrintLintByLine(str, BannaerName string) (string, string) {
	err := ""

	for row := 0; row < 8; row++ {

		for columns := 0; columns < len(str); columns++ {
			character := str[columns]
			result, err = TakeFromFile((Line_num(character) + row), BannaerName)
			if err != "" {
				return "", err
			}

		}
		result = result + "\n"
	}


	//Take new line after printing the art
	return result, ""

}
