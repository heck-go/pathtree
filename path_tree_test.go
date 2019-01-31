package pathtree

import (
	"fmt"
	"testing"
)

func TestPathTree_Add(t *testing.T) {
	pathTree := NewPathTree()
	pathTree.Add("/", 0, []string{"*"}, map[string]string{})
	pathTree.Add("/api", 1, []string{"*"}, map[string]string{})
	pathTree.Add("/api/book/create", 11, []string{"*"}, map[string]string{})
	pathTree.Add("/api/book/:id", 12, []string{"*"}, map[string]string{})
	pathTree.Add("/api/book/:id/edit", 13, []string{"*"}, map[string]string{})
	pathTree.Add("/api/book/:id/*", 14, []string{"*"}, map[string]string{})
	pathTree.Add("/api/user/create", 21, []string{"*"}, map[string]string{})
	pathTree.Add("/api/user/:id", 22, []string{"*"}, map[string]string{})
	pathTree.Add("/api/user/:id/edit", 23, []string{"*"}, map[string]string{})
	pathTree.Add("/api/user/:id/street/:street", 24, []string{"*"}, map[string]string{})
	
	fmt.Println(pathTree.Match(PathToSegments("/"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/book/create"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/book/123"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/book/123/edit"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/book/123/author/123"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/user/create"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/user/123"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/user/123/edit"), "GET"))
	fmt.Println(pathTree.Match(PathToSegments("api/user/123/street/123"), "GET"))
}
