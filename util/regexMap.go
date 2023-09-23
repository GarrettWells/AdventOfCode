package util

import "regexp"

func CreateMap(input string, regex *regexp.Regexp) map[string]string {
	match := regex.FindStringSubmatch(input)

	if match == nil {
		return nil
	}

	result := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
