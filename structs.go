package main
import "fmt" 

var (
	userName string
)

const (
	userCompany string = "XYZ"
)

type User struct {
	name string
	phone int
	company string
}
func main() {
	user1 := User { "Ayushi", 8223454322, "XYZ"}
	//to create on heap:
	user2 := new (User)
	user2.name = "Ayushi"
	user2.phone = 8223454322
	fmt.Println(user1)
	fmt.Println(user2)
}