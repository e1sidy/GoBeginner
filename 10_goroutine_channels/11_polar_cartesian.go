package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	radius, angle float64
}

type cartesian struct {
	x, y float64
}

const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}
}

func main() {
	polarChan := make(chan polar)
	defer close(polarChan)
	cartesianChan := createSolver(polarChan)
	defer close(cartesianChan)
	interact(polarChan, cartesianChan)
}

func createSolver(in chan polar) chan cartesian {
	out := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-in
			angle := polarCoord.angle * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(angle)
			y := polarCoord.radius * math.Sin(angle)
			out <- cartesian{x, y}
		}
	}()
	return out
}

func interact(out chan polar, in chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and the angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]

		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatsToStrings(numbers)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid number\n")
				continue
			}
			out <- polar{polars[0], polars[1]}
			coord := <-in
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintf(os.Stderr, "invalid number\n")
		}
	}
	fmt.Println()
}

func floatsToStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
