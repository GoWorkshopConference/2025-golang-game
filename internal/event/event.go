package event

type Event interface{}

type EventType string

var (
	ControllerButtonTouch EventType = "CONTROLLER_BUTTON_TOUCH"
	ShootButtonTouch      EventType = "SHOOT_BUTTON_TOUCH"
)
