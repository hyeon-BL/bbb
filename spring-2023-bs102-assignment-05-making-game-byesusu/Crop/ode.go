package main

import (
	"fmt"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/termbox"
)

// go get github.com/mum4k/termdash
// go get github.com/mum4k/termdash/terminal/termbox

func gauge() {
	// Set initial conditions
	x := 3.0 // crop size
	t := 5.0 // time
	k := 0.3 // constant for growth rate
	L := 1.0 // max size

	// Set the time step.
	dt := t / 4

	// Create a terminal.
	t, err := termbox.New()
	if err != nil {
		fmt.Println("Error creating terminal:", err)
		return
	}
	defer t.Close()

	// Create a gauge widget.
	g, err := cell.NewGauge(cell.Fill())
	if err != nil {
		fmt.Println("Error creating gauge:", err)
		return
	}

	// Create a container for the gauge widget.
	c, err := container.New(t, container.Border())
	if err != nil {
		fmt.Println("Error creating container:", err)
		return
	}
	defer c.Close()

	// Add the gauge to the container.
	if err := c.Update("Crop Growth", g); err != nil {
		fmt.Println("Error updating container:", err)
		return
	}

	// Simulate crop growth.
	for i := 0; i < 1000; i++ {
		// Evaluate the ODE.
		dx := k * x * (1.0 - x/L)

		// Update the size of the crop.
		x += dx * dt

		// Set the gauge value.
		if err := g.Percent(int(x * 100)); err != nil {
			fmt.Println("Error setting gauge value:", err)
			return
		}

		// Refresh the terminal.
		if err := c.Refresh(); err != nil {
			fmt.Println("Error refreshing terminal:", err)
			return
		}

		// Sleep for a short duration to visualize the growth.
		time.Sleep(100 * time.Millisecond)

		// Update the time.
		t += dt
	}
}

func main() {
	gauge()
}
