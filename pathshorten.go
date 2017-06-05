package pathshorten

import (
	"os"
	"strings"
)

func Run(path string) string {
	var (
		s     string
		skip  bool
		slash bool
	)

	pos := 0
	sep := string(os.PathSeparator)
	count := strings.Count(path, sep)

	if len(path) == 1 && path == "." {
		return path
	}

	if !strings.Contains(path, sep) {
		return path
	}

	for _, p := range strings.Split(path, "") {
		switch p {
		case sep:
			s += p
			skip = false
			slash = true
			pos += 1
		case ".":
			if slash {
				s += p
				skip = false
			}
			slash = false
		default:
			if count == pos {
				s += p
				continue
			}
			if skip {
				continue
			}
			s += p
			skip = true
			slash = false
		}
	}

	return s
}
