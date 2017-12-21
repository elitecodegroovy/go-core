// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
)

const (
	win            = 100 // The winning score in a game of Pig
	gamesPerSeries = 10  // The number of games per series to simulate
)

// A score includes scores accumulated in previous turns for each player,
// as well as the points scored by the current player in this turn.
type score struct {
	player, opponent, thisTurn int
}

// An action transitions stochastically to a resulting score.
type action func(current score) (result score, turnIsOver bool)

// roll returns the (result, turnIsOver) outcome of simulating a die roll.
// If the roll value is 1, then thisTurn score is abandoned, and the players'
// roles swap.  Otherwise, the roll value is added to thisTurn.
func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// stay returns the (result, turnIsOver) outcome of staying.
// thisTurn score is added to the player's score, and the players' roles swap.
func stay(s score) (score, bool) {
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

// A strategy chooses an action for any given score.
type strategy func(score) action

// stayAtK returns a strategy that rolls until thisTurn is at least k, then stays.
func stayAtK(k int) strategy {
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// play simulates a Pig game and returns the winner (0 or 1).
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	currentPlayer := rand.Intn(2) // Randomly decide who plays first
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

// roundRobin simulates a series of games between every pair of strategies.
func roundRobin(strategies []strategy) ([]int, int) {
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1) // no self play
	return wins, gamesPerStrategy
}

// ratioString takes a list of integer values and returns a string that lists
// each value and its percentage of the sum of all values.
// e.g., ratios(1, 2, 3) = "1/6 (16.7%), 2/6 (33.3%), 3/6 (50.0%)"
func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtK(k + 1)
	}
	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Wins, losses staying at k =% 4d: %s\n",
			k+1, ratioString(wins[k], games-wins[k]))
	}
}

/**
----Game overview----
Pig is a two-player game played with a 6-sided die. Each turn, you may roll or stay.
If you roll a 1, you lose all points for your turn and play passes to your opponent. Any other roll adds its value to your turn score.
If you stay, your turn score is added to your total score, and play passes to your opponent.
The first person to reach 100 total points wins.

The score type stores the scores of the current and opposing players, in addition to the points accumulated during the current turn.
*/

/**
----User-defined function types--
In Go, functions can be passed around just like any other value. A function's type signature describes the types of its arguments and return values.

The action type is a function that takes a score and returns the resulting score and whether the current turn is over.

If the turn is over, the player and opponent fields in the resulting score should be swapped, as it is now the other player's turn.

*/

/**
----Multiple return values--
Go functions can return multiple values.

The functions roll and stay each return a pair of values. They also match the action type signature. These action functions define the rules of Pig.
*/

/**
----Higher-order functions--
A function can use other functions as arguments and return values.

A strategy is a function that takes a score as input and returns an action to perform.
(Remember, an action is itself a function.)

*/

/**
----Function literals and closures--
Anonymous functions can be declared in Go, as in this example. Function literals are closures: they inherit the scope of the function in which they are declared.

One basic strategy in Pig is to continue rolling until you have accumulated at least k points in a turn, and then stay. The argument k is enclosed by this function literal,
 which matches the strategy type signature.
*/

/**
---Simulating games
We simulate a game of Pig by calling an action to update the score until one player reaches 100 points. Each action is selected by calling the strategy function associated with the current player.

*/

/**
----Simulating a tournament
The roundRobin function simulates a tournament and tallies wins. Each strategy plays each other strategy gamesPerSeries times

*/

/**
----Variadic function declarations
Variadic functions like ratioString take a variable number of arguments. These arguments are available as a slice inside the function.

*/

/**
----Simulation results
The main function defines 100 basic strategies, simulates a round robin tournament, and then prints the win/loss record of each strategy.

Among these strategies, staying at 25 is best, but the optimal strategy for Pig is much more complex.
*/
