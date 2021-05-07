package main

import (
	"math"
)

type vec3 struct {
	e []float64
}

func (v *vec3) vec3(x float64, y float64, z float64) {
	e := []float64{x, y, z}
	v.e = e
}

func (v *vec3) x() float64 { return v.e[0] }
func (v *vec3) y() float64 { return v.e[1] }
func (v *vec3) z() float64 { return v.e[2] }
func (v *vec3) r() float64 { return v.e[0] }
func (v *vec3) g() float64 { return v.e[1] }
func (v *vec3) b() float64 { return v.e[2] }

func (v *vec3) plus(v2 vec3) vec3 {
	rtn := vec3{}
	rtn.vec3(v.e[0]+v2.e[0], v.e[1]+v2.e[1], v.e[2]+v2.e[2])
	return rtn
}

func (v *vec3) minus(v2 vec3) vec3 {
	rtn := vec3{}
	rtn.vec3(v.e[0]-v2.e[0], v.e[1]-v2.e[1], v.e[2]-v2.e[2])

	return rtn
}

func (v *vec3) mult(v2 vec3) vec3 {
	return vec3{
		e: []float64{v.e[0] * v2.e[0], v.e[1] * v2.e[1], v.e[2] * v2.e[2]},
	}
}

func (v *vec3) div(v2 vec3) vec3 {
	return vec3{
		e: []float64{v.e[0] / v2.e[0], v.e[1] / v2.e[1], v.e[2] / v2.e[2]},
	}
}

func (v *vec3) multiplyT(t float64) vec3 {
	rtn := vec3{}
	rtn.vec3(t*v.e[0], t*v.e[1], t*v.e[2])

	return rtn
}

func (v *vec3) divideT(t float64) vec3 {
	rtn := vec3{}
	rtn.vec3(v.e[0]/t, v.e[1]/t, v.e[2]/t)
	return rtn
}

func (v *vec3) negate() {
	v.e[0] = v.e[0] - v.e[0] - v.e[0]
	v.e[1] = v.e[1] - v.e[1] - v.e[1]
	v.e[2] = v.e[2] - v.e[2] - v.e[2]
}
func (v *vec3) index(i int) float64 { return v.e[i] }

//Maybe not needed
func (v *vec3) indexP(i int) *float64 { return &v.e[i] }

// +=
func (v *vec3) plusEq(v2 vec3) {
	v.e[0] = v.e[0] + v2.e[0]
	v.e[1] = v.e[1] + v2.e[1]
	v.e[2] = v.e[2] + v2.e[2]
}

// -=
func (v *vec3) minEq(v2 vec3) {
	v.e[0] = v.e[0] - v2.e[0]
	v.e[1] = v.e[1] - v2.e[1]
	v.e[2] = v.e[2] - v2.e[2]
}

// *=
func (v *vec3) multEq(v2 vec3) {
	v.e[0] = v.e[0] * v2.e[0]
	v.e[1] = v.e[1] * v2.e[1]
	v.e[2] = v.e[2] * v2.e[2]
}

// /=
func (v *vec3) divEq(v2 vec3) {
	v.e[0] = v.e[0] / v2.e[0]
	v.e[1] = v.e[1] / v2.e[1]
	v.e[2] = v.e[2] / v2.e[2]
}

// *=
func (v *vec3) multEqf(t float64, v2 vec3) {
	v.e[0] = t * v2.e[0]
	v.e[1] = t * v2.e[1]
	v.e[2] = t * v2.e[2]
}

// /=
func (v *vec3) divEqfv(t float64) {
	v.e[0] = v.e[0] / t
	v.e[1] = v.e[1] / t
	v.e[2] = v.e[2] / t
}

func (v *vec3) length() float64 {
	return float64(math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]))
}

func (v *vec3) squaredLength() float64 {
	return float64(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

func (v *vec3) makeUnitVector() {
	k := float64(1.0 / math.Sqrt(v.e[0]*v.e[0]+v.e[1]*v.e[1]+v.e[2]*v.e[2]))
	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k
}

func (v *vec3) unitVector() vec3 {

	length := v.length()
	rtn := v.divideT(length)
	return rtn
}

func unitVector(v vec3) vec3 {
	length := v.length()
	rtn := v.divideT(length)
	return rtn
}

func dot(v1 vec3, v2 vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}
