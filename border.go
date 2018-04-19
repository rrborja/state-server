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

func (line Line) PointA() Point {
	return line.P1
}

func (line Line) PointB() Point {
	return line.P2
}

// Checks whether a given point ray casts to the right in a straight line towards the given line
func (line Line) DoIntersect(p Point) bool {
	a := line.P1
	b := line.P2
	return (a.Y > p.Y) != (b.Y > p.Y) &&
		p.X < (b.X-a.X)*(p.Y-a.Y)/(b.Y-a.Y)+a.X
}

// A simple ray casting algorithm that checks all borders may be intersected. It counts how many
// times borders are intersected. If there are even times that have been intersected, then the
// point in question is said to be inside the polygon. Otherwise, it is outside the polygon.
func (polygon Polygon) Within(point Point) (i bool) {
	for _, line := range polygon {
		if line.DoIntersect(point) {
			i = !i
		}
	}
	return
}

