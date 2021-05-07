package main

type ray struct {
	A vec3
	B vec3
}

func (r *ray) origin() vec3 { return r.A }

func (r *ray) direction() vec3 { return r.B }

func (r *ray) pointAtParameter(t float64) vec3 {
	r.B = r.B.multiplyT(t)
	return r.A.plus(r.B)
}

func (r *ray) ray(a vec3, b vec3) {
	r.A = a
	r.B = b
}
