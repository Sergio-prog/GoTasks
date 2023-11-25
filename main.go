package main

import (
	"GoTasks/api"
)

func main() {
	// file := ToFile.File{Path: "tasks.json"}
	// data, err := file.GetData()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))

	// err = file.AddData("Test Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	api.Run()
}
