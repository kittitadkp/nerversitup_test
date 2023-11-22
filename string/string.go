package permutation

import (
	"regexp"
	"sort"
	"strings"
)

func GeneratePermutations(input string) []string {
	var result []string
	if strings.TrimSpace(input) == "" {
		return result
	}

	permute([]rune(input), 0, &result)
	result = removeDuplicates(result)
	sort.Strings(result)
	return result
}

func removeDuplicates(input []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, str := range input {
		if !encountered[str] {
			encountered[str] = true
			result = append(result, str)
		}
	}

	return result
}

func permute(chars []rune, start int, result *[]string) {
	if start == len(chars)-1 {
		// Convert the rune slice to a string and append to the result
		*result = append(*result, string(chars))
		return
	}

	// Recursive permutation
	for i := start; i < len(chars); i++ {
		// Swap characters at positions start and i
		chars[start], chars[i] = chars[i], chars[start]

		// Recursively generate permutations for the remaining characters
		permute(chars, start+1, result)

		// Backtrack by swapping the characters back
		chars[start], chars[i] = chars[i], chars[start]
	}
}

func CountSmileys(s []string) int {
	count := 0

	pattern := `[:;][-~]?[)D]`

	re := regexp.MustCompile(pattern)

	// Check each string in the array against the pattern
	for _, face := range s {
		if re.MatchString(face) {
			count++
		}
	}

	return count
}
