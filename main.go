package main

import (
	"fmt" 
	"flag"
	"math"
)

func main(){
	inputNumber := flag.Float64("number",  math.Pow(10,5), "The number to solve")
	flag.Parse();

	var sum,i int64
	
	for i=1; i <= int64(*inputNumber); i++ {
		sum += CalcSigma2(i)
	}
	result := (sum % int64( math.Pow(10,5)))
	fmt.Println(sum, result  )
}

func CalcSigma2(n int64) (int64) {

	var sum,i int64

	for i = 1; i<=n; i++ {
		if(n%i == 0){
			sum += i*i
		}
	}
	return sum
}