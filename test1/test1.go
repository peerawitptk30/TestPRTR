package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func findtheways(findtheways [][]int) int {

	sum := 0
	// 
	for _, sublist := range findtheways {
		//
		for _, num := range sublist {
			sum += num
		}
	}
	return sum
}

func main() {
	// อ่าน JSON ไฟล์
	data, err := ioutil.ReadFile("hard.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    //กำหนดค่าเพื่อเก็บ LIST ข้อมูล
    var nestedList [][]int

    // แปลงค่าจาก json file ให้เป็น list ข้อมูล
    err = json.Unmarshal(data, &nestedList)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }
	// เรียกใช้ function findtheways เพิ่อหาค่าที่มากที่สุด
    result := findtheways(nestedList)

    fmt.Println("ค่าที่มากที่สุด:", result)
}