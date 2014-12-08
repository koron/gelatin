package args

import (
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	err := Root.Parse()
	if err != nil {
		t.Errorf("shold be success: %v", err)
	}
}

func TestUnknownOptions(t *testing.T) {
	err := Root.Parse("--foo")
	if err == nil {
		t.Errorf("shold be failure for --foo")
	}
	if ert := reflect.TypeOf(err); ert != reflect.TypeOf(ErrorUnknownOption{}) {
		t.Errorf("%v is not ErrorUnknownOption", ert)
	}
	if err.Error() != "unknown option:--foo in mode:(global)" {
		t.Errorf("error message mismatch: %#v", err)
	}
}
