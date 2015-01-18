package main

import (
	//"fmt"
	"os"
	"log"
	"text/template"
	"io/ioutil"
	"bytes"
)

const stack = `
package stack

type Stack []{{.Type}}

func (s *Stack) Push(v {{.Type}}) (e error) {
	e, *s = Push(*s, v)
	return e
}

func (s *Stack) Pop() (e error, v {{.Type}}) {
	e,*s, v = Pop(*s)
	return e, v
}

func Push(a Stack, v {{.Type}}) (e error, r Stack) {
	if len(a) == cap(a) {
		t := make(Stack, len(a), 2*len(a)+1)
		copy(t, a)
		a = t
	}
	r = a[0 : len(a)+1 : cap(a)]
	r[len(r)-1] = v
	return e, r
}

func Pop(a Stack) (e error, r Stack, v {{.Type}}) {
	var z {{.Type}}	
	v = a[len(a)-1]
	a[len(a)-1] = z
	r = a[0 : len(a)-1 : cap(a)]
	if len(r) <= cap(r)/4 {
		t := make(Stack, len(r), len(r))
		copy(t, r)
		r = t
	}
	return e, r, v
}
`

func main() {
	t := template.Must(template.New("Stack").Parse(stack))
	bb := bytes.NewBuffer(make([]byte,0,0))
	data := struct{ Type string }{Type: "int64"}
	if err := t.Execute(bb, data); err != nil {
		log.Fatal(err)
	}
   	os.Mkdir("." + string(os.PathSeparator) + data.Type ,0644)
	if err := ioutil.WriteFile( data.Type +"/Stack.go",bb.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

}
