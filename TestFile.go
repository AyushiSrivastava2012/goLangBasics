package main
import "fmt" 

var (
	userName string
)

const (
	userCompany string = "XYZ"
)
func main() {
	userName = "Ayushi"
	print("Enter 1 for username and 2 for company name")

	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1 :
		print(userName)
	case 2:
		print(userCompany)
	default:
		print("Invalid choice")
	}
}

func print(name string){
		fmt.Println(name)
	
}