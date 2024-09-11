package View3

import (
	"bufio"   //pentru operatia de citire
	"fmt"     //pentru afisarea in fisier
	"os"      //pentru utilizarea
	"strconv" //convertirea la numar
	"strings" //pentru lucrul cu siruri de caractere
)

func Read() (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) { //citirea din fisier a parametrilor
	var rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx float64

	filei, _ := os.Open("input.txt") //deschiderea fisierului de input
	defer filei.Close()
	scanner := bufio.NewScanner(filei) //crearea scanner-ului pentru citirea datelor

	//scanner.Scan() citeste secvential datele
	//strings.TrimPrefix creaza un sir ce contine doar numarul de tip float
	//strconv.ParseFloat converteste sirul la un numar float

	scanner.Scan()
	line1 := scanner.Text() //noua linie pentru rho
	rhoStr := strings.TrimPrefix(line1, "rho=")
	rho, _ = strconv.ParseFloat(strings.TrimSpace(rhoStr), 64)

	scanner.Scan()
	line2 := scanner.Text() //noua linie pentru G
	GStr := strings.TrimPrefix(line2, "G=")
	G, _ = strconv.ParseFloat(strings.TrimSpace(GStr), 64)

	scanner.Scan()
	line3 := scanner.Text() //noua linie pentru S
	SStr := strings.TrimPrefix(line3, "S=")
	S, _ = strconv.ParseFloat(strings.TrimSpace(SStr), 64)

	scanner.Scan()
	line4 := scanner.Text() //P
	PStr := strings.TrimPrefix(line4, "P=")
	P, _ = strconv.ParseFloat(strings.TrimSpace(PStr), 64)

	scanner.Scan()
	line5 := scanner.Text() //Cx0
	Cx0Str := strings.TrimPrefix(line5, "Cx0=")
	Cx0, _ = strconv.ParseFloat(strings.TrimSpace(Cx0Str), 64)

	scanner.Scan()
	line6 := scanner.Text() //noua linie pentru eta
	etaStr := strings.TrimPrefix(line6, "eta=")
	eta, _ = strconv.ParseFloat(strings.TrimSpace(etaStr), 64)

	scanner.Scan()
	line7 := scanner.Text() //noua linie pentru k
	kStr := strings.TrimPrefix(line7, "k=")
	k, _ = strconv.ParseFloat(strings.TrimSpace(kStr), 64)

	scanner.Scan()
	line8 := scanner.Text() //noua linie pentru Cz
	CzStr := strings.TrimPrefix(line8, "Cz=")
	Cz, _ = strconv.ParseFloat(strings.TrimSpace(CzStr), 64)

	scanner.Scan()
	line9 := scanner.Text() //noua linie pentru Czmax
	CzmaxStr := strings.TrimPrefix(line9, "Czmax=")
	Czmax, _ = strconv.ParseFloat(strings.TrimSpace(CzmaxStr), 64)

	scanner.Scan()
	line10 := scanner.Text() //noua linie pentru Cx
	CxStr := strings.TrimPrefix(line10, "Cx=")
	Cx, _ = strconv.ParseFloat(strings.TrimSpace(CxStr), 64)

	return rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx //returnarea valorilor citite
}

func DispRes(rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx, v, Vmax, gamma, Gammamax float64) { //afisarea in fisier
	fileo, _ := os.Create("output.txt") //crearea fisierului de afisare
	defer fileo.Close()
	fileo.WriteString(fmt.Sprintf("v: %f\n", v))               //afisarea vitezei
	fileo.WriteString(fmt.Sprintf("Vmax: %f\n", Vmax))         //afisarea vitezei maxime
	fileo.WriteString(fmt.Sprintf("gamma: %f\n", gamma))       //afisarea unghiului
	fileo.WriteString(fmt.Sprintf("Gammamax: %f\n", Gammamax)) //afisarea unghiului maxim
}
