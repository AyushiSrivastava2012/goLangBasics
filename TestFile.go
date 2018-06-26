package main
import "fmt" 


var (
	userName string
)

const (
	userCompany string = "Airtel"
)
func main() {
	userName = "Ayushi"
	printName(userName,userCompany)
}

func printName(name... string){
	
	for i=0; i< len(name);i++ {
		fmt.Println(name[i])
	}	
}