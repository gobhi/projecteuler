package problem_test

import (
  "testing"
  "github.com/gobhi/projecteuler/problem"
)

func TestMultiplesOfThreeAndFiveUnder1000(t *testing.T) {
  sum := problem.MultiplesOfThreeAndFive()
  if sum != 233168 {
    t.Errorf("Problem 1: The sum of all the multiples of 3 or 5 below 1000 is 233168, got %d.\n", sum)
  }
}