package strkit_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  "github.com/Liquid-Labs/strkit/go/strkit"
)

func TestRandStringLetters(t *testing.T) {
  minSize := 16
  for i := 0; i < 10; i += 1 {
    randString := strkit.RandString(strkit.Letters, minSize + i)
    assert.Equal(t, minSize + i, len(randString))
    assert.Regexp(t, `^[a-zA-Z]`, randString)
  }
}
