package main

type hitRecord struct {
	t      float64
	p      vec3
	normal vec3
	matPtr material
}

type hitTable interface {
	hitFunc(r ray, tMin float64, tMax float64, rec hitRecord) (bool, hitRecord)
}

type hitTableList struct {
	list []hitTable
	rec  hitRecord
}

func (h *hitTableList) hitFunc(r ray, tMin float64, tMax float64, rec hitRecord) (bool, hitRecord) {

	hitAnything := false
	hit := false
	closestSoFar := tMax
	tmprec := hitRecord{}

	for i := 0; i < len(h.list); i++ {
		//fmt.Println(i)
		hit, rec = h.list[i].hitFunc(r, tMin, closestSoFar, tmprec)
		if hit==true {
			hitAnything = true
			closestSoFar = rec.t
			tmprec = rec			
		}

	}

	return hitAnything, rec

}
