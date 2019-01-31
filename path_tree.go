package pathtree

import (
	"regexp"
)

// PathTree helps in matching a path and tag against a value.
// Use NewPathTree function to construct a PathTree instance.
type PathTree struct {
	subtree *SubTree
}

// NewPathTree creates and returns a new PathTree instance
func NewPathTree() *PathTree {
	return &PathTree{
		subtree: NewSubTree(),
	}
}

// Add adds a path to the PathTree
func (self *PathTree) Add(path string, value interface{}, tags []string, pathRegex map[string]string) {
	segs := PathToSegments(path)

	subtree := self.subtree
	numSegs := len(segs)

	for i, seg := range segs {
		// Is it glob segment?
		if i == numSegs-1 && (seg == "*" || (seg[0] == ':' && seg[len(seg)-1] == '*')) {
			for _, tag := range tags {
				subtree.globValue[tag] = value
			}
			return
		}

		// Named segment?
		if seg[0] == ':' {
			segName := seg[1:]
			if _, found := pathRegex[segName]; found {
				next := NewSubTree()
				re, err := regexp.Compile(pathRegex[segName])
				if err != nil {
					panic("Invalid regular expression provided!")
				}
				subtree.regexp = append(subtree.regexp, pair{regexp: re, subtree: next})
				subtree = next
			} else {
				next := subtree.varSubTree
				if next == nil {
					next = NewSubTree()
					subtree.varSubTree = next
				}
				subtree = next
			}
			continue
		}

		// Fixed segment
		next, segFound := subtree.fixed[seg]
		if !segFound {
			next = NewSubTree()
			subtree.fixed[seg] = next
		}
		subtree = next
	}

	for _, tag := range tags {
		subtree.value[tag] = value
	}
}

// Match returns the value that matches the given path segments and tag. If a match
// for path segments and tag are not found, nil is returned.
func (self *PathTree) Match(segments []string, tag string) interface{} {
	return  self._match(self.subtree, segments, tag)
}

func (self *PathTree) _match(root *SubTree, segments []string, tag string) interface{} {
	subtree := root
	
	for i, seg := range segments {
		next := subtree.fixed[seg]
		if subtree.HasVarPaths() {
			if next != nil {
				ret := self._match(next, segments[i+1:], tag)
				if ret != nil {
					return  ret
				}
				next = nil
			}
			
			// Look if there is a regexp match
			for _, p := range subtree.regexp {
				if p.regexp.MatchString(seg) {
					if next != nil {	// Multiple regex matches
						ret := self._matchRegex(subtree, segments[i:], tag)
						if ret != nil {
							return  ret
						}
						next = nil
						break
					}
					next = p.subtree
				}
			}
			
			// Look for named segment
			if next == nil {
				if subtree.varSubTree != nil {
					if subtree.globValue != nil {
						ret := self._match(subtree.varSubTree, segments[i+1:], tag)
						if ret != nil {
							return  ret
						}
						ret = subtree.globValue[tag]
						if ret != nil {
							return ret
						}
						ret = subtree.globValue["*"]
						if ret != nil {
							return ret
						}
					}
					next = subtree.varSubTree
				}
				if next == nil {
					ret := subtree.globValue[tag]
					if ret != nil {
						return ret
					}
					ret = subtree.globValue["*"]
					if ret != nil {
						return ret
					}
					return nil
				}
			}
		}
		
		if next == nil {
			return  nil
		}
		subtree = next
	}
	
	ret := subtree.value[tag]
	if ret != nil {
		return ret
	}
	ret = subtree.globValue[tag]
	if ret != nil {
		return ret
	}
	ret = subtree.value["*"]
	if ret != nil {
		return ret
	}
	ret = subtree.globValue["*"]
	if ret != nil {
		return ret
	}
	return nil
}

func (self *PathTree) _matchRegex(root *SubTree, segments []string, tag string) interface{} {
	for _, p := range root.regexp {
		if p.regexp.MatchString(segments[0]) {
				ret := self._match(p.subtree, segments[1:], tag)
				if ret != nil {
					return  ret
				}
		}
	}
	return nil
}