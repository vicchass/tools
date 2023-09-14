package utisli

import (
	"fmt"
	"math"
)

// return an float64 , which is an averahge of []int , return 1 after coma : 10,3
func Average_slice_int(the_slice []int) float64 {
	var final float64
	var sum int

	for _, v := range the_slice {
		sum += v
	}

	final = float64(sum) / float64(len(the_slice))
	final = (math.Round(final * 10)) / 10
	return final
}

// return the number of times a string is present in a slice of strings
func Number_string_slice(the_slice []string, word string) int {
	var num int
	for _, v := range the_slice {
		if v == word {
			num++
		}

	}
	return num

}

// create a slice of int ,min and max value you can use
func Create_slice_int(min, max int) []int {
	var out_slice []int
	//loop to create slice
	for {
		var to_add int
		fmt.Println("type int to add,-1 to stop ")
		fmt.Scan(&to_add)
		if to_add == -1 {
			fmt.Println("returning", out_slice)
			return out_slice
		}
		if to_add < min || to_add > max {
			fmt.Println("error with range , try again")
			continue
		} else {
			out_slice = append(out_slice, to_add)

		}

	}

}
