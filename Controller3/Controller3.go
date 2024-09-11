package Controller3

import (
	"MVC_FILES/Model3" //pentru a utiliza algoritmul de determinare a parametrilor
)

func DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax float64) (v, gamma, Vmax, Gammamax float64) { //functie ce determina valorile celor 4 parametri
	solutions := Model3.Optimize(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax) //functia de optimizare, ce determina vectorul de solutii pentru v si gamma
	v = solutions[0]
	gamma = solutions[1]
	Vmax = Model3.CalcVmax(G, rho, S, Czmax)
	Gammamax = Model3.CalcGammamax(P, eta, rho, S, Vmax, Cx0, k, Czmax, G)
	return
}
