package main

import "fmt"

func average_calculator(student_info map[string]float64) float64 {
	tot := 0.0
	for _, v := range student_info {
		tot += v
	}
	student_info["total"] = tot
	length := len(student_info)
	average := tot / float64(length)
	return average
}

func display_info(student_name string, average float64, student_info map[string]float64) {
	fmt.Printf("\n---------------CALCULATED RESULT FOR %v----------------\n\n",student_name)
	for k, v := range student_info {
		if k == "total"{
			continue
		}
		fmt.Printf("%v : %0.2f \n",k,v)
	}
	fmt.Printf("%v : %0.2f\n","Total",student_info["total"])

	fmt.Printf("%v : %0.3f\n","Average",average)
	fmt.Println("\n---------------THANKS----------------\n")

}

func accept_info() (string,map[string]float64){
	fmt.Println("---------------WELCOME TO AVERAGE CALCULATOR----------------")
	name:=""
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	subject := 0
	fmt.Print("How much subject do you want to enter: ")
	fmt.Scan(&subject)
	var student_info map[string]float64 = map[string]float64{}
	var sub_name string
	var point float64
	for i:=1;i<subject+1; i++{
		fmt.Printf("Enter Subject %v name: ",i)
		fmt.Scan(&sub_name)
		fmt.Printf("Enter Subject %v point: ",i)
		fmt.Scan(&point)
		if point>100 || point<0{
			fmt.Print("Please Enter correct value (0-100)) for the point again : ")
			fmt.Scan(&point)
		}
		student_info[sub_name] = point
	}
	return name, student_info
}

func main() {
	name, student_info:=accept_info()
	average := average_calculator(student_info)
	display_info(name,average,student_info)

}


/*
		Test Cases

Test Case 1: Normal Case

studentInfo1 := map[string]float64{
  "Math":    85,
  "English": 90,
  "Science": 95,
}

expectedAverage1 := 90.0 // (85 + 90 + 95) / 3

Test Case 2: Single Subject

studentInfo2 := map[string]float64{
  "Math": 100,
}

expectedAverage2 := 100.0 // 100 / 1

Test Case 3: No Subjects

studentInfo3 := map[string]float64{}

expectedAverage3 := 0.0 // Sum is 0 and there are no subjects

*/