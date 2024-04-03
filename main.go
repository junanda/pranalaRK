package main

import "junanda/prk_test/jawaban"

func main() {
	/*
		jawaban soal 3
	*/
	// namaFile := "dataDummy.json"
	// jawaban.JsonReadModified(namaFile)

	/*
		jawaban soal no.4
	*/
	// jawaban.Run()

	/*
		jawaban soal no.5
	*/
	db := jawaban.Conenct()

	db.AutoMigrate(&jawaban.User{})

	// data := jawaban.User{
	// 	Name:  "junanda",
	// 	Email: "jpatihullah@gmail.com",
	// 	Age:   "35",
	// }

	// jawaban.CreateData(data, db)

	jawaban.ReadData("jpatihullah@gmail.com", db)
}
