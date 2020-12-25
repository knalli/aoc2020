package day25

import (
	"errors"
	"fmt"
	"github.com/knalli/aoc"
	"math"
)

func solve1(lines []string) error {
	return solve(lines)
}

func solve2(lines []string) error {
	return solve(lines)
}

func solve(lines []string) error {
	cardPubKey, doorPubKey := parse(lines)
	fmt.Printf("card public key = %d, door public key = %d\n", cardPubKey, doorPubKey)
	cardLoopSize := resolveLoopSize(7, cardPubKey)
	doorLoopSize := resolveLoopSize(7, doorPubKey)
	fmt.Printf("card loop size = %d, door loop size = %d\n", cardLoopSize, doorLoopSize)
	cardEncryptionKey := applyLoopSize(1, doorPubKey, cardLoopSize)
	doorEncryptionKey := applyLoopSize(1, cardPubKey, doorLoopSize)
	fmt.Printf("card encryption key = %d, door encryption key = %d\n", cardEncryptionKey, doorEncryptionKey)
	if cardEncryptionKey == doorEncryptionKey {
		aoc.PrintSolution(fmt.Sprintf("encryption key is %d", doorEncryptionKey))
		return nil
	} else {
		return errors.New("invalid encryption key")
	}
}

func parse(lines []string) (cardPubKey int, doorPubKey int) {
	cardPubKey = aoc.ParseInt(lines[0])
	doorPubKey = aoc.ParseInt(lines[1])
	return
}

func resolveLoopSize(s int, n int) int {
	result := 1
	for i := 1; i < math.MaxInt64; i++ {
		result *= s
		result %= 20201227
		if result == n {
			return i
		}
	}
	return -1
}

func applyLoopSize(start int, s int, l int) int {
	result := start
	for i := 0; i < l; i++ {
		result *= s
		result %= 20201227
	}
	return result
}
