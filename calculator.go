package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int
	Minutes int
}

func main() {

	// Solicitar a primeira hora do usuário
	fmt.Println("Digite a primeira hora (no formato HH:MM):")
	hora1 := readTimeInput()

	// Solicitar a segunda hora do usuário
	fmt.Println("Digite a segunda hora (no formato HH:MM):")
	hora2 := readTimeInput()

	//Solicitar a operação de soma ou subtração
	fmt.Println("Deseja somar ou subtrair as horas? (soma/subtracao):")
	operacao := readOperationInput()

	// Calcular o resultado com base na operação escolhida
	var resultado Time
	if operacao == "soma" {
		resultado = sumTimes(hora1, hora2)
	} else if operacao == "subtracao" {
		resultado = subtractTimes(hora1, hora2)
	} else {
		fmt.Println("Operação invalida.")
		return
	}

	// Imprimir o resultado em horas e minutos
	fmt.Printf("O resultado da operacao : %02d:%02d\n", resultado.Hours, resultado.Minutes)
}

func readTimeInput() Time {
	var input string
	fmt.Scanln(&input)
	horasMinutos := strings.Split(input, ":")
	horas, _ := strconv.Atoi(horasMinutos[0])
	minutos, _ := strconv.Atoi(horasMinutos[1])
	return Time{Hours: horas, Minutes: minutos}
}

func readOperationInput() string {
	var operacao string
	fmt.Scanln(&operacao)
	return strings.ToLower(operacao)
}

func sumTimes(hora1, hora2 Time) Time {
	totalMinutos := (hora1.Hours*60 + hora1.Minutes) + (hora2.Hours*60 + hora2.Minutes)
	horas := totalMinutos / 60
	minutos := totalMinutos % 60
	return Time{Hours: horas, Minutes: minutos}
}

func subtractTimes(hora1, hora2 Time) Time {
	totalMinutos := (hora1.Hours*60 + hora1.Minutes) - (hora2.Hours*60 + hora2.Minutes)
	horas := totalMinutos / 60
	minutos := totalMinutos % 60
	return Time{Hours: horas, Minutes: minutos}
}
