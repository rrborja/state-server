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
	"net/http"
	"log"

	"strconv"
	"encoding/json"
)

var states States

type ErrorResponse struct {
	Error string
}

func handler(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	longitude, err1 := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err1 != nil {
		err := ErrorResponse{err1.Error()}
		w.WriteHeader(500)
		writer.Encode(err)
		return
	}

	latitude, err2 := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err2 != nil {
		err := ErrorResponse{err2.Error()}
		w.WriteHeader(500)
		writer.Encode(err)
		return
	}

	location := Point{longitude, latitude}

	var searchedStates []string

	for name, state := range states {
		if state.Within(location) {
			searchedStates = append(searchedStates, name)
		}
	}

	writer.Encode(searchedStates)
}

func RunStateServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	stateDescriptions, err := StateDescriptions("states.json")
	if err != nil {
		log.Fatal(err)
	}

	states = Initialize(stateDescriptions)

	RunStateServer()
}