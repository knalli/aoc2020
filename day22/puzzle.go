package day22

import (
	"crypto/sha256"
	"fmt"
	"github.com/knalli/aoc"
	"hash"
	"sort"
	"strings"
)

func solve1(lines []string) error {
	_, score := playGame1(parseCardStacks(lines), false)
	aoc.PrintSolution(fmt.Sprintf("The winning player's score is %d", score))
	return nil
}

func solve2(lines []string) error {
	_, score := playGame2(parseCardStacks(lines), 0, false)
	aoc.PrintSolution(fmt.Sprintf("The winning player's score is %d", score))
	return nil
}

func playGame1(queues []*aoc.Queue, debug bool) (int, int) {
	return playGame(queues, 0, 0, false, debug, sha256.New())
}

func playGame2(queues []*aoc.Queue, game int, debug bool) (int, int) {
	return playGame(queues, game, game+1, true, debug, sha256.New())
}

func playGame(queues []*aoc.Queue, game int, nextGame int, recursive bool, debug bool, h hash.Hash) (int, int) {

	hashes := make([]string, 0)

	queueToList := func(q *aoc.Queue) []string {
		result := make([]string, 0)
		for i := q.Len(); i > 0; i-- {
			v := q.Head()
			result = append(result, fmt.Sprintf("%d", v.(int)))
			q.Add(v)
		}
		return result
	}

	hashRound := func(queues []*aoc.Queue) string {
		s := ""
		for p, q := range queues {
			clone := q.Clone()
			s += fmt.Sprintf(";%d=", p)
			for !clone.IsEmpty() {
				s += fmt.Sprintf("%d,", clone.Head())
			}
		}
		h.Reset()
		h.Write([]byte(s))
		return fmt.Sprintf("%x", h.Sum(nil))
	}

	copyDeckToCache := func(queues []*aoc.Queue) {
		hashes = append(hashes, hashRound(queues))
	}

	existAlreadyInDeck := func(queues []*aoc.Queue) bool {
		search := hashRound(queues)
		for _, h := range hashes {
			if h == search {
				return true
			}
		}
		return false
	}

	debugMsg := func(m string) {
		if debug {
			fmt.Println(m)
		}
	}

	ifDebugMsg := func(f func() string) {
		if debug {
			fmt.Println(f())
		}
	}

	debugMsg(fmt.Sprintf("=== Game %d ===\n", game+1))

	round := 0
	for {

		playersLeft := 0
		for _, q := range queues {
			if !q.IsEmpty() {
				playersLeft++
			}
		}

		if existAlreadyInDeck(queues) {
			return 0, 0
		}
		copyDeckToCache(queues)

		// Finish?
		if playersLeft == 1 {
			score := 0
			winner := -1
			for p := range queues {
				if !queues[p].IsEmpty() {
					winner = p
				}
			}
			for !queues[winner].IsEmpty() {
				lvl := queues[winner].Len()
				v := queues[winner].Head().(int)
				score += lvl * v
			}
			debugMsg(fmt.Sprintf("The winner of game %d is player %d!", game+1, winner+1))
			return winner, score
		}

		debugMsg(fmt.Sprintf("-- Round %d (Game %d) --", round+1, game+1))
		if debug {
			for player, q := range queues {
				ifDebugMsg(func() string {
					return fmt.Sprintf("Player %d's deck: %s", player+1, strings.Join(queueToList(q), ", "))
				})
			}
		}
		deck := make(map[int]int)
		for player, q := range queues {
			if q.Len() > 0 {
				v := q.Head().(int)
				debugMsg(fmt.Sprintf("Player %d plays: %d", player+1, v))
				deck[player] = v
			}
		}

		{
			// Get winner
			winner := -1
			deckValues := make([]int, 0)

			// recursive
			if recursive {
				r := true
				for p, v := range deck {
					if queues[p].Len() < v {
						r = false
						break
					}
				}
				if r {
					debugMsg("Playing a sub-game to determine the winner..")
					clones := make([]*aoc.Queue, len(queues))
					for p, q := range queues {
						t := q.Clone()
						clones[p] = aoc.NewQueue()
						for i := 0; i < deck[p]; i++ {
							clones[p].Add(t.Head())
						}
					}
					winner, _ = playGame(clones, nextGame, nextGame+1, recursive, debug, h)
					nextGame++
					debugMsg(fmt.Sprintf("The winner of game %d is player %d!", game+2, winner+1))
					debugMsg(fmt.Sprintf("... anyway, back to game %d.\n", game+1))

					queues[winner].Add(deck[winner])
					for p, v := range deck {
						if p != winner {
							queues[winner].Add(v)
						}
					}
				}
			}

			if winner == -1 {
				winnerValue := -1
				for p, v := range deck {
					if v > winnerValue {
						winner = p
						winnerValue = v
					}
				}
				debugMsg(fmt.Sprintf("Player %d wins round %d of game %d!", winner+1, round+1, game+1))

				// only required if not already chosen
				for _, v := range deck {
					deckValues = append(deckValues, v)
				}
				sort.Sort(sort.Reverse(sort.IntSlice(deckValues)))
				for _, v := range deckValues {
					queues[winner].Add(v)
				}
			}
			debugMsg("")
		}

		round++
	}
}

func parseCardStacks(lines []string) []*aoc.Queue {
	result := make([]*aoc.Queue, 0)

	var queue *aoc.Queue
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Index(line, "Player") == 0 {
			queue = aoc.NewQueue()
			result = append(result, queue)
		} else if queue != nil {
			queue.Add(aoc.ParseInt(line))
		}
	}
	return result
}
