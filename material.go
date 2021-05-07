package main

//Interface for material types.
type material interface {
	scatter(rayIn ray, rec *hitRecord, attenuation vec3, scattered *ray) (bool, *ray)
}


