package worlds

import (
	"log"

	"github.com/solarlune/resolv"
)

type portalObject struct {
	worldKey int // we use key because not all worlds are set when the portals are being init'd
	x        int
	y        int
}

func initPortal(o *resolv.Object, wKey int, x int, y int) {
	o.Data = &portalObject{
		worldKey: wKey,
		x:        x,
		y:        y,
	}
}

func assertPortalData(o *resolv.Object) *portalObject {
	if data, ok := o.Data.(*portalObject); ok {
		return data
	}

	log.Fatal("portal data is not set, this a critical error")
	return nil
}

func portalData(o *resolv.Object) *portalObject {
	data := assertPortalData(o)
	return data
}
