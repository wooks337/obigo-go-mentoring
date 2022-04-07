package main

import (
	"obigo-go-mentoring/max/13_package/packageAAA"
	packageAAABBB "obigo-go-mentoring/max/13_package/packageAAA/packageBBB"
	"obigo-go-mentoring/max/13_package/packageBBB"
	_ "obigo-go-mentoring/max/13_package/packageCCC"
)

func main() {

	packageAAA.Pub()
	packAAA := packageAAA.PubStruct{} //대문자만 접근가능
	packAAA.AAA = 123                 //대문자만 접근가능

	packageBBB.PackBBB()
	packageAAABBB.PackAAABBB()
}
