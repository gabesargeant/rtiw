package main

//Interface for material types.
type material interface {
	scatter(rayIn ray, rec *hitRecord, attenuation vec3, scattered ray) (bool, ray, vec3)
}

type lambertain struct {
	albedo vec3
}

type metal struct {
	albedo vec3
}

func (m *metal) metal(a vec3) {
	m.albedo = a
}

func (l *lambertain) lambertain(a vec3) {
	l.albedo = a
}

func (l *lambertain) scatter(rayIn ray, rec *hitRecord, attenuation vec3, scattered ray) (bool, ray, vec3) {
	//target := rec.p.plus(rec.normal.plus(randomInInitSphere()))
	//scattered.ray(target,rec.p)
	attenuation = l.albedo
	return true, scattered, attenuation
}

func reflect(v vec3, n vec3) vec3 {
	return v.minus(n.multiplyT(2.0 * dot(v, n)))
}

func (m *metal) scatter(rayIn ray, rec *hitRecord, attenuation vec3, scattered ray) (bool, ray, vec3) {

	reflected := reflect(unitVector(rayIn.direction()), rec.normal)
	scattered.ray(rec.p, reflected)
	attenuation = m.albedo
	hit := dot(scattered.direction(), rec.normal) > 0
	return hit, scattered, attenuation
}
