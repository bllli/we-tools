package playground

import (
	"fmt"
	"testing"
)

func TestEmptyStruct(t *testing.T) {
	var null struct{}

	m := map[string]struct{}{
		"a": null,
		"b": null,
	}
	if _, ok := m["a"]; !ok {
		t.Error("a is not in the map")
	}
	if _, ok := m["c"]; ok {
		t.Error("c is in the map")
	}

}

type C struct {
	intValue int
}

// C intValue setter
func (c *C) SetIntValueWithPointerReceiver(intValue int) {
	c.intValue = intValue
}

// C intValue getter
func (c *C) GetIntValue() int {
	return c.intValue
}

func (c C) SetIntValueWithoutPointerReceiver(intValue int) {
	c.intValue = intValue
	fmt.Printf("C.SetIntValueWithoutPointerReceiver: %p\n", &c) // 能观察到本函数里的c被拷贝了一份 改了也没用
}

func TestTheClass(t *testing.T) {
	c := C{
		intValue: 123,
	}
	fmt.Printf("TestTheClass: %p\n", &c)
	C.SetIntValueWithoutPointerReceiver(c, 100) // 拷贝了 所以没有改变c的值
	if c.GetIntValue() != 123 {
		t.Error("c.GetIntValue() != 123")
	}
	c.SetIntValueWithoutPointerReceiver(100)
	if c.GetIntValue() != 123 {
		t.Error("c.GetIntValue() != 123")
	}

	(&c).SetIntValueWithPointerReceiver(200)
	if c.GetIntValue() != 200 {
		t.Error("c.GetIntValue() != 200")
	}
	c.intValue = 123
	c.SetIntValueWithPointerReceiver(200)
	if c.GetIntValue() != 200 {
		t.Error("c.GetIntValue() != 200")
	}

}
