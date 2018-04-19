/*
 * State Server API
 * Copyright (C) 2018  Ritchie Borja
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntersectToVerticalLine(t *testing.T) {
	line := Line{Point{0,8}, Point{0, -8}}
	assert.True(t, line.DoIntersect(Point{-8, 0}))
}

func TestIntersectToHorizontalLine(t *testing.T) {
	line := Line{Point{-8, 0}, Point{8, 0}}
	assert.False(t, line.DoIntersect(Point{-13, 0}))
}

func TestIntersectAfterAVerticalLine(t *testing.T) {
	line := Line{Point{0,8}, Point{0, -8}}
	assert.False(t, line.DoIntersect(Point{8, 0}))
}

func TestIntersectToNegativeSlopedLine(t *testing.T) {
	line := Line{Point{1, 8}, Point{12, -13}}
	assert.True(t, line.DoIntersect(Point{5, -8}))
}

func TestIntersectToSlopedLine(t *testing.T) {
	line := Line{Point{10, 8}, Point{3, -13}}
	assert.True(t, line.DoIntersect(Point{5, 0}))
}

func Square(p, d float64) Polygon {
	p1 := Point{p, d}
	p2 := Point{d, d}
	p3 := Point{d, p}
	p4 := Point{0, p}

	return Polygon{Line{p1, p2}, Line{p2, p3}, Line{p3, p4}, Line{p4, p1}}
}

// Test if the point lies within the polygon regardless of shape
func TestToCheckPointInSquare(t *testing.T) {
	square := Square(0, 20)

	assert.True(t, square.Within(Point{5, 13}))
}

// Test the left side of the square but point lies outside of the polygon
func TestToCheckPointLeftsideOfSquare(t *testing.T) {
	square := Square(0, 15)

	assert.False(t, square.Within(Point{-13, 7}))
}

func TestToCheckPointRightsideOfSquare(t *testing.T) {
	square := Square(0, 15)

	assert.False(t, square.Within(Point{-40, 7}))
}

func SampleConvex() Polygon {
	p1 := Point{2, 30}
	p2 := Point{6, 16}
	p3 := Point{8, 20}
	p4 := Point{9, 0}
	p5 := Point{-3, 0}

	return []Liner{Line{p1,p2}, Line{p2,p3}, Line{p3,p4}, Line{p4,p5}, Line{p5,p1}}
}

func TestIfPointIsInConvexPolygonExpectingMultipleIntersections(t *testing.T) {
	point := Point{3, 18}

	convex := SampleConvex()

	assert.True(t, convex.Within(point))
}

func TestIfLocationIsInPennsylvania(t *testing.T) {
	location := Point{-77.036133, 40.513799}

	p1 := Point{-77.475793, 39.719623}
	p2 := Point{-80.524269, 39.721209}
	p3 := Point{-80.520592, 41.986872}
	p4 := Point{-74.705273, 41.375059}
	p5 := Point{-75.142901, 39.881602}
	p6 := Point{-77.475793, 39.719623}

	var PABorders Polygon = []Liner{
		Line{p1, p2},
		Line{p2, p3},
		Line{p3, p4},
		Line{p4, p5},
		Line{p6, p1},
	}

	assert.True(t, PABorders.Within(location))
}