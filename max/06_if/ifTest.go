package main

func main() {
	if name, age, result := getMyInfo(); result && age == 28 && name == "max" {
		println("Success")
	} else {
		println("Fail")
	}
}

func getMyInfo() (string, int, bool) {
	return "max", 28, true
}
