package main

import (
	"fmt"
	"math"
)

/*
func main() {

	var n, k int
	fmt.Scan(&n, &k)

	factN := factorial(n)
	factK := factorial(k)
	nSubtractK := n - k
	factNSubtractK := factorial(nSubtractK)

	fmt.Println(factN / (factK * factNSubtractK))
}

func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	return fact
}
*/
/*
func main() {
	var a, b int
	fmt.Scan(&a, &b)

	z := sign(a) + sign(b)
	fmt.Println(z)
}

func sign(x int) int {
	var signx int
	if x < 0 {
		signx = -1
	}
	if x == 0 {
		signx = 0
	}
	if x > 0 {
		signx = 1
	}
	return signx
}
*/
/*
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(razr(n) * razr(m))
}

func razr(x int) int {
	count := 0
	for i := x; i > 0; i /= 10 {
		count++
	}
	return count
}
*/
/*
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if lucky(a) == 1 && lucky(b) == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}

func lucky(x int) int {
	sum := 0
	sum1 := 0
	for i := x; i > 0; i /= 10 {
		digit := i % 10
		if i > 999 {
			sum += digit
		}
		if i < 1000 {
			sum1 += digit
		}
	}
	if sum == sum1 {
		return 1
	} else {
		return 0
	}
}
*/
/*
func main() {
	var n, m float64
	fmt.Scan(&n, &m)
	fmt.Println(avg(n) + avg(m))
}

func avg(x float64) float64 {
	var average float64
	sum := float64(0)
	for i := float64(1); i <= x; i++ {
		sum += i
	}
	average = sum / x
	return average
}
*/
/*
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(reverse(a) + reverse(b))
}

func reverse(x int) int {
	new_int := 0
	for x > 0 {
		digit := x % 10
		new_int *= 10
		new_int += digit
		x /= 10
	}
	return new_int
}
*/
/*
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fact2(a), fact2(b), fact2(c))
}

func fact2(x int) int {
	fact := 1
	if x%2 != 0 {
		for i := 1; i <= x; i++ {
			if i%2 != 0 {
				fact *= i
			}
		}
	}
	if x%2 == 0 {
		for i := 1; i <= x; i++ {
			if i%2 == 0 {
				fact *= i
			}
		}
	}
	return fact
}
*/
/*
func main() {
	var sizea int
	fmt.Scan(&sizea)
	a := make([]int, sizea)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	var sizeb int
	fmt.Scan(&sizeb)
	b := make([]int, sizeb)
	for i := 0; i < len(b); i++ {
		fmt.Scan(&b[i])
	}
	fmt.Println(maxmass(a) * maxmass(b))
}

func maxmass(x []int) int {
	max := math.MinInt
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}
*/
/*
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	counta, countb := countNumbers(a), countNumbers(b)
	if counta > countb {
		fmt.Println(1)
	} else {
		if countb > counta {
			fmt.Println(2)
		} else {
		}
		if counta == countb {
			fmt.Println(0)
		}
	}
}

func countNumbers(x int) int {
	count := 0
	for i := x; i > 0; i /= 10 {
		count++
	}
	return count
}
*/
/*
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	counta, countb := countNumbers(a), countNumbers(b)

	if counta > countb {
		fmt.Println(1)
	} else {
		if countb > counta {
			fmt.Println(2)
		} else {
		}
		if counta == countb {
			fmt.Println(0)
		}
	}

}

func countNumbers(x int) int {
	sum := 0
	for i := x; i > 0; i /= 10 {
		digit := i % 10
		sum += digit
	}
	return sum
}
*/
/*
func main() {
	scan := bufio.NewScanner(os.Stdin)
	_ = scan.Scan()
	s1 := scan.Text()
	_ = scan.Scan()
	s2 := scan.Text()
	fmt.Println(countb(s1) + countb(s2))
}

func countb(s string) int {
	count := 0
	for i := range s {
		if s[i] == 'b' {
			count++
		}
	}
	return count
}
*/
/*
func main() {
	var n int
	fmt.Scan(&n)
	if simple(n) == true {
		fmt.Println("prime")
	} else {
		if simple(n) == false {
			fmt.Println("composite")
		}
	}
}

func simple(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%int(i) == 0 {
			return false
		}
	}
	return true
}
*/
/*
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(SumRange(a, b) + SumRange(b, c))
}

func SumRange(x, y int) int {
	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}
	return sum
}
*/
func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(sumznach(a, b))
}

func sumznach(m, n float64) int {
	sum := float64(0)
	for i := float64(m); i <= n; i++ {
		if i < -5 {
			sum = sum + 2*math.Abs(i) - 1
		}
		if i >= -5 && i <= 5 {
			sum += i
		}
		if i > 5 {
			sum = sum + 2*i
		}
	}
	return int(sum)
}
