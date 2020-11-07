package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestCheckType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `json:"name"` //Struct Tag
	Age        int    `json:"age"`  //Struct Tag
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name:value(%[1]v),Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:json", nameField.Tag.Get("json"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(10)})
	t.Log("Update Age:", e)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "four"}

	// fmt.Println(a == b) // map can only be compared to nil
	fmt.Println(reflect.DeepEqual(a, b)) // 输出 false

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	// fmt.Println(s1 == s2) //slice can only be compared to nil
	fmt.Printf("s1 == s2 ? %t\n", reflect.DeepEqual(s1, s2)) // 输出 s1 == s2 ? true
	fmt.Printf("s2 == s3 ? %t\n", reflect.DeepEqual(s2, s3)) // 输出 s2 == s3 ? false

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(reflect.DeepEqual(c1, c2)) // 输出 true
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem 获得指针指向的结构
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type.")
		}
	}
	if settings == nil {
		return errors.New("the setting should not be nil.")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 20}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
