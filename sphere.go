package main

import (
	"fmt"
	"math"
)

type sphere struct {
	center vec3
	radius float64
	mat    material
}

func (s *sphere) getCenter() vec3 {
	return s.center
}

func (s *sphere) sphere(cen vec3, r float64, m material) {
	s.center = cen
	s.radius = r
	s.mat = m
	fmt.Println(m.getName())
}

func (s *sphere) hitFunc(r ray, tMin float64, tMax float64, rec hitRecord) (bool, hitRecord) {

	oc := r.origin()
	oc = oc.minus(s.center)
	a := dot(r.direction(), r.direction())
	b := dot(oc, r.direction())
	c := dot(oc, oc) - s.radius*s.radius
	discriminant := (b*b) - (a*c)

	rec.matPtr = s.mat

	if discriminant > 0 {
		tmp := (-b - math.Sqrt(b*b-a*c)) / a
		if tmp < tMax && tmp > tMin {
			rec.t = tmp
			rec.p = r.pointAtParameter(rec.t)
			nml := rec.p.minus(s.center)
			rslt := nml.divideT(s.radius)
			rec.normal = rslt

			return true, rec
		}
		tmp = (-b + math.Sqrt(b*b-a*c)) / a
		if tmp < tMax && tmp > tMin {
			rec.t = tmp
			rec.p = r.pointAtParameter(rec.t)

			nml := rec.p.minus(s.center)
			rec.normal = nml.divideT(s.radius)

			return true, rec
		}
	}
	return false, rec
}
