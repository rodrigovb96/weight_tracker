package logger

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"fmt"
)

type DayMeasure struct {
	Weight float32
	Date string
}

type WeekMeasure struct {
	Week int
	Sum float32
	Mean float32
	Days []DayMeasure
}

// If the file already exists 
// We need to update it
// First we need to read it's content and then add the new content
func updateFile(filePath string, content WeekMeasure) bool {

	file, _ := ioutil.ReadFile(filePath)

	// Reading the contents
	data := WeekMeasure{}
	json.Unmarshal([]byte(file), &data)

	fmt.Println(data.Days[0])

	// Adding the new day
	data.Days = append(data.Days,content.Days[0])
	data.Sum += content.Days[0].Weight

	data.Mean = data.Sum / float32(len(data.Days))

	// Starting to rewrite the new file 
	file, _ = json.MarshalIndent(data, "", "")

	_ = ioutil.WriteFile(filePath ,file,0644)


	return true
}

func dataFactory(weight_ float32) (w WeekMeasure, nowTime string) {

	nowTime = time.Now().Format("01-02-2006 15:04:05")

	weekNumber, _ := strconv.Atoi(nowTime[3:5])
	weekNumber /= 7

	w = WeekMeasure{
		Week: weekNumber,
	}

	w.Days = make([]DayMeasure,1)

	w.Days[0] = DayMeasure{
			Weight: weight_,
			Date: nowTime,
		    }

	w.Sum += weight_

	w.Mean = w.Sum


	return

}

// Receives the weight and log to a file 
func LogToFile(weight float32) bool{

	w, nowTime := dataFactory(weight)

	// Path to the month folder
	path := "files/data/"+nowTime[0:2]+"/"

	// The actual file path
	filePath := path + strconv.Itoa(w.Week)+ ".json"

	// If File already exist, we need to update it 
	if _,err := os.Stat(filePath); err == nil {
		return updateFile(filePath,w)
	}

	// If the folder does not exist we need to create it
	if _,err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path,0777)
	}

	file, err := json.MarshalIndent(w, "", "")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filePath ,file,0644)

	return true;
}
