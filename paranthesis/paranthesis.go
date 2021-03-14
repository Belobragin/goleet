//Problem 20 "Valid paranthesis"

//official answer (func only):

func isValid(s string) bool {
    var otkr=map[string] int{"[":1, "(":2, "{":3}
    var zakr=map[int] string{1:"]", 2:")", 3:"}"}
    var parant []string
    var lastsymb string
    var _, vvv = math.Modf(float64(len(s))/2)
    if vvv != 0.0 {return false}
	for i, symb := range s {
	//fmt.Println(i, len(s))
		ww := len(parant)		
		oo := string(symb)
		//fmt.Println("last", oo, "pre-last", lastsymb, "parant", parant)
		switch oo {			
			case "[", "(", "{": 
				parant=append(parant, string(symb))
				lastsymb= oo
				//fmt.Println(parant)
			case "]", ")", "}": 
				if len(parant)==0 {
					//mt.Println("Short") 
					return false}
				if rr:=otkr[lastsymb]; oo==zakr[rr] {					
					//if len(parant)!=1 { //return true}
					parant=parant[:ww-1]
					//fmt.Println(len(parant))
					if len(parant)==0 {lastsymb=""
					}else {lastsymb=parant[ww-2]}
						
					//fmt.Println("Pomenyali",lastsymb, "New parant", parant)
					if i-len(s)+1 == 0 && len(parant) ==0  {return true}

				   }else {
					fmt.Println("Y") 
					return false}	
				//}	
            default: {return false}//fmt.Println("Panic - this symblol is not a paranthesis")
			}
		 
    }
return false}

//Full answer (package + debugging):

package main

import (
	"fmt"
	"math"
)

func isValid(s string) bool {
    var otkr=map[string] int{"[":1, "(":2, "{":3}
    var zakr=map[int] string{1:"]", 2:")", 3:"}"}
    var parant []string
    var lastsymb string
    var _, vvv = math.Modf(float64(len(s))/2)
    if len(s)>10000 {return false}
	if vvv != 0.0 {return false}
	for i, symb := range s {
	fmt.Println(i, len(s))
		ww := len(parant)		
		oo := string(symb)
		fmt.Println("last", oo, "pre-last", lastsymb, "parant", parant)
		switch oo {			
			case "[", "(", "{": 
				parant=append(parant, string(symb))
				lastsymb= oo
				fmt.Println(parant)
			case "]", ")", "}": 
				if len(parant)==0 {
					fmt.Println("Short") 
					return false}
				if rr:=otkr[lastsymb]; oo==zakr[rr] {					
					//if len(parant)!=1 { //return true}
					parant=parant[:ww-1]
					//fmt.Println(len(parant))
					if len(parant)==0 {lastsymb=""
					}else {lastsymb=parant[ww-2]}
						
					fmt.Println("Pomenyali",lastsymb, "New parant", parant)
					if i-len(s)+1 == 0 && len(parant) ==0  {return true}
						
				   }else {
					fmt.Println("Y") 
					return false}		
			default: fmt.Println("Panic - this symblol is not a paranthesis")
			}
		 
    }
return false}

func main(){
fmt.Println(isValid("({{{{}}}))"))
}
