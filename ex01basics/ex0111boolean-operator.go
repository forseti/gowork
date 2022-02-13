package main

import "fmt"

func main() {
	var exam = 80
	var attendance = 80

	var passExam = exam >= 80
	var passAttendance = attendance >= 80
	fmt.Println("passExam =", passExam)
	fmt.Println("passAttendance =", passAttendance)

	var passOverall = passAttendance && passExam
	fmt.Println("passOverall =", passOverall)

	fmt.Println("exam >= 80 && attendance >= 80 =", exam >= 80 && attendance >= 80)
}
