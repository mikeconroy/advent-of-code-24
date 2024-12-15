package main

import (
	"flag"
	"fmt"
	"github.com/mikeconroy/advent-of-code-24/day9"
	"github.com/mikeconroy/advent-of-code-24/day10"
	"github.com/mikeconroy/advent-of-code-24/day11"
	"github.com/mikeconroy/advent-of-code-24/day12"
	"github.com/mikeconroy/advent-of-code-24/day13"
)

func main() {
	dayToRun := flag.Int("day", 0, "Which Day to run? Defaults to 0 (all)")
	flag.Parse()

	days := []func() (string, string){
		day9.Run,
		day10.Run,
		day11.Run,
		day12.Run,
		day13.Run,
	}

	if *dayToRun == 0 {
		for day, run := range days {
			runDay(day+9, run)
		}
	} else {
		runDay(*dayToRun, days[*dayToRun-9])
	}
}

func runDay(dayNum int, run func() (string, string)) {
	p1, p2 := run()
	fmt.Printf("Day %d\n\tP1: %s\n\tP2: %s\n", dayNum, p1, p2)
}
