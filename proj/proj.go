package proj

// #cgo CXXFLAGS: -std=c++14
// #cgo linux LDFLAGS: -lrt -lm -lpthread
// #cgo windows LDFLAGS: -lshlwapi -lrpcrt4
//
// #include "proj.h"
import "C"

func IsLatLng() bool {
	return false
}
