#!/bin/bash

FMT_RED="\033[0;31m"
FMT_GREEN="\033[0;32m"
FMT_RESET="\033[0m"

function errecho() {
  echo >&2 -ne "${FMT_RED}"
  echo >&2 -n "$@"
  echo >&2 -e "${FMT_RESET}"
}

function okecho() {
  echo >&2 -ne "${FMT_GREEN}"
  echo >&2 -n "$@"
  echo >&2 -e "${FMT_RESET}"
}

function initDay() {
  local day
  local id
  day="$1"
  id=$(printf "%02d" "$day")
  if [ -d "day${id}" ]; then
    errecho "day directory already exist"
    return 1
  fi
  if [ -r "day${id}/init.go" ]; then
    errecho "day file already exist"
    return 1
  fi
  mkdir "day${id}"
  cat <<EOF >"day${id}/init.go"
package day${id}

import (
	"github.com/knalli/aoc"
)

func init() {
	aoc.Registry.Register(${day}, main)
}

func main(args []string) error {
	aoc.PrintDayHeader(${day}, "Day${day}")
	return nil
}

EOF
  okecho "Day $day created at 'day${id}/init.go'"
  return 0
}

if ! initDay "$1"; then
  errecho "Failed initializing day"
  exit 1
fi
