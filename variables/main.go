package main

import("fmt")


func main()  {

	fmt.Println("men")
	var name string = "hello"

	var cchar[2] string 
	cchar[0] = "man"
	cchar[1] = "woman"

	for i := 0; i < len(cchar); i++ {
		fmt.Println(cchar[i])
	}

	fmt.Println(name)
}