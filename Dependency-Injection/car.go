package dependencyinjection

type Speeder interface {
	MaxSpeed() int
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	currentSpeed := 90
	if c.Speeder.MaxSpeed() < 10 {
		return 15
	}
	
	if currentSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}

	return currentSpeed
}

type DefaultEngine struct{}

func (de DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct{}

func (te TurboEngine) MaxSpeed() int {
	return 100
}
