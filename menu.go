package main
import("fmt")
func main(){
	fmt.Println("menu")
	var input string =""
	for {
		fmt.Println("pl enter your choice:")
		fmt.Scan(&input)
		switch input{
		case "hello":
			fmt.Println("hello!")
		case "help":
			fmt.Println("This is help cmd")
		case "quit":
			return 
		default:
			fmt.Println("Wrong cnm")

		}
	    

	}
}