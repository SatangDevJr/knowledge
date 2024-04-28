package permutations

func Permutations(stringInput string) []string {
	if len(stringInput) <= 1 {
		return []string{stringInput} // if len less than one don't need to process next just back there but change to slice
	}

	var output []string

	for index, charactor := range stringInput {
		strackCharactor := Permutations(stringInput[:index] + stringInput[index+1:]) // call itself to keep strack charactor
		for _, strack := range strackCharactor {
			count := 0
			charactorOutput := string(charactor) + strack
			for _, value := range output {
				if value == charactorOutput {
					count++
				}
			}
			if count <= 0 {
				output = append(output, charactorOutput)
			}
		}
	}

	return output
}
