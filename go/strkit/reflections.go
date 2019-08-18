package strkit

import (
  "path/filepath"
  "reflect"
  "runtime"
  "strings"
)

func FuncNameOnly(f interface{}) string {
  funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
  funcName = filepath.Ext(funcName)
  return strings.TrimPrefix(funcName, ".")
}
