package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)
	// count chars
	var lower,upper,digits,spec int
	for _,x:=range input{
		if 'a'<=x&&x<='z'{
			lower++
		}else if 'A'<=x&&x<='Z'{
			upper++
		}else if '0'<=x&&x<='9' {
			digits++
		}else{
			spec++
		}
	}

	// display output
	fmt.Println(lower,upper,digits,spec)
	// ping https://www.omnicalculator.com/other/password-entropy
}
