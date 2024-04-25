package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type BeefSummary struct {
	TBone    int `json:"t-bone"`
	Fatback  int `json:"fatback"`
	Pastrami int `json:"pastrami"`
	Pork     int `json:"pork"`
	Meatloaf int `json:"meatloaf"`
	Jowl     int `json:"jowl"`
	Enim     int `json:"enim"`
	Bresaola int `json:"bresaola"`
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
    // เรียกใช้งาน API เพื่อรับข้อมูลจากอินเทอร์เน็ต
    response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
    if err != nil {
        http.Error(w, "Failed to fetch data from external API", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    // อ่านข้อมูลจาก response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        http.Error(w, "Failed to read response body", http.StatusInternalServerError)
        return
    }

    // แยกคำ
    words := strings.Fields(string(body))
    
    // สร้างตัวแปรเพื่อเก็บค่าการนับ
    beefCounts := BeefSummary{}


    var wg sync.WaitGroup

    var mutex sync.Mutex


    for _, word := range words {
        wg.Add(1)
        go func(word string) {
            defer wg.Done()


            mutex.Lock()
            defer mutex.Unlock()
            switch word {
            case "t-bone", "T-bone":
                beefCounts.TBone++
            case "fatback", "Fatback":
                beefCounts.Fatback++
            case "pastrami", "Pastrami":
                beefCounts.Pastrami++
            case "pork", "Pork":
                beefCounts.Pork++
            case "meatloaf":
                beefCounts.Meatloaf++
            case "jowl", "Jowl":
                beefCounts.Jowl++
            case "enim", "Enim":
                beefCounts.Enim++
            case "bresaola", "Bresaola":
                beefCounts.Bresaola++
            }
        }(word)
    }

    // รอให้การทำงานทั้งหมดเสร็จสมบูรณ์
    wg.Wait()

    // สร้าง JSON response
    respData := map[string]BeefSummary{"beef": beefCounts}
    respJSON, err := json.Marshal(respData)
    if err != nil {
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
