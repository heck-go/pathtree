package pathtree

import "regexp"

type pair struct {
	regexp *regexp.Regexp
	subtree *SubTree
}

type SubTree struct {
	// Values for exact match
	value map[string]interface{}

	// Value for glob path segment
	globValue map[string]interface{}

	// Subtree for fixed path segment
	fixed map[string]*SubTree

	// Subtree for variable path segment
	varSubTree *SubTree

	// Subtree for regular expression
	regexp []pair
}

func NewSubTree() *SubTree {
	return &SubTree{
		value: make(map[string]interface{}),
		globValue: make(map[string]interface{}),
		fixed: make(map[string]*SubTree),
	}
}

func (self *SubTree) HasVarPaths() bool  {
	if len(self.regexp) != 0 {
		return true
	}
	if self.varSubTree != nil {
		return true
	}
	if len(self.globValue) != 0 {
		return true
	}
	return false
}