package main

import (
	"MVC_FILES/Controller3" //pentru obtinerea rezultatelor
	"MVC_FILES/View3"       //pentru aflarea rezultatelor
)

func main() {
	// input
	rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx := View3.Read()

	// obtinerea rezultatelor
	v, gamma, Vmax, Gammamax := Controller3.DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)

	// output
	View3.DispRes(rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx, v, Vmax, gamma, Gammamax)
}
