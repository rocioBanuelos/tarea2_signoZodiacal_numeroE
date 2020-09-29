package main

import "fmt"

func main() {
	var mes int
	var dia int
	var signo string

	fmt.Scanf("%d\n", &dia)
	fmt.Scanf("%d\n", &mes)

	switch {
	case mes == 1:
		if dia > 0 && dia <= 19 {
			signo = "Capricornio"
		} else if dia >= 20 && dia <= 31 {
			signo = "Acuario"
		} else {
			signo = "No es un día válido"
		}
	case mes == 2:
		if dia > 0 && dia <= 18 {
			signo = "Acuario"
		} else if dia >= 19 && dia <= 29 {
			signo = "Piscis"
		} else {
			signo = "No es un día válido"
		}
	case mes == 3:
		if dia > 0 && dia <= 20 {
			signo = "Piscis"
		} else if dia >= 21 && dia <= 31 {
			signo = "Aries"
		} else {
			signo = "No es un día válido"
		}
	case mes == 4:
		if dia > 0 && dia <= 19 {
			signo = "Aries"
		} else if dia >= 20 && dia <= 30 {
			signo = "Tauro"
		} else {
			signo = "No es un día válido"
		}
	case mes == 5:
		if dia > 0 && dia <= 20 {
			signo = "Tauro"
		} else if dia >= 21 && dia <= 31 {
			signo = "Geminis"
		} else {
			signo = "No es un día válido"
		}
	case mes == 6:
		if dia > 0 && dia <= 20 {
			signo = "Geminis"
		} else if dia >= 21 && dia <= 30 {
			signo = "Cancer"
		} else {
			signo = "No es un día válido"
		}
	case mes == 7:
		if dia > 0 && dia <= 22 {
			signo = "Cancer"
		} else if dia >= 23 && dia <= 31 {
			signo = "Leo"
		} else {
			signo = "No es un día válido"
		}
	case mes == 8:
		if dia > 0 && dia <= 22 {
			signo = "Leo"
		} else if dia >= 23 && dia <= 31 {
			signo = "Virgo"
		} else {
			signo = "No es un día válido"
		}
	case mes == 9:
		if dia > 0 && dia <= 22 {
			signo = "Virgo"
		} else if dia >= 23 && dia <= 30 {
			signo = "Libra"
		} else {
			signo = "No es un día válido"
		}
	case mes == 10:
		if dia > 0 && dia <= 22 {
			signo = "Libra"
		} else if dia >= 23 && dia <= 31 {
			signo = "Escorpio"
		} else {
			signo = "No es un día válido"
		}
	case mes == 11:
		if dia > 0 && dia <= 21 {
			signo = "Escorpio"
		} else if dia >= 22 && dia <= 30 {
			signo = "Sagitario"
		} else {
			signo = "No es un día válido"
		}
	case mes == 12:
		if dia > 0 && dia <= 21 {
			signo = "Sagitario"
		} else if dia >= 22 && dia <= 31 {
			signo = "Capricornio"
		} else {
			signo = "No es un día válido"
		}
	default:
		signo = "No es un mes válido"
	}

	fmt.Println(signo)
}

