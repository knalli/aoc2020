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
	aoc.PrintDayHeader(1, "Day ${day}")
	if err := step1(args); err != nil {
		return err
	}
	if err := step2(args); err != nil {
		return err
	}
	return nil
}

func step1(args []string) error {
	aoc.PrintStepHeader(1)
	if lines, err := aoc.ReadFileToArray("day${id}/puzzle1.txt"); err != nil {
		return err
	} else {
		return solve1(lines)
	}
}

func step2(args []string) error {
	aoc.PrintStepHeader(2)
	if lines, err := aoc.ReadFileToArray("day${id}/puzzle1.txt"); err != nil {
		return err
	} else {
		return solve2(lines)
	}
}

EOF
  okecho "Day $day created at 'day${id}/init.go'"
  cat <<EOF >"day${id}/puzzle.go"
package day${id}

func solve1(lines []string) error {
	return nil
}

func solve2(lines []string) error {
	return nil
}

EOF
  okecho "Day $day created at 'day${id}/puzzle.go'"
  touch "day${id}/sample1.txt" && okecho "Day $day created at 'day${id}/sample1.txt'"
  touch "day${id}/puzzle1.txt" && okecho "Day $day created at 'day${id}/puzzle1.txt'"
  # Patch main.go
  local search
  local adding
  search=$(grep '//_ "github.com/knalli/aoc2020/dayXX"' main.go)
  adding=$(echo "$search" | sed 's/\/\///g')
  adding=$(echo "$adding" | sed -e "s/XX/$id/g")
  if [[ "$OSTYPE" == "darwin"* ]] || [[ "$OSTYPE" == "freebsd"* ]]; then
    sed -i '' "s#$search#$adding\n$search#" main.go
  else
    sed -i "s#$search#$adding\n$search#" main.go
  fi
  return 0
}

if ! initDay "$1"; then
  errecho "Failed initializing day"
  exit 1
fi
