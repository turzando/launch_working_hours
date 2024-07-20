package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {

	email := ""
	repeat := true

	var date, firstHour, secondHour, company, team, description, paidHours string

	for repeat {
		date, firstHour, secondHour, company, team, description, paidHours = getUserInput()

		printUserAnswers(date, firstHour, secondHour, company, team, description, paidHours)

		repeat = !confirmData()
	}

	formsUrl := fmt.Sprintf("https://docs.google.com/forms/d/e/1FAIpQLSdNWxeCckE8_YtuJTiie7EcZY4nwRjXZKSv1qD-imkMgzndiA/formResponse?emailAddress=%s&entry.771844714=%s&entry.1840133132=%s&entry.234363642=%s&entry.1523459200=%s&entry.444546027=%s&entry.287786779=%s&entry.1068426552=%s", email, date, firstHour, secondHour, company, team, description, paidHours)
	fmt.Println(formsUrl)
}

func getUserInput() (string, string, string, string, string, string, string) {
	var date, firstHour, secondHour, company, team, description, paidHours string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Informe a data: ")
	fmt.Scan(&date)
	fmt.Println("Informe o horário que começou: ")
	fmt.Scan(&firstHour)
	fmt.Println("Informe o horário que parou: ")
	fmt.Scan(&secondHour)
	fmt.Println("Informe a empresa para a qual trabalhou: ")
	scanner.Scan()
	company = url.QueryEscape(scanner.Text())
	fmt.Println("Informe o time em que trabalhou: ")
	scanner.Scan()
	team = url.QueryEscape(scanner.Text())
	fmt.Println("Informe a Descrição: ")
	scanner.Scan()
	description = url.QueryEscape(scanner.Text())

	for {
		fmt.Print("Informe se foi hora paga (y/N): ")
		paidHours, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		paidHours = strings.TrimSpace(strings.ToLower(paidHours))

		if paidHours == "" || paidHours == "n" {
			paidHours = "Não"
			break
		} else if paidHours == "y" {
			paidHours = "Sim"
			break
		}
	}

	return date, firstHour, secondHour, company, team, description, paidHours
}

func printUserAnswers(date, firstHour, secondHour, company, team, description, paidHours string) {
	fmt.Println()
	fmt.Println("Essas foram suas respostas: ")
	fmt.Println()
	fmt.Println("Data: ", date)
	fmt.Println("Horário que começou: ", firstHour)
	fmt.Println("Horário que terminou: ", secondHour)
	fmt.Println("Empresa para a qual trabalhou: ", company)
	fmt.Println("Time em que trabalhou: ", team)
	fmt.Println("Descrição: ", description)
	fmt.Println("Horas faturadas: ", paidHours)
}

func confirmData() bool {

	for {
		fmt.Println()
		fmt.Println("Estão corretas (Y/n)?")
		reader := bufio.NewReader(os.Stdin)
		userConfirmation, _ := reader.ReadString('\n')
		userConfirmation = strings.TrimSpace(strings.ToLower(userConfirmation))
		if userConfirmation == "" || userConfirmation == "y" {
			return true
		} else if userConfirmation == "n" {
			return false
		}
	}
}
