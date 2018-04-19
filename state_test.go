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
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// New Jersey, Hawaii, Alaska, Rhode Island, Delaware, Connecticut

// Tests to check whether the states.json file is present in the current project directory
//
// What I've noticed in the states.json file is that not all 50 US States are considered
// mainly because what I have thought is that one of these states are either not
// contiguous or too small to consider. Michigan is also excluded because it is the only
// US state that has two peninsulas which would had been considered two polygons in this
// test. Subsequently, the following state are excluded:
//
// 			Alaska
// 			Connecticut
// 			Delaware
// 			Hawaii
// 			Michigan
// 			New Jersey
// 			Rhode Island
func TestInitializeCorrectFile(t *testing.T) {
	states, err := StateDescriptions("states.json")
	assert.NoError(t, err)

	statesNotToConsider := []string{
		"Alaska", "Connecticut", "Delaware", "Hawaii", "Michigan", "New Jersey", "Rhode Island"}
	for _, state := range states {
		for _, invalidState := range statesNotToConsider {
			assert.NotEqual(t, invalidState, state.State,
				fmt.Sprintf("%s is a state that is not part of the original states.json file", invalidState))
		}
	}

	// Let's check if there are exactly all 50 states but the excluded states stated above in the file
	assert.Len(t, states, 50-len(statesNotToConsider))
}

// In this test, we'll have to assume that the points of the given border in the
// states.json file are geometrically sorted.
func TestInitializeCorrectState(t *testing.T) {
	stateDescriptions, _ := StateDescriptions("states.json")
	states := Initialize(stateDescriptions)

	// Let's check whether all states in the states.json file are correctly sorted
	// At least when the first point in the list is equal to the last point in the
	// list so that we'll know that the state's borders are drawable virtually
	for name, border := range states {
		assert.EqualValues(t, border[0].PointA(), border[len(border)-1].PointB(),
			fmt.Sprintf("%s is corrupted from the states.json file", name))
	}

	PABorders := Polygon{
		&Line{Point{-77.475793, 39.719623}, Point{-80.524269, 39.721209}},
		&Line{Point{-80.524269, 39.721209}, Point{-80.520592, 41.986872}},
		&Line{Point{-80.520592, 41.986872}, Point{-74.705273, 41.375059}},
		&Line{Point{-74.705273, 41.375059}, Point{-75.142901, 39.881602}},
		&Line{Point{-75.142901, 39.881602}, Point{-77.475793, 39.719623}},
	}

	pennsylvania := states["Pennsylvania"]

	for i, border := range PABorders {
		assert.Equal(t, border, pennsylvania[i])
	}
}

// We'll make sure all borders for this test assumes no horizontal lines
func TestToCheckNoHorizontalLines(t *testing.T) {
	stateDescriptions, _ := StateDescriptions("states.json")
	states := Initialize(stateDescriptions)

	for _, border := range states {
		for _, b := range border {
			y1 := b.PointA().Y
			y2 := b.PointB().Y
			assert.NotEqual(t, y1, y2)
		}
	}
}

func TestNoSuchFile(t *testing.T) {
	_, err := StateDescriptions("states-dummy.json")
	assert.Error(t, err)
}
