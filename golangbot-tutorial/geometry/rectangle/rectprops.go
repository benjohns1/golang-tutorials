package rectangle

import "math"
import "fmt"

func init() {
	fmt.Println("rectangle package initialized")
}

// Area calculates and returns the area of a rectangle given its length and width
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal calculates and returns the diagonal length of a rectangle given its length and width
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
