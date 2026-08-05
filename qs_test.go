package qs

import "testing"

func TestParse(t *testing.T) {
	got := Parse("?a=1&b=hello%20world&flag=")
	if got["a"] != "1" {
		t.Errorf("a = %q, want 1", got["a"])
	}
	if got["b"] != "hello world" {
		t.Errorf("b = %q, want decoded spaces", got["b"])
	}
	if _, ok := got["flag"]; !ok {
		t.Error("a valueless key should still be present")
	}
}

func TestBuildRoundTrip(t *testing.T) {
	in := map[string]string{"q": "go & rust", "page": "2"}
	got := Parse(Build(in))
	for k, v := range in {
		if got[k] != v {
			t.Errorf("round trip of %q = %q, want %q", k, got[k], v)
		}
	}
}

func TestBuildIsStable(t *testing.T) {
	in := map[string]string{"b": "2", "a": "1"}
	if got := Build(in); got != "a=1&b=2" {
		t.Errorf("Build = %q, want a=1&b=2", got)
	}
}
