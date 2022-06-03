package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 24}
	vip := VIPUser{
		User{"화랑", "hwarang", 34},
		4,
		340,
	}
	fmt.Println(user.Name, user.ID, user.Age)
	fmt.Println(vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price)
}
