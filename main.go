package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type myData struct {
	name string
	age  int
}

func main() {
	var buf bytes.Buffer
	var data = []myData{
		{name: "Sandra Juarez", age: 19},
		{name: "Marco Dias", age: 21},
		{name: "Jorge Alvarez", age: 15},
	}

	for i, v := range data {
		var aux = []byte(v.name)
		var auxAge = []byte(strconv.Itoa(v.age))
		aux = append(aux, ' ')
		aux = append(aux, auxAge...)

		if i < len(data)-1 {
			aux = append(aux, '\n')
		}

		var n, err = buf.Write(aux)
		if err != nil || n != len(aux) {
			break
		}
	}
	fmt.Println(buf.String())
}
