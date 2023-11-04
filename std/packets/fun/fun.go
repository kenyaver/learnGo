package fun

import (
  "errors"
)

func Sum(x, y int) int {
  return x + y
}

func Div(x, y int) (int, error) {
  if y == 0 {
    return 0, errors.New("0 низя")
  }
  return x / y, nil
}