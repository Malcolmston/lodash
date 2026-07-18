package lodash

import "testing"

func TestParseInt(t *testing.T) {
	cases := []struct {
		s    string
		base int
		want int64
		err  bool
	}{
		{"42", 10, 42, false},
		{"  -7 ", 10, -7, false},
		{"ff", 16, 255, false},
		{"0x1A", 0, 26, false},
		{"0b101", 0, 5, false},
		{"101", 2, 5, false},
		{"nope", 10, 0, true},
	}
	for _, c := range cases {
		got, err := ParseInt(c.s, c.base)
		if c.err {
			if err == nil {
				t.Errorf("ParseInt(%q,%d) expected error", c.s, c.base)
			}
			continue
		}
		if err != nil || got != c.want {
			t.Errorf("ParseInt(%q,%d) = %d,%v want %d", c.s, c.base, got, err, c.want)
		}
	}
}

func TestParseFloat(t *testing.T) {
	got, err := ParseFloat("  3.14 ")
	if err != nil || got != 3.14 {
		t.Errorf("ParseFloat = %v,%v", got, err)
	}
	if _, err := ParseFloat("x"); err == nil {
		t.Error("ParseFloat(x) expected error")
	}
}
