package ui

import (
  "fmt"
)

const CLEAR = "\033[H\033[2J"
const CLEARLN = "\033[2K"

const (
  WHITE string = "\033[0m"
  RED          = "\033[31m"
  GREEN        = "\033[32m"
  YELLOW       = "\033[33m"
)

func Goto_xy(x, y int) string {
  return fmt.Sprintf("\033[%d;%dH",x,y)
}
