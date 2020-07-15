package condition

import (
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1=1")
	}
	// if result, err := someFunc(); err == nil {
	// 	t.Log("have not err")
	// } else {
	// 	t.Log("have err")
	// }
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux")
	default:
		t.Logf("%s.", os)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 1, 3:
			t.Log(i, "Odd")
		case 0, 2:
			t.Log(i, "Even")
		default:
			t.Log(i, "it is not 0-3")
		}
	}
}

func TestSwitchConditionCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "Odd")
		default:
			t.Log("unknow")
		}
	}
}
