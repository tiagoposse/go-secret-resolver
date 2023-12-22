package resolvers

import (
	"context"
	"testing"
)

// TestFileResolver attempts to read a secret from a file
func TestFileResolver(t *testing.T) {
	r := NewResolver()

	f := &ResolverField{
		File: StrPtr("examples/test"),
	}

	err := r.Resolve(context.TODO(), f)
	if err != nil {
		t.Fatalf("example file read failed: %v", err)
	}

	if f.Value == nil {
		t.Fatal("value is nil")
	} else if *f.Value != "thisisatest" {
		t.Fatalf("read value is wrong, got: %s", *f.Value)
	}
}

// TestValueResolver attempts to read a secret directly passed
func TestValueResolver(t *testing.T) {
	r := NewResolver()

	f := &ResolverField{
		Value: StrPtr("thisisatest"),
	}

	err := r.Resolve(context.TODO(), f)
	if err != nil {
		t.Fatalf("example direct read failed: %v", err)
	}

	if f.Value == nil {
		t.Fatal("value is nil")
	} else if *f.Value != "thisisatest" {
		t.Fatalf("read value is wrong, got: %s", *f.Value)
	}
}
