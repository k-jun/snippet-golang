
package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

func WordCount(phrase string) map[string]int {
	m := make(map[string]int)

	s := strings.Replace(phrase, ",", " ", -1)
	reg := regexp.MustCompile(`[^A-Za-z1-9']`)
	s = reg.ReplaceAllString(s, " ")

	fmt.Println(s)

	for _, w := range strings.Split(s, " ") {
		w = strings.TrimSpace(w)
		w = strings.ToLower(w)
		reg = regexp.MustCompile(`.+'t`)
		if !reg.MatchString(w) {
			w = strings.Replace(w, "'", "", -1)
		}
		if w != "" {
			m[w]++
		}
	}

	return m
}
