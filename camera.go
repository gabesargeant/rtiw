package main

type camera struct {
	lowerLeftCorner vec3
	horizontal      vec3
	vertical        vec3
	origin          vec3
}

func (c *camera) camera() {
	c.lowerLeftCorner.vec3(-2.0, -1.0, -1.0)
	c.horizontal.vec3(4.0, 0.0, 0.0)
	c.vertical.vec3(0.0, 2.0, 0.0)
	c.origin.vec3(0.0, 0.0, 0.0)
}

func (c *camera) getRay(u float64, v float64) ray {

	hoz := c.horizontal.multiplyT(u)
	vert := c.vertical.multiplyT(v)
	llc := c.lowerLeftCorner.plus(hoz)
	llc = llc.plus(vert)
	r := ray{}
	r.ray(c.origin, llc)
	return r
}
