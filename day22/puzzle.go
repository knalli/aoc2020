package day22

import (
	"crypto/sha256"
	"fmt"
	"github.com/knalli/aoc"
	"sort"
	"strings"
)

func solve1(lines []string) error {
	queues := parseCardStacks(lines)
	_, score := playGame(queues, 0, false, false)
	aoc.PrintSolution(fmt.Sprintf("The winning player's score is %d", score))
	return nil
}

func solve2(lines []string) error {
	queues := parseCardStacks(lines)
	_, score := playGame(queues, 0, true, false)
	aoc.PrintSolution(fmt.Sprintf("The winning player's score is %d", score))
	return nil
}

func playGame(queues []*aoc.Queue, game int, recursive bool, debug bool) (int, int) {

	hashes := make([]string, 0)

	/*
		cloneQueue := func(q *aoc.Queue) *aoc.Queue {
			clone := q.Clone()
			for i := q.Len(); i > 0; i-- {
				v := q.Head()
				clone.Add(v)
				q.Add(v)
			}
			return clone
		}*/

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
		h := sha256.New()
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
					winner, _ = playGame(clones, game+1, recursive, debug)
					debugMsg(fmt.Sprintf("The winner of game %d is player %d!", game+2, winner+1))
					debugMsg(fmt.Sprintf("... anyway, back to game %d.", game+1))

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
