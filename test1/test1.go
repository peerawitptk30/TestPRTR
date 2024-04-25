package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxPathSum(pathmostsum [][]int) int {
	a := make([]int, len(pathmostsum))
	copy(a, pathmostsum[len(pathmostsum)-1])
	for i := len(pathmostsum) - 2; i >= 0; i-- {
		for j := 0; j < len(pathmostsum[i]); j++ {
			a[j] = pathmostsum[i][j] + max(a[j], a[j+1])
		}
	}
	return a[0]
}

func main() {
	// อ่าน JSON ไฟล์
	test1 := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	
	data, err := ioutil.ReadFile("hard.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    //กำหนดค่าเพื่อเก็บ LIST ข้อมูล
    var test2 [][]int

    // แปลงค่าจาก json file ให้เป็น list ข้อมูล
    err = json.Unmarshal(data, &test2)
	// fmt.Println(nestedList)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }
	// เรียกใช้ function findtheways เพิ่อหาค่าที่มากที่สุด
	result1 := maxPathSum(test1)
    result2 := maxPathSum(test2)
	fmt.Println("เส้นทางที่มีค่ามากที่สุด TEST1 =", result1)
    fmt.Println("เส้นทางที่มีค่ามากที่สุด TEST2 =", result2)
}