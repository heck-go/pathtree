package pathtree

import "strings"

// Splits path into path segments
func PathToSegments(path string) []string {
	paths := strings.Split(path, "/")
	if len(paths) == 0 {
		return  paths
	}
	if len(paths[0]) == 0 {
		paths = paths[1:]
	}
	if len(paths) == 0 {
		return  paths
	}
	if len(paths[len(paths) - 1]) == 0 {
		paths = paths[:len(paths) - 1]
	}

	ret := make([]string, 0, len(paths))

	for _, v := range paths {
		if len(v) != 0 {
			ret = append(ret, v)
		}
	}

	return ret
}
