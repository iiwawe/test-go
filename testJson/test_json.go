package testJson

import (
	"testing"
	"fmt"
	"bonc.com/testGo/testJson/entity"
	"encoding/json"
	"os"
	"log"
)

func TestMarshal(t *testing.T) {
	st := testJson.NewStudent("Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99, )
	strStu, err := json.Marshal(st)
	if err != nil {
		return
	}
	fmt.Println(string(strStu[:]))

	stu := &testJson.Student{}
	err = json.Unmarshal(strStu, stu)

	if err != nil {
		return
	}
	fmt.Println(stu.Name)
}

func TestDecoder(t *testing.T) {
	file, err := os.Open("D://json.txt")
	defer file.Close()
	file2, err := os.Create("D://jsonOut.txt")
	defer file2.Close()
	if err != nil {
		log.Println("decode err")
		return
	}
	dec := json.NewDecoder(file)
	enc := json.NewEncoder(file2)
	for {
		var kvs map[string]interface{}
		if err := dec.Decode(&kvs); err != nil {
			log.Println(err)
			return
		}
		for k := range kvs {
			log.Printf("%s: %s\n", k, kvs[k])
			//if k != "Name" {
			//	delete(kvs, k)
			//}
		}
		if err := enc.Encode(&kvs); err != nil {
			log.Println(err)
		}
	}

}

