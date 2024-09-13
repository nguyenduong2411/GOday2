package main

import "fmt"

type StructSv struct {
	ID     int
	Name   string
	Age    int
	Gender string
}

func main() {
	SinhVienMap := make(map[string]StructSv)
	SinhVienMap["SV1"] = StructSv{ID: 1, Name: "Nguyen Van A", Age: 20, Gender: "Nam"}
	SinhVienMap["SV2"] = StructSv{ID: 2, Name: "Nguyen Van B", Age: 21, Gender: "Nu"}
	SinhVienMap["SV3"] = StructSv{ID: 3, Name: "Nguyen Van C", Age: 22, Gender: "Nam"}
	for id, sv := range SinhVienMap {
		fmt.Printf("ID:%s, Name:%s, Age:%d,Gender:%s", id, sv.Name, sv.Age, sv.Gender)
	}
	//insert
	newID := "SV4"
	newSinhvienmap := StructSv{ID: 4, Name: "Nguyen Van D", Age: 23, Gender: "Nu"}
	SinhVienMap[newID] = newSinhvienmap
	//update
	updatedID := "SV2"
	updatedSinhvienmap := StructSv{ID: 2, Name: "Nguyen Van B", Age: 25, Gender: "Nu"}
	SinhVienMap[updatedID] = updatedSinhvienmap
	//delete
	deleteID := "SV3"
	delete(SinhVienMap, deleteID)

	printSinhvienByGender := func(gender string) {
		fmt.Printf("Sinh vien co gioi tinh '%s':\n", gender)
		for id, sv := range SinhVienMap {
			if sv.Gender == gender {
				fmt.Printf("ID:%s, Name:%s, Age:%d\n", id, sv.Name, sv.Age)
			}
		}
	}
	printSinhvienByGender("Nam")
	printSinhvienByGender("Nu")
}
