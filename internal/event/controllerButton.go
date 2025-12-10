package event

import "github.com/samber/lo"

type ControllerButtonTouchEvent struct {
	TouchPos lo.Tuple2[float64, float64]
}

func NewControllerButtonTouchEvent(touchPos lo.Tuple2[float64, float64]) *ControllerButtonTouchEvent {
	return &ControllerButtonTouchEvent{
		TouchPos: touchPos,
	}
}

func (e *ControllerButtonTouchEvent) Type() EventType {
	return ControllerButtonTouch
}
