package stringtonumber

import "strconv"

func StringToNumber(str string) int {
       result, err := strconv.Atoi(str)
       if err != nil {
	       return 0 // or handle error as needed
       }
       return result
}