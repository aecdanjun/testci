package main

import (
  "testing"
)

func TestAdd(t *testing.T) {
  r := add(1,3)

  if r != 4 {
    t.Fail()
  }
}
