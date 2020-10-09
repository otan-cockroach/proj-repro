package pkgb

import (
	"github.com/cockroachdb/proj-repro/proj"
	"github.com/cockroachdb/proj-repro/repro"
)

func SomeOtherFunc() bool {
	return proj.IsLatLng()
}

func CircularFunc() bool {
	return repro.CallMe()
}
