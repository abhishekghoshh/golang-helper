package builder

import (
	"design-pattern/src/creational/builder/car"
	"fmt"
)

func Main() {
	car := car.Builder().TopSpeed(50).Paint("blue").Build()
	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
