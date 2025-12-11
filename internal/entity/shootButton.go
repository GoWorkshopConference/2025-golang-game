package entity

import (
	"image/color"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/samber/lo"
)

var (
	shootButtonMargin    = 20.0
	shootButtonRadius    = 40.0
	shootButtonCenterPos = lo.T2(
		internal.WindowWidth-controllerButtonCenterPos.A,
		controllerButtonCenterPos.B)
)

var _ Entity = &ShootButton{}

type ShootButton struct {
	centerPos     lo.Tuple2[float64, float64]
	radius        float64
	isTouched     bool
	touchPos      lo.Tuple2[float64, float64]
	lastTouchTime *time.Time
}

type ShootButtonTouchEvent struct{}

func NewShootButton() *ShootButton {
	return &ShootButton{
		centerPos:     shootButtonCenterPos,
		radius:        shootButtonRadius,
		isTouched:     false,
		lastTouchTime: nil,
	}
}

func NewShootButtonTouchEvent() *ShootButtonTouchEvent {
	return &ShootButtonTouchEvent{}
}

func (b *ShootButton) Update() *ShootButtonTouchEvent {
	touchIDs := ebiten.AppendTouchIDs([]ebiten.TouchID{})

	hitFlag := false
	for _, touchID := range touchIDs {
		x, y := ebiten.TouchPosition(touchID)
		hit, _ := CircleHit(b.centerPos, b.radius+touchCircleMargin, float64(x), float64(y))
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
		return NewShootButtonTouchEvent()
	}

	return nil
}

func (b *ShootButton) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen,
		float32(b.centerPos.A), float32(b.centerPos.B),
		float32(b.radius), color.RGBA{0x66, 0x66, 0x66, 0x66}, true)

	return
}

func (b *ShootButton) GetShape() *Shape {
	return nil
}

func (b *ShootButton) GetHitBox() *Shape {
	return nil
}
