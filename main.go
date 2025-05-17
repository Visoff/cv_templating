package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Data struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    ImageURL  string `json:"image_url"`
    Title     string `json:"title"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Website   string `json:"website"`
    Summary   string `json:"summary"`
    Languages []struct {
        Name string `json:"name"`
        Level int    `json:"level"`
    } `json:"languages"`
    Skills    []struct {
        Name   string `json:"name"`
        Level  int    `json:"level"`
    } `json:"skills"`
    Experience []struct {
        Position     string `json:"position"`
        Company      string `json:"company"`
        StartDate    string `json:"start_date"`
        EndDate      string `json:"end_date"`
        Description  []string `json:"description"`
    } `json:"experience"`
    Education []struct {
        Degree       string `json:"degree"`
        FieldOfStudy string `json:"field_of_study"`
        Institution  string `json:"institution"`
        StartDate    string `json:"start_date"`
        EndDate      string `json:"end_date"`
    } `json:"education"`
    Projects []struct {
        Name        string `json:"name"`
        Description []string `json:"description"`
        URL        string `json:"url"`
        Technologies []string `json:"technologies"`
    } `json:"projects"`

	GitHub struct {
		Username string `json:"username"`
	} `json:"github"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var t *template.Template
		var err error
		if r.URL.Path == "/md" {
			t, err = template.ParseFiles("profile.md")
		} else {
			t, err = template.ParseFiles("index.html")
		}
        if err != nil {
            panic(err)
    }
        fp, err := os.Open("data.json")
        if err != nil {
            panic(err)
        }
        defer fp.Close()
        var data Data
        data_bytes, err := io.ReadAll(fp)
        if err != nil {
            panic(err)
        }
        json.Unmarshal(data_bytes, &data)
        err = t.Execute(w, data)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
        }
    })

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.ListenAndServe(":8080", nil)
}
