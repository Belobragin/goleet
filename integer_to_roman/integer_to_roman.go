// Problem 12 Integer to Roman
// https://leetcode.com/problems/integer-to-roman/
// it accepts integer inputs < 4000

package main

import (
	"fmt"
	"math"
)

func makeRoman (X, Y, Z string, poryadok int) (Answer string) {
	switch poryadok {
	case 1,2,3: for i:=0; i< poryadok; i+=1 {
		Answer = Answer + X
		//fmt.Println("RomanNum 1-3", Answer)
		}
	case 4: Answer = X+Y
	case 5: Answer = Y
	case 6, 7, 8: 	for i:=0; i< poryadok-5; i+=1 {
				Answer = Answer + X
				}
			Answer = Y + Answer
			//fmt.Println("RomanNum 6-8", Answer)
	case 9: Answer = X+Z
	
	default: break
	}
	return	
}


func intToRoman(num int) string {
   Romans:= [7]string {"I", "V", "X", "L", "C", "D", "M"}
   var RomanNum string   
   var promeg [4]int 
   var (slugmax=float64(num)
        slugmin float64
	)
   
   for j:=0; j<4; j+=1 {
   	slugmax, slugmin = math.Modf(float64(slugmax)/10.0)
   	promeg[j] = int(math.Round(slugmin*10))
   	}

	
   fmt.Println(promeg)

   RomanNum = RomanNum + makeRoman(Romans[6], "", "", promeg[3])
   for ii:=3; ii>0; ii-=1 {RomanNum = RomanNum + makeRoman(Romans[2*ii-2], Romans[2*ii-1], Romans[2*ii], promeg[ii-1])}
   //fmt.Println(RomanNum)
	
   
return RomanNum	
}

func main() {
    fmt.Println(intToRoman(582))
}