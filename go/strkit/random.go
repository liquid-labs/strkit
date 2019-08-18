package strkit

import (
  "math/rand"
  "strings"
  "time"
)

type runeSource struct {
  letterBytes []byte
  noLetters   uint
  idxBits     uint
  idxMask     int64
  idxMax      uint
}

func DefineRuneSource(runes []byte) runeSource {
  var idxBits = uint(1)
  var maxCount = uint(idxBits * 2)
  for maxCount < uint(len(runes)) {
    idxBits += 1
    maxCount *= 2
  }
  return runeSource{
    letterBytes : runes,
    noLetters   : uint(len(runes)),
    idxBits     : idxBits,
    idxMask     : 1<<idxBits - 1,
    idxMax      : 63 / idxBits,
  }
}

var Letters, LettersAndNumbers, FriendlyPassword runeSource
var randSrc rand.Source
func init() {
  Letters = DefineRuneSource([]byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
  LettersAndNumbers = DefineRuneSource([]byte(`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
  // remove easily confused characters: 0, 1, I, l, and O
  FriendlyPassword = DefineRuneSource([]byte(`23456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ_-.+`))

  randSrc = rand.NewSource(time.Now().UnixNano())
}

// RandString produces a random string of letters and numbers. Credit for the implementation goes to [izca](https://stackoverflow.com/users/1705598/icza) from StackOverflow. Refer [here](https://stackoverflow.com/a/31832326/929494) for a detailed breakdown. Since we use dynamic rune sources, our implementation is not quite as fast.
func RandString(rs runeSource, n int) string {
  sb := strings.Builder{}
  sb.Grow(n)
  // A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
  for i, cache, remain := n-1, randSrc.Int63(), rs.idxMax; i >= 0; {
    if remain == 0 {
      cache, remain = randSrc.Int63(), rs.idxMax
    }
    if idx := uint(cache & rs.idxMask); idx < rs.noLetters {
      sb.WriteByte(rs.letterBytes[idx])
      i--
    }
    cache >>= rs.idxBits
    remain--
  }

  return sb.String()
}
