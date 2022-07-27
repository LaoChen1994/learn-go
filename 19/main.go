package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type hr struct {
	Name   string `json:name`
	Age    int    `json:age`
	Gender int    `json:gender`
}

type request struct {
	Code int32  `json: code`
	Data []hr   `json: data`
	Msg  string `json: msg`
}

func main() {
	var db []hr

	var people hr

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "欢迎使用人力资源管理系统")
		w.Write([]byte("欢迎光临"))
	})

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			err := r.ParseForm()

			if err != nil {
				fmt.Println("err ->", err)
			} else {
				name := r.FormValue("name")
				age, _ := strconv.Atoi(r.FormValue("age"))
				gender, _ := strconv.Atoi(r.FormValue("gender"))
				fmt.Println(name, age, gender)

				db = append(db, hr{
					Name:   name,
					Age:    age,
					Gender: gender,
				})

				fmt.Fprintf(w, "添加了"+name)
			}
		}

		if r.Method == "POST" {
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)

			if err != nil {
				fmt.Fprintln(w, err)
				panic(err)
			}

			jsonErr := json.Unmarshal(body, &people)

			if jsonErr != nil {
				fmt.Fprintln(w, jsonErr)
				panic(jsonErr)
			}

			db = append(db, hr{
				Name:   people.Name,
				Age:    people.Age,
				Gender: people.Gender,
			})

			fmt.Fprintf(w, "插入了"+people.Name)

		}
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			b, err := json.Marshal(request{
				Msg:  "ok",
				Code: 200,
				Data: db,
			})

			fmt.Println(b)

			if err != nil {
				fmt.Fprintln(w, err)
				panic(err)
			}

			w.Header().Set("Content-Type", "text/json")
			w.Write(b)
		}
	})

	err := http.ListenAndServe(":8208", nil)

	if err != nil {
		fmt.Println("err ->", err)
	}
}
