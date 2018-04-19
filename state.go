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
	"io/ioutil"
	"encoding/json"
)

type Point struct {
	X float64
	Y float64
}

type Line struct {
	P1 Point
	P2 Point
}

type Polygon []Liner
type States map[string]Polygon

type Liner interface {
	PointA() Point
	PointB() Point
	Caster
}

type Caster interface {
	DoIntersect(Point) bool
}

type StateDescription struct {
	State string		`json:"state"`
	Border [][]float64	`json:"border"`
}

// Load the state configuration file containing all U.S. states' names
// and the geometric representation of the corresponding state
func StateDescriptions(filename string) ([]StateDescription, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var stateDescriptions []StateDescription
	json.Unmarshal(raw, &stateDescriptions)

	return stateDescriptions, nil
}

// Initialize the loaded state file into the state server's State
// internal database
func Initialize(stateDescriptions []StateDescription) States {
	states := make(States, len(stateDescriptions))

	for _, stateDescription := range stateDescriptions {
		polygon := make(Polygon, len(stateDescription.Border) - 1)

		// Scan all the points so that they can be processed as
		// segments a.k.a. the border. This assumes that all
		// points are in order
		for i := 0; i < len(stateDescription.Border) - 1; i++ {
			fpoint1 := stateDescription.Border[i]
			fpoint2 := stateDescription.Border[i+1]

			// Each polygon contains at least three borders all connected
			polygon[i] = &Line{Point{fpoint1[0], fpoint1[1]}, Point{fpoint2[0], fpoint2[1]}}
		}

		// Store the resulting polygon into the state
		states[stateDescription.State] = polygon
	}

	return states
}