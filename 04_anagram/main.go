package main

func groupAnagram(strs []string) [][]string {
	result := [][]string{}
	counter := map[[26]int][]string{}

	for _, str := range strs {
		chars := [26]int{}

		for _, s := range str {
			chars[s-'a']++
		}

		counter[chars] = append(counter[chars], str)
	}

	for _, value := range counter {
		result = append(result, value)
	}

	return result
}
