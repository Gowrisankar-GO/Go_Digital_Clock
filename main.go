package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	zero = []string{"███",
		            "█ █",
		            "█ █",
		            "█ █",
		            "███"}

	one = []string{"██ ",
		           " █ ",
		           " █ ",
		           " █ ",
		           "███"}

	two = []string{"███",
		           "  █",
		           "███",
		           "█  ",
		           "███"}

	three = []string{"███",
		             "  █",
		             "███",
		             "  █",
		             "███"}

	four = []string{"█ █",
		            "█ █",
		            "███",
		            "  █",
		            "  █"}

	five = []string{"███",
		            "█  ",
		            "███",
		            "  █",
		            "███"}

	six = []string{"███",
		           "█  ",
		           "███",
		           "█ █",
		           "███"}

	seven = []string{"███",
		             "  █",
		             "  █",
		             "  █",
		             "  █"}

	eight = []string{"███",
		             "█ █",
		             "███",
		             "█ █",
		             "███"}

	nine = []string{"███",
		            "█ █",
		            "███",
		            "  █",
		            "███"}

	separator = []string{"   ",
		                 " ░ ",
		                 "   ",
		                 " ░ ",
		                 "   "}

	digits = [][]string{zero, one, two, three, four, five, six, seven, eight, nine}
)

func main() {
	for {
		fmt.Println("\r")
		var (
			currentHour          = strconv.Itoa(time.Now().Hour())
			currentMinute        = strconv.Itoa(time.Now().Minute())
			currentSecond        = strconv.Itoa(time.Now().Second())
			hour, minute, second string
		)

		if time.Now().Hour() < 10 {
			hour = "0" + currentHour
		} else {
			hour = currentHour
		}

		if time.Now().Minute() < 10 {
			minute = "0" + currentMinute
		} else {
			minute = currentMinute
		}

		if time.Now().Second() < 10 {
			second = "0" + currentSecond
		} else {
			second = currentSecond
		}

		var (
			timeArr [][]string
			h0, _   = strconv.Atoi(string(hour[0]))
			h1, _   = strconv.Atoi(string(hour[1]))
			m0, _   = strconv.Atoi(string(minute[0]))
			m1, _   = strconv.Atoi(string(minute[1]))
			s0, _   = strconv.Atoi(string(second[0]))
			s1, _   = strconv.Atoi(string(second[1]))
		)
		timeArr = append(timeArr, digits[h0], digits[h1], separator, digits[m0], digits[m1], separator, digits[s0], digits[s1])

		i, rows := 0, 5
		for i < rows {
		loop2:
			for _, digit := range timeArr {
				for index, part := range digit {
					if index == i {
						fmt.Printf("%5s ", part)
						continue loop2
					}
				}
			}
			if i != 4 {
				fmt.Println()
			}
			i++
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\033[H")
	}
}
