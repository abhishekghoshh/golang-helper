package car

type builder interface {
	TopSpeed(int) builder
	Paint(string) builder
	Build() Vehicle
}

type CarBuilder struct {
	speedOption int
	color       string
}

func (cb *CarBuilder) TopSpeed(speed int) builder {
	cb.speedOption = speed
	return cb
}

func (cb *CarBuilder) Paint(color string) builder {
	cb.color = color
	return cb
}

func (cb *CarBuilder) Build() Vehicle {
	return &Car{
		topSpeed: cb.speedOption,
		color:    cb.color,
	}
}
func Builder() builder {
	return &CarBuilder{}
}
