package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type address struct {
	Str string
	No  int
}

type phone struct {
	N1 int
}

type user struct {
	Name  string    `json:"name,omitempty"`
	ID    int       `json:"id,omitempty"`
	Add   []address `json:"add,omitempty"`
	phone `json:"phone,omitempty"`
}

func main() {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}
	http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("")
		fmt.Println(r.Method+" ", r.Host+" ", r.URL)
		fmt.Println("==========Request==========")
		for key, value := range r.Header {
			fmt.Println(key+":", value)
		}

		header := w.Header()
		header.Add("Action", "query")
		header.Add("Content-Type", "application/json; charset=utf-8")

		fmt.Println("==========Response==========")
		for key, value := range header {
			fmt.Println(key+":", value)
		}

		u := user{
			Name: "lyl",
			ID:   30,
			Add:  []address{address{"shanghai", 1035}},
		}
		bs, _ := json.Marshal(u)
		w.Write(bs)

		//w.WriteHeader(200)
	}))
}
