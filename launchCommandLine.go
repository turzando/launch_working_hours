// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	var email string = "artur.silva@ubots.com.br"
// 	var date string = "" // 2024-07-19
// 	var firstHour string = ""
// 	var secondHour string = "" // 08:30
// 	var company string = ""
// 	var team string = "Professional Service"
// 	var description string = ""
// 	var paidHours string = "Não" // Sim ou Não

// 	for index, arg := range os.Args {
// 		switch arg {
// 		case "-dt":
// 			date = os.Args[index+1]
// 		case "-1h":
// 			firstHour = os.Args[index+1]
// 		case "-2h":
// 			secondHour = os.Args[index+1]
// 		case "-c":
// 			company = os.Args[index+1]
// 		case "-t":
// 			team = os.Args[index+1]
// 		case "-d":
// 			description = os.Args[index+1]
// 		case "-p":
// 			paidHours = os.Args[index+1]
// 		}
// 	}

// 	fmt.Println(email)
// 	fmt.Println(date)
// 	fmt.Println(firstHour)
// 	fmt.Println(secondHour)
// 	fmt.Println(company)
// 	fmt.Println(team)
// 	fmt.Println(description)
// 	fmt.Println(paidHours)
// }
