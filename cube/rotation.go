package cube

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

// Rotation describes the rotation of an object in the world. It holds a yaw
// (r[0]) and pitch value (r[1]). Yaw is in the range (-180, 180) while pitch is
// in the range (-90, 90). A positive pitch implies an entity is looking
// downwards, while a negative pitch implies it is looking upwards.
type Rotation [2]float32

// Yaw returns the yaw of r (r[0]).
func (r Rotation) Yaw() float32 {
	return r[0]
}

// Pitch returns the pitch of r (r[1]).
func (r Rotation) Pitch() float32 {
	return r[1]
}

// Elem extracts the elements of the Rotation for direct value assignment.
func (r Rotation) Elem() (yaw, pitch float32) {
	return r[0], r[1]
}

// Add adds the values of two Rotations element-wise and returns a new Rotation.
// If the yaw or pitch would otherwise exceed their respective range as
// described, they are 'overflown' to the other end of the allowed range.
func (r Rotation) Add(r2 Rotation) Rotation {
	return Rotation{r[0] + r2[0], r[1] + r2[1]}.fix()
}

// Opposite returns the Rotation opposite r, so that
// r.Vec3().Add(r.Opposite().Vec3()).Len() is equal to 0.
func (r Rotation) Opposite() Rotation {
	return Rotation{r[0] + 180, -r[1]}.fix()
}

// Direction returns the horizontal Direction that r points towards based on the
// yaw of r.
func (r Rotation) Direction() Direction {
	yaw := r.fix().Yaw()
	switch {
	case yaw > 45 && yaw <= 135:
		return West
	case yaw > -45 && yaw <= 45:
		return South
	case yaw > -135 && yaw <= -45:
		return East
	case yaw <= -135 || yaw > 135:
		return North
	}
	return 0
}

// Orientation returns an Orientation value that most closely matches the yaw
// of r.
func (r Rotation) Orientation() Orientation {
	const step = 360 / 16.0

	yaw := r.fix().Yaw()
	if yaw < -step/2 {
		yaw += 360
	}
	return Orientation(math32.Round(yaw / step))
}

// Vec3 returns the direction vector of r. The length of the mgl32.Vec3 returned
// is always 1.
func (r Rotation) Vec3() mgl32.Vec3 {
	yaw, pitch := r.fix().Elem()
	yawRad, pitchRad := mgl32.DegToRad(yaw), mgl32.DegToRad(pitch)

	m := math32.Cos(pitchRad)
	return mgl32.Vec3{
		-m * math32.Sin(yawRad),
		-math32.Sin(pitchRad),
		m * math32.Cos(yawRad),
	}
}

// fix 'overflows' the Rotation's values to make sure they are within the range
// as described above.
func (r Rotation) fix() Rotation {
	return Rotation{
		math32.Mod(r[0]+180, 360) - 180,
		math32.Mod(r[1]+90, 180) - 90,
	}
}
