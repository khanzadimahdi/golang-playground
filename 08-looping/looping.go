package main

import "fmt"

func main() {
	for i, j := 0, 10; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	j := 0
	for {
		if j%2 == 0 {
			continue
		}

		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value, s[key])
	}

}
