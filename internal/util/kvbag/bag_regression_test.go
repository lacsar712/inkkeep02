package kvbag

import "testing"

// Regression: a freshly created Bag must be safe to write to even though
// callers never initialize the backing map themselves. Previously New
// returned a Bag with a nil map, so the first Set panicked with
// "assignment to entry in nil map".
func TestNewBagSetGetRoundtrip(t *testing.T) {
	b := New()

	// Set on a brand-new Bag must not panic.
	b.Set("a", "1")
	b.Set("b", "2")

	for k, want := range map[string]string{"a": "1", "b": "2"} {
		got, ok := b.Get(k)
		if !ok || got != want {
			t.Fatalf("Get(%q)=%q ok=%v, want %q true", k, got, ok, want)
		}
	}

	if got := b.Len(); got != 2 {
		t.Fatalf("Len()=%d, want 2", got)
	}

	// Overwrite an existing key.
	b.Set("a", "9")
	if got, _ := b.Get("a"); got != "9" {
		t.Fatalf("Get(%q)=%q, want %q after overwrite", "a", got, "9")
	}
}
