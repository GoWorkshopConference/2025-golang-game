package entity

import (
	"image/color"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/samber/lo"
)

var (
	controllerButtonMargin    = 20.0
	controllerButtonRadius    = 60.0
	controllerButtonCenterPos = lo.T2(controllerButtonMargin+controllerButtonRadius, internal.WindowHeight-(controllerButtonMargin+controllerButtonRadius))
	touchInterval             = time.Duration(300 * time.Millisecond)

	subCircleRadius   = 30.0
	touchCircleMargin = 20.0
)

var _ Entity = &ControllerButton{}

type ControllerButton struct {
	centerPos     lo.Tuple2[float64, float64]
	radius        float64
	isTouched     bool
	touchPos      lo.Tuple2[float64, float64]
	lastTouchTime *time.Time
}

type ControllerButtonTouchEvent struct {
	Direction lo.Tuple2[float64, float64]
}

var (
	buttonAnchor = 20
	buttonSize   = 40.0
)

func NewControllerButton() *ControllerButton {
	return &ControllerButton{
		centerPos:     controllerButtonCenterPos,
		radius:        controllerButtonRadius,
		isTouched:     false,
		lastTouchTime: nil,
	}
}

func calculateDirection(
	touchPos lo.Tuple2[float64, float64],
	centerPos lo.Tuple2[float64, float64],
) lo.Tuple2[float64, float64] {
	direction := lo.T2(
		touchPos.A-centerPos.A,
		touchPos.B-centerPos.B,
	)
	direction = lo.T2(
		direction.A/controllerButtonRadius,
		direction.B/controllerButtonRadius,
	)

	if direction.A > 1 {
		direction.A = 1
	}
	if direction.B > 1 {
		direction.B = 1
	}
	if direction.A < -1 {
		direction.A = -1
	}
	if direction.B < -1 {
		direction.B = -1
	}

	return direction
}

func NewControllerButtonTouchEvent(
	touchPos lo.Tuple2[float64, float64],
	centerPos lo.Tuple2[float64, float64],
) *ControllerButtonTouchEvent {

	return &ControllerButtonTouchEvent{
		Direction: calculateDirection(touchPos, centerPos),
	}
}

func (b *ControllerButton) Update() *ControllerButtonTouchEvent {
	touchIDs := ebiten.AppendTouchIDs([]ebiten.TouchID{})

	hitFlag := false
	for _, touchID := range touchIDs {
		x, y := ebiten.TouchPosition(touchID)
		hit, _ := utils.CircleHit(b.centerPos, b.radius+touchCircleMargin, float64(x), float64(y))
		if hit {
			b.touchPos = lo.T2(float64(x), float64(y))
			hitFlag = true
			break
		}
	}

	if hitFlag && !b.isTouched {
		b.isTouched = true
	}

	if !hitFlag {
		if b.lastTouchTime == nil {
			b.lastTouchTime = lo.ToPtr(time.Now())
		}
		if b.lastTouchTime != nil && time.Since(*b.lastTouchTime) >= touchInterval {
			b.lastTouchTime = nil
			b.isTouched = false
		}
	}

	if b.isTouched {
		return NewControllerButtonTouchEvent(b.touchPos, b.centerPos)
	}

	return nil
}

func (b *ControllerButton) Draw(screen *ebiten.Image) {
	col := color.RGBA{0x66, 0x66, 0x66, 0x66}

	vector.FillCircle(screen,
		float32(b.centerPos.A), float32(b.centerPos.B),
		float32(b.radius), col, true)

	if b.isTouched {
		vector.FillCircle(screen,
			float32(b.touchPos.A), float32(b.touchPos.B),
			float32(subCircleRadius), col, true)
	}

	return
}
