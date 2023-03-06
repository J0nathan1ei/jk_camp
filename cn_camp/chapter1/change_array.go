package main

import "fmt"

func main() {
	d := [5]string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < len(d); i++ {
		if i == 2 {
			d[i] = "smart"
		} else if i == 4 {
			d[i] = "strong"
		}
	}
	fmt.Println(d)
}
