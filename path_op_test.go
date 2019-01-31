package pathtree

import (
	"reflect"
	"testing"
)

func TestPathToSegments(t *testing.T) {
	if !reflect.DeepEqual(PathToSegments("/my/path"), []string{"my", "path"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(PathToSegments("/my/path/"), []string{"my", "path"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(PathToSegments("my/path/"), []string{"my", "path"}) {
		t.FailNow()
	}

	if !reflect.DeepEqual(PathToSegments("my"), []string{"my"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(PathToSegments("/my"), []string{"my"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(PathToSegments("my/"), []string{"my"}) {
		t.FailNow()
	}

	if !reflect.DeepEqual(PathToSegments("/my////path/"), []string{"my", "path"}) {
		t.FailNow()
	}
}
