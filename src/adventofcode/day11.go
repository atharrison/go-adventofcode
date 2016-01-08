package adventofcode

import (
	"fmt"
)

func ExecuteDay11(input string) {

	fmt.Printf("Starting with %s\n", input)
	valid := false
	nextPwd := input
	for !valid {
		nextPwd = IncrementPassword(nextPwd)
		valid = CheckPasswordValidity(nextPwd)
	}
	fmt.Printf("Next Valid: [%s]\n", nextPwd)
}

func CheckPasswordValidity(password string) bool {
	valid := HasIncreasingStraight(password, 3) && AnyTwoPairs(password)
	return valid
}

func AnyTwoPairs(input string) bool {

	pair1Found := false
	pair1Char := byte(1)
	for idx := 0; idx < len(input)-1; idx++ {
		if input[idx] == input[idx+1] {
			if !pair1Found {
				pair1Found = true
				pair1Char = byte(input[idx])
			} else if pair1Found && input[idx] != pair1Char {
				return true
			}
		}
	}
	return false
}

func HasIncreasingStraight(input string, minLength int) bool {
	for idx := 0; idx < len(input)-(minLength-1); idx++ {

		potentialStraight := input[idx : idx+minLength]
		//		fmt.Printf("Potential: [%s]\n", potentialStraight)
		valid := true
		for j := 0; j < minLength-1; j++ {
			//			fmt.Printf("Checking if %s + %d is %s?\n", string(potentialStraight[j]), minLength-j-1, string(potentialStraight[minLength-1]))
			if int(potentialStraight[j])+(minLength-j-1) != int(potentialStraight[minLength-1]) {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}
	return false
}

func IncrementPassword(password string) string {

	newPwd := password
	ptr := len(password) - 1

	charCode := newPwd[ptr]
	next, overflow := NextLowercaseLetterCode(charCode)
	for {
		fmt.Printf("%s\t >%s<(%d) After %d is %d, overflow: %v\n", password, string(charCode), ptr, charCode, next, overflow)
		newPwd = ReplaceCharInString(newPwd, next, ptr)
		fmt.Printf("New: [%s]\n", newPwd)
		if !overflow {
			break
		}

		ptr--
		charCode = newPwd[ptr]
		next, overflow = NextLowercaseLetterCode(charCode)
	}

	return newPwd
}

func ReplaceCharInString(input string, newChar byte, loc int) string {
	front := input[:loc]
	back := input[loc+1:]
	fmt.Printf("Converting %s into [%s]%s[%s]\n", input, front, string(newChar), back)
	newStr := front + string(newChar) + back
	return newStr
}

func NextLowercaseLetterCode(code uint8) (uint8, bool) {
	nextCode := code + 1
	overflow := false
	switch nextCode {
	case byte('{'): // '{' is code after 'z'
		nextCode = byte('a')
		overflow = true
	case byte('i'), byte('l'), byte('o'): // Skip invalid letters
		nextCode++
		overflow = false
	}

	return nextCode, overflow
}
