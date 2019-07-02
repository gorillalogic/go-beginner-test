// Copyright 2019 Gorilla Logic, Inc.
// All rights reserved.
package duplicates

import (
	"testing"
	"bytes"
)

func TestDuplicates(t *testing.T) {
    origin := []byte {1,3,2,3,24,36,2,5,90,122,15,45,32,24}
	expected := []byte {2,3,24}

	result := Duplicates(origin)
	
    if !bytes.Equal(result, expected) {
       t.Errorf("Result was incorrect, got: %v, expected: %v.", result, expected)
    }
}