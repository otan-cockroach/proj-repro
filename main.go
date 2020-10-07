package main

import "fmt"

// #cgo CXXFLAGS: -std=c++14
// #cgo linux LDFLAGS: -lrt -lm -lpthread
// #cgo windows LDFLAGS: -lshlwapi -lrpcrt4
//
// #include "proj.h"
import "C"

func main() {
	spec := `GEOGCS["WGS 84",DATUM["WGS_1984",SPHEROID["WGS 84",6378137,298.257223563,AUTHORITY["EPSG","7030"]],AUTHORITY["EPSG","6326"]],PRIMEM["Greenwich",0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.0174532925199433,AUTHORITY["EPSG","9122"]],AUTHORITY["EPSG","4326"]] | +proj=longlat +datum=WGS84 +no_defs`
	cstr := C.CString(spec)
	var ret C.int
	var major C.double
	var ecc C.double
	C.CR_PROJ_GetProjMetadata(cstr, &ret, &major, &ecc)
	fmt.Println("%v %v %v\n", ret, major, ecc)
}
