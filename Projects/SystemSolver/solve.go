//Main package creates a webapp for solving systems of equations
package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	http.ListenAndServe(":8080", handler())
}

//handler checks if the request url is valid
func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/solve" {
			process(w, r)
		} else {
			http.Error(w, "Bad Request: Incorrect URL", http.StatusBadRequest)
		}
	}
}

//process queries the input from the URL
func process(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query()["coef"]
	if len(input[0]) != 23 {
		http.Error(w, "Bad Request: Incorrect number of coefficients.", http.StatusBadRequest)
	} else {
		coef := ConvertInput(input[0])
		matrix := CreateInitialMatrix(coef)
		FindSolution(matrix, w)
	}
}

//ConvertInput converts the string query into float64 and returns an array of coefficients
func ConvertInput(input string) [12]float64 {
	removeSpecial := regexp.MustCompile(`(?m)[^0-9]`)
	input = removeSpecial.ReplaceAllString(input, "")
	var coef [12]float64
	for i, num := range input {
		floatNum := float64(num - '0')
		coef[i] = floatNum
	}

	return coef
}

//CreateInitialMatrix takes an array of coefficients and returns a 2D matrix
func CreateInitialMatrix(coef [12]float64) [3][4]float64 {
	var mat [3][4]float64
	i := 0
	for j := 0; j < 3; j++ {
		for k := 0; k < 4; k++ {
			mat[j][k] = coef[i]
			i++
		}
	}
	return mat
}

//GetDeterminant returns the determinant of a given matrix
func GetDeterminant(mat [3][3]float64) float64 {
	var ans float64
	ans = mat[0][0]*(mat[1][1]*mat[2][2]-mat[2][1]*mat[1][2]) - mat[0][1]*(mat[1][0]*mat[2][2]-mat[1][2]*mat[2][0]) + mat[0][2]*(mat[1][0]*mat[2][1]-mat[1][1]*mat[2][0])
	return ans
}

//FindSolution solves the system of equations using cramer's rule
func FindSolution(coef [3][4]float64, w http.ResponseWriter) {
	// Matrix d using coef as given in cramer's rule
	d := [3][3]float64{
		{coef[0][0], coef[0][1], coef[0][2]},
		{coef[1][0], coef[1][1], coef[1][2]},
		{coef[2][0], coef[2][1], coef[2][2]},
	}
	// Matrix d1 using coef as given in cramer's rule
	d1 := [3][3]float64{
		{coef[0][3], coef[0][1], coef[0][2]},
		{coef[1][3], coef[1][1], coef[1][2]},
		{coef[2][3], coef[2][1], coef[2][2]},
	}
	// Matrix d2 using coef as given in cramer's rule
	d2 := [3][3]float64{
		{coef[0][0], coef[0][3], coef[0][2]},
		{coef[1][0], coef[1][3], coef[1][2]},
		{coef[2][0], coef[2][3], coef[2][2]},
	}
	// Matrix d3 using coef as given in cramer's rule
	d3 := [3][3]float64{
		{coef[0][0], coef[0][1], coef[0][3]},
		{coef[1][0], coef[1][1], coef[1][3]},
		{coef[2][0], coef[2][1], coef[2][3]},
	}

	D := GetDeterminant(d)
	D1 := GetDeterminant(d1)
	D2 := GetDeterminant(d2)
	D3 := GetDeterminant(d3)
	fmt.Fprintf(w, "system:\n")
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", coef[0][0], coef[0][1], coef[0][2], coef[0][3])
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", coef[1][0], coef[1][1], coef[1][2], coef[1][3])
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", coef[2][0], coef[2][1], coef[2][2], coef[2][3])
	if D != 0 {
		x := D1 / D
		y := D2 / D
		z := D3 / D

		fmt.Fprintf(w, "solution:\n")
		fmt.Fprintf(w, "x = %.2f, y = %.2f, z = %.2f\n", x, y, z)
	} else {
		if D1 == 0 && D2 == 0 && D3 == 0 {
			fmt.Fprintln(w, "dependent - with multiple solutions")

		} else if D1 != 0 || D2 != 0 || D3 != 0 {
			fmt.Fprintln(w, "inconsistent - no solution")
		}
	}

}
