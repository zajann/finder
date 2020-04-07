package finder

import (
	"net/url"
	"regexp"
	"strings"
)

const URLRexp string = `[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

func URL(str string) []string {
	r := regexp.MustCompile(URLRexp)
	list := r.FindAllString(str, -1)

	var filteredURL []string
	for _, l := range list {
		u, err := url.Parse("http://" + l)
		if err != nil {
			continue
		}
		parts := strings.Split(u.Hostname(), ".")
		if _, ok := AllowExtension[strings.ToUpper(parts[len(parts)-1])]; ok {
			filteredURL = append(filteredURL, l)
		}
	}
	return filteredURL
}
