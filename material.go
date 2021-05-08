package main

//Interface for material types.
type material interface {
	scatter(rayIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool
}

type lambertain struct{
	albedo vec3
}

type metal struct{
	albedo vec3
}

func (m *metal) metal(a vec3){
	m.albedo = a
}

func (l *lambertain) lambertain(a vec3){
	l.albedo = a
}

func (l *lambertain) scatter(rayIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool{
	target := rec.p.plus(rec.normal.plus(randomInInitSphere()))
	//(rec.p, target.minus(rec.p))
	rr := ray{}
	rr.ray(rec.p, target.minus(rec.p))
	scattered = &rr
	attenuation = &l.albedo
	return true

}

func reflect(v vec3, n vec3) vec3 {
	return v.minus(n.multiplyT(2.0 * dot(v,n)))
}

func (m *metal) scatter(rayIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool{
	
	reflected := reflect(unitVector(rayIn.direction()), rec.normal)
	rr := ray{}
	rr.ray(rec.p, reflected)
	scattered = &rr
	attenuation = &m.albedo
	return true
}


