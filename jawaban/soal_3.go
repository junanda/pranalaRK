package jawaban

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type DataRead struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

func JsonReadModified(filePath string) {
	var (
		dataRead DataRead
	)

	dataFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Faiel Read File JOSN: ", err.Error())
	}

	err = json.Unmarshal(dataFile, &dataRead)
	if err != nil {
		log.Println("Error Unmarshal data:", err.Error())
	}

	dataRead.Email = "johndoe@example.com"
	dataRead.Age += 1

	dataNew, err := json.Marshal(dataRead)
	if err != nil {
		log.Println("error marshal data:", err.Error())
	}

	err = ioutil.WriteFile(filePath, dataNew, 0655)
	if err != nil {
		log.Println("Error Create File: ", err.Error())
	}

	log.Println("Data Json: ", dataRead)
}
