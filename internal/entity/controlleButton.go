package entity

import (
	"image/color"
	"math"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/samber/lo"
)

var (
	controllerButtonMargin    = 10.0
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
	touchID       *ebiten.TouchID
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

	nowTouched := false
	for _, touchID := range touchIDs {
		x, y := ebiten.TouchPosition(touchID)
		hit, _ := CircleHit(b.centerPos, b.radius+touchCircleMargin, float64(x), float64(y))
		if hit {
			b.touchID = lo.ToPtr(touchID)
			b.touchPos = lo.T2(float64(x), float64(y))
			nowTouched = true
			b.isTouched = true
			break
		}
	}

	nowOutOfCircleTouched := false
	if !nowTouched && b.touchID != nil && lo.Contains(touchIDs, *b.touchID) {
		nowOutOfCircleTouched = true

		x, y := ebiten.TouchPosition(*b.touchID)
		touchPos := lo.T2(float64(x), float64(y))

		// centerPos からの方向ベクトルを計算
		dx := touchPos.A - b.centerPos.A
		dy := touchPos.B - b.centerPos.B

		// 距離を計算
		distance := math.Sqrt(dx*dx + dy*dy)

		// 半径 b.radius+touchCircleMargin の位置を計算
		maxRadius := b.radius + touchCircleMargin
		if distance > 0 {
			// 方向ベクトルを正規化して、最大半径を掛ける
			normalizedDx := dx / distance
			normalizedDy := dy / distance
			b.touchPos = lo.T2(
				b.centerPos.A+normalizedDx*maxRadius,
				b.centerPos.B+normalizedDy*maxRadius,
			)
		} else {
			// 距離が0の場合は中心位置を使用
			b.touchPos = b.centerPos
		}
	}

	if nowTouched || nowOutOfCircleTouched {
		return NewControllerButtonTouchEvent(b.touchPos, b.centerPos)
	}

	b.isTouched = false
	b.touchID = nil

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

func (b *ControllerButton) GetShape() *Shape {
	return nil
}

func (b *ControllerButton) GetHitBox() *Shape {
	return nil
}
