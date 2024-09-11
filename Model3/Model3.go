package Model3

import (
	"math" //accesarea functiilor matematice

	"gonum.org/v1/gonum/floats"   // lucru cu vectori și matrici
	"gonum.org/v1/gonum/optimize" // optimizare (minimalizare)
)

func Ecuatii(x []float64, rho, S, P, G, Cx0, k, eta, Cz, Cx float64) []float64 { //determinarea celor 2 functii F1 si F2
	v := x[0]                                                                     // viteza
	gamma := x[1]                                                                 // unghiul
	F1 := P*eta - (rho / 2 * S * math.Pow(v, 3) * Cx) - (v * G * math.Sin(gamma)) // prima functie
	F2 := G*math.Cos(gamma) - (rho / 2 * S * math.Pow(v, 2) * Cz)                 // a doua functie
	return []float64{F1, F2}
}

func FunctieObiectiv(x []float64, rho, S, P, G, Cx0, k, eta, Cz, Cx float64) float64 { //functia obiectiv ce va fi optimizata prin minimizare
	f := Ecuatii(x, rho, S, P, G, Cx0, k, eta, Cz, Cx) // aflăm cele 2 funcții
	return floats.Norm(f, 2)                           // calculăm sqrt(F1^2+F2^2), astfel ecuațiile F1=0 și F2=0 se reduc doar la una
}

func Optimize(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax float64) []float64 { //functia de optimizare
	problem := optimize.Problem{ //definirea problemei ce va fi supusa optimizarii
		Func: func(x []float64) float64 { //functia obiectiv
			return FunctieObiectiv(x, rho, S, P, G, Cx0, k, eta, Cz, Cx)
		},
	}
	start := []float64{50, 0.01}                             // valori inițiale pentru v și gamma
	result, _ := optimize.Minimize(problem, start, nil, nil) //realizarea operatiei de optimizare, parametrii nil sunt setari suplimentare
	return result.X                                          //returnarea vectorului ce contine solutiile pentru v si gamma
}

func CalcVmax(G, rho, S, Czmax float64) float64 { //determinarea vitezei maxime
	return math.Sqrt((2 * G) / (rho * S * Czmax))
}

func CalcGammamax(P, eta, rho, S, Vmax, Cx0, k, Czmax, G float64) float64 { //determinarea unghiului maxim
	argAsin := (P*eta - (rho/2)*S*math.Pow(Vmax, 3)*(Cx0+k*math.Pow(Czmax, 2))) / (Vmax * G)
	return math.Asin(argAsin)
}
