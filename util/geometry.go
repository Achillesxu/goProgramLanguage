package util

import "math"

type PointF struct {
	X, Y float64
}

// traditional function
func Distance(p, q PointF) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p PointF) Distance(q PointF) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
