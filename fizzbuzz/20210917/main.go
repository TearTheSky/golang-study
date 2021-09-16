package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz Start!")
	i := 1.00
	result := ""
	threeTimes, fiveTimes, fifteenTimes := 1.00, 1.00, 1.00

	for i < 101 {
		i3 := i / 3.00
		i5 := i / 5.00
		i15 := i / 15.00

		if i15-fifteenTimes == 0.00 {
			result = fmt.Sprintf("%d:\tFizzBuzz", int(i))
			fifteenTimes++
			fiveTimes++
			threeTimes++
		} else if i5-fiveTimes == 0.00 {
			result = fmt.Sprintf("%d:\tBuzz", int(i))
			fiveTimes++
		} else if i3-threeTimes == 0.00 {
			result = fmt.Sprintf("%d:\tFizz", int(i))
			threeTimes++
		} else {
			result = fmt.Sprintf("%d:\t", int(i))
		}

		fmt.Println(result)
		result = ""
		i += 1
	}
}
