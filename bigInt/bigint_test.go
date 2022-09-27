package bigint

import (
	"testing"
)

func TestNewInt(t *testing.T) {

	a, err := NewInt("-9999999999")
	var want Bigint = Bigint{Value: "-9999999999"}
	if a != want || err == nil {
		t.Errorf("got %q, wanted %q", a, want)
	}

	a.Set("2")
	var wan Bigint = Bigint{Value: "2"}
	if a != wan || err == nil {
		t.Errorf("got %q, wanted %q", a, wan)
	}

	b, err := NewInt("0000000000009999999999")
	var want1 Bigint = Bigint{Value: "9999999999"}

	if b != want1 || err == nil {
		t.Errorf("got %q, wanted %q", b, want1)
	}

	c, err := NewInt("-11111111111111111111111111111111111111111111")
	var want2 Bigint = Bigint{Value: "-11111111111111111111111111111111111111111111"}

	if c != want2 || err == nil {
		t.Errorf("got %q, wanted %q", c, want2)
	}

	d, err := NewInt("+11111111111111111111111111111111111111111111")
	var want3 Bigint = Bigint{Value: "11111111111111111111111111111111111111111111"}

	if d != want3 || err == nil {
		t.Errorf("got %q, wanted %q", d, want3)
	}

}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "Two minus", num1: Bigint{Value: "-999"}, num2: Bigint{Value: "-111"}, expected: "-1110"},
		{name: "Two plus number", num1: Bigint{Value: "+999"}, num2: Bigint{Value: "+1111"}, expected: "2110"},
		{name: "Two plus number", num1: Bigint{Value: "-9999"}, num2: Bigint{Value: "+1111"}, expected: "-8888"},
		{name: "Two plus number", num1: Bigint{Value: "-9999"}, num2: Bigint{Value: "+11111"}, expected: "1112"},
		{name: "Two plus number", num1: Bigint{Value: "+9999"}, num2: Bigint{Value: "-1111"}, expected: "8888"},
		{name: "Two plus number", num1: Bigint{Value: "+9999"}, num2: Bigint{Value: "-11111"}, expected: "-1112"},
		{name: "Two plus number", num1: Bigint{Value: "+58642"}, num2: Bigint{Value: "+235894"}, expected: "294536"},
		{name: "Two minus", num1: Bigint{Value: "-11111111111111111111111111111111111111111111"}, num2: Bigint{Value: "-1"}, expected: "-11111111111111111111111111111111111111111112"},
		{name: "Two plus number", num1: Bigint{Value: "+1"}, num2: Bigint{Value: "+11111111111111111111111111111111111111111111"}, expected: "11111111111111111111111111111111111111111112"},
	}

	for _, test := range tests {
		output := Add(test.num1, test.num2)
		if output.Value != test.expected {
			t.Errorf(" Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "Two minus", num1: Bigint{Value: "-999"}, num2: Bigint{Value: "-111"}, expected: "-888"},
		{name: "Two plus number", num1: Bigint{Value: "+999"}, num2: Bigint{Value: "+111"}, expected: "888"},
		{name: "Two different operator minus plus", num1: Bigint{Value: "-999999999999999"}, num2: Bigint{Value: "+11111111111"}, expected: "-1000011111111110"},
		{name: "Two different operator plus minus", num1: Bigint{Value: "+999999999999999"}, num2: Bigint{Value: "-11111111111"}, expected: "1000011111111110"},
	}

	for _, test := range tests {
		output := Sub(test.num1, test.num2)
		if output.Value != test.expected {
			t.Errorf(" Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "Two minus", num1: Bigint{Value: "-999"}, num2: Bigint{Value: "-111"}, expected: "110889"},
		{name: "Two plus number", num1: Bigint{Value: "+999"}, num2: Bigint{Value: "+111"}, expected: "110889"},
		{name: "Two different operator minus plus", num1: Bigint{Value: "-999999999999999"}, num2: Bigint{Value: "+11111111111"}, expected: "-4611686016279904256"},
		{name: "Two different operator plus minus", num1: Bigint{Value: "+999999999999999"}, num2: Bigint{Value: "-11111111111"}, expected: "-4611686016279904256"},
	}

	for _, test := range tests {
		output := Multiply(test.num1, test.num2)
		if output.Value != test.expected {
			t.Errorf(" Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "Two minus", num1: Bigint{Value: "-999"}, num2: Bigint{Value: "-111"}, expected: "0"},
		{name: "Two plus number", num1: Bigint{Value: "+999"}, num2: Bigint{Value: "+111"}, expected: "0"},
		{name: "Two different operator minus plus", num1: Bigint{Value: "-999999999999999"}, num2: Bigint{Value: "+11111111111"}, expected: "-9999"},
		{name: "Two different operator plus minus", num1: Bigint{Value: "+999999999999999"}, num2: Bigint{Value: "-11111111111"}, expected: "9999"},
	}

	for _, test := range tests {
		output := Mod(test.num1, test.num2)
		if output.Value != test.expected {
			t.Errorf(" Output %q not equal to expected %q", output, test.expected)
		}
	}
}
