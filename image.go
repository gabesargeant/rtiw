package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func outputImage(filename string) {

	width := 200
	height := 100
	ns := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	image := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	world := new(hitTableList)
	cam := camera{}
	cam.camera()

	s1 := &sphere{}
	s1v := vec3{}
	s1v.vec3(0, 0, -1)
	m1 := lambertain{}
	m1v := vec3{}
	m1v.vec3(0.8,0.3,0.3)
	m1.lambertain(m1v)
	s1.sphere(s1v, 0.5, &m1)
	//
	s2 := &sphere{}
	s2v := vec3{}
	s2v.vec3(0, -100.5, -1)
	m2 := lambertain{}
	m2v := vec3{}
	m2v.vec3(0.8,0.8,0.0)
	m2.lambertain(m2v)
	s2.sphere(s2v, 100, &m2)
	//
	s3 := &sphere{}
	s3v := vec3{}
	s3v.vec3(1, 0, -1)
	m3 := lambertain{}
	m3v := vec3{}
	m3v.vec3(0.8,0.6,0.2)
	m3.lambertain(m3v)
	s3.sphere(s3v, 0.5, &m3)
	//
	s4 := &sphere{}
	s4v := vec3{}
	s4v.vec3(-1, 0, -1)
	m4 := lambertain{}
	m4v := vec3{}
	m4v.vec3(0.8,0.3,0.3)
	m4.lambertain(m4v)
	s4.sphere(s4v, 0.5, &m4)

	list := []hitTable{s1, s2, s3, s4}

	world.list = list
	fmt.Println(len(world.list))
	for i := 0; i < len(world.list); i++ {
		fmt.Println(world.list[i])
	}

	for y := height; y >= 0; y-- {
		for x := 0; x < width; x++ {
			col := vec3{}
			col.vec3(0,0,0)
			for a := 0; a < ns; a++ {
				ur := rand.Float64()
				vr := rand.Float64()
				u := (float64(x) + ur) / float64(width)
				v := (float64(y) + vr) / float64(height)

				r := cam.getRay(u, v)
				r.pointAtParameter(2.0)

				col.plusEq(colour(r, world))

			}
			col = col.divideT(float64(ns))

			col.e[0] = math.Sqrt(col.e[0])
			col.e[1] = math.Sqrt(col.e[1])
			col.e[2] = math.Sqrt(col.e[2])

			ir := uint8(255.99 * col.e[0])
			ig := uint8(255.99 * col.e[1])
			ib := uint8(255.99 * col.e[2])

			posY := int(math.Abs(float64(height - y)))

			image.Set(x, posY, color.RGBA{ir, ig, ib, 0xff})

		}
	}
	fmt.Println("Writing out: ", filename)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file")
		os.Exit(-99)
	}

	png.Encode(file, image)

}

func colour(r ray, world *hitTableList, depth int) vec3 {

	rec := &hitRecord{}
	hit, rec := world.hitFunc(r, 0.001, math.MaxFloat64, rec)
	

	if hit {

		scattered := ray{}
		attenuation := vec3{}
		if(depth < 50 && rec.matPtr.scatter(r, rec, &attenuation, &scattered)){
			return attenuation.mult(colour(scattered, world, depth+1))
		}else{
			tmp := vec3{}
			tmp.vec3(0,0,0)
			return tmp
		}
	}

	unitDirection := unitVector(r.direction())
	t := 0.5 * (unitDirection.y() + 1.0)
	v1 := vec3{}
	v2 := vec3{}
	v1.vec3(1.0, 1.0, 1.0)
	v2.vec3(0.5, 0.7, 1.0)
	start := v1.multiplyT(1.0 - t)
	end := v2.multiplyT(t)

	return start.plus(end)

}

func randomInInitSphere() vec3 {
	p := vec3{}
	for {

		a:= vec3{}
		a.vec3(rand.Float64(),rand.Float64(),rand.Float64())
		b:= vec3{}
		b.vec3(1,1,1)
		c := a.minus(b)
		p = c.multiplyT(2.0)

		if p.squaredLength() >= 1.0 {
			break;
		}
		
	}
	return p
}

func createOutputDirectory(outDir string) {

	err := os.Mkdir(outDir, 0777)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error with creating out dir")
	}
}
