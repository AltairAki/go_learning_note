package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
 "basic_info":{
	"name":"Mike",
	"age":20
 },
 "job_info":{
	 "skills":["Java","Go","C"]
 }
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e) // {{Mike 20} {[Java Go C]}}
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v)) // {"basic_info":{"name":"Mike","age":20},"job_info":{"skills":["Java","Go","C"]}}
	} else {
		t.Error(err)
	}
}
