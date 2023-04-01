package main

import (
	"log"

	"github.com/solarlune/resolv"
)

type portalObject struct {
	worldKey int
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

	log.Fatalf("hitbox data is not set, this a critical error: ID:%v\nContents:%v\n", &o.Data, o.Data)
	return nil
}

func portalData(o *resolv.Object) *portalObject {
	data := assertPortalData(o)
	return data
}
