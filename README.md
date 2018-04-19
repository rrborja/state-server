# State Server API

[![GoDoc](https://godoc.org/github.com/rrborja/state-server?status.svg)](https://godoc.org/github.com/rrborja/state-server)
[![License: GPL v2+](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl.txt)  
[![Build Status](https://travis-ci.org/rrborja/state-server.svg?branch=master)](https://travis-ci.org/rrborja/state-server)
[![codecov](https://codecov.io/gh/rrborja/state-server/branch/master/graph/badge.svg)](https://codecov.io/gh/rrborja/state-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrborja/state-server)](https://goreportcard.com/report/github.com/rrborja/state-server)

# Install

This project is tested on a Go 1.10 but it may work on previous versions of Go starting 1.6

#### Cloning and building the source code

1. Run terminal and execute `git clone https://github.com/rrborja/state-server`
2. Enter the cloned project's main directory and execute `go build`
3. At this point, you must have the compiled executable `state-server` if you're on a UNIX-based OS.

#### Cloning and running the source code using Go

1. Run terminal and execute `git clone https://github.com/rrborja/state-server`
2. Enter the cloned project's main directory
3. Execute `go run main.go`

#### Installing directly using Go install

1. Run terminal and execute `go run https://github.com/rrborja/state-server`
2. Execute `state-server`

# Usage

If you have installed the state-server through Option 1, you can perform:

1. If you are still in the working directory after Option 1, then in your terminal, run `./state-server &`
2. Test the state server by checking if the given point is within Pennsylvania land by running `curl  -d "longitude=-77.036133&latitude=40.513799" http://localhost:8080/` while still in your current working directory
3. You should expect an output `["Pennsylvania"]`