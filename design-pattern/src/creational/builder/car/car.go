package car

import "strconv"

type Vehicle interface {
	Drive() string
	Stop() string
}

type Car struct {
	topSpeed int
	color    string
}

func (c *Car) Drive() string {
	return "Driving at speed: " + strconv.Itoa(c.topSpeed)
}

func (c *Car) Stop() string {
	return "Stopping a " + c.color + " car"
}
