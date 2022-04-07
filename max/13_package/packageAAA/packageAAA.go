package packageAAA

const (
	PubConst  = 1
	privConst = 2
)

var PubInt int
var privInt int

type PubStruct struct {
	aaa int
	AAA int
}

type privStruct struct {
	aaa int
	AAA int
}

func Pub() {
	//공개
}

func priv() {
	//비공개
}
