package strkit_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  "github.com/Liquid-Labs/strkit/go/strkit"
)

func TestJustFuncName(t *testing.T) {
  assert.Equal(t, `FuncNameOnly`, strkit.FuncNameOnly(strkit.FuncNameOnly))
}
