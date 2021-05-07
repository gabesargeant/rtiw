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
	s1.sphere(s1v, 0.5)
	s2 := &sphere{}
	s2v := vec3{}
	s2v.vec3(0, -100.5, -1)
	s2.sphere(s2v, 100)

	list := []hitTable{s1, s2}

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
			//	fmt.Println("ur ", ur,  "  vr ", vr)
				u := (float64(x) + ur) / float64(width)
				v := (float64(y) + vr) / float64(height)

				r := cam.getRay(u, v)
				r.pointAtParameter(2.0)

				col.plusEq(colour(r, world))

			}
			col = col.divideT(float64(ns))

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

func colour(r ray, world *hitTableList) vec3 {

	rec := &hitRecord{}
	hit, rec := world.hitFunc(r, 0.0, math.MaxFloat64, rec)

	if hit {

		a := vec3{}

		a.vec3(rec.normal.x()+1, rec.normal.y()+1, rec.normal.z()+1)
		return a.multiplyT(0.5)
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

// func hitSphere(center vec3, radius float64, r ray) float64 {
// 	origin := r.origin()
// 	oc := origin.minus(center)
// 	a := dot(r.direction(), r.direction())
// 	b := dot(oc, r.direction())
// 	c := dot(oc, oc) - radius*radius
// 	discriminant := b*b - 4*a*c

// 	if discriminant < 0 {
// 		return -1.0
// 	}

// 	return (-b - math.Sqrt(discriminant)/(2.0*a))

// }

func createOutputDirectory(outDir string) {

	err := os.Mkdir(outDir, 0777)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error with creating out dir")
	}
}
