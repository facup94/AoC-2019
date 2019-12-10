package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	t.Run("12", func(t *testing.T) {
		mass := 12

		got := calculateFuelPart1(mass)
		want := 2

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("14", func(t *testing.T) {
		mass := 14

		got := calculateFuelPart1(mass)
		want := 2

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("1969", func(t *testing.T) {
		mass := 1969

		got := calculateFuelPart1(mass)
		want := 654

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("100756", func(t *testing.T) {
		mass := 100756

		got := calculateFuelPart1(mass)
		want := 33583

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("part 2 - 14", func(t *testing.T) {
		mass := 14

		got := calculateFuelPart2(mass)
		want := 2

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("part 2 - 1969", func(t *testing.T) {
		mass := 1969

		got := calculateFuelPart2(mass)
		want := 966

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("part 2 - 100756", func(t *testing.T) {
		mass := 100756

		got := calculateFuelPart2(mass)
		want := 50346

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
