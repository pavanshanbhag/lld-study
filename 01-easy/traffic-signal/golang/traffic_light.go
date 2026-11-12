package trafficsignalsystem

type TrafficLight struct {
	ID             string
	CurrentSignal  Signal
	RedDuration    int
	YellowDuration int
	GreenDuration  int
}

func NewTrafficLight(id string, redDuration, yellowDuration, greenDuration int) *TrafficLight {
	return &TrafficLight{
		ID:             id,
		RedDuration:    redDuration,
		YellowDuration: yellowDuration,
		GreenDuration:  greenDuration,
		CurrentSignal:  Red,
	}
}

func (tl *TrafficLight) ChangeSignal(newSignal Signal) {
	tl.CurrentSignal = newSignal
}

func (tl *TrafficLight) Signal() Signal {
	return tl.CurrentSignal
}
