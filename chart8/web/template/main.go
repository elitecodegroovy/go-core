package main

import (
	"log"
	"net/http"
	"github.com/elitecodegroovy/go-core/chart8/web/template/tmpl"
)

type User struct {
	Name        string
	City        string
	Nationality string
}

type Skill struct {
	Language string
	Level    string
}

type SkillSets []*Skill



func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.RenderTemplate(w, "templates/index.tmpl", nil)
	if err != nil {
		log.Println(err)
	}
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	userData := &User{Name: "刘继刚", City: "广州", Nationality: "中国"}
	err := tmpl.RenderTemplate(w, "templates/aboutme.tmpl", userData)
	if err != nil {
		log.Println(err)
	}
}

func skillSetHandler(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&Skill{Language: "Golang", Level: "资深"},
		&Skill{Language: "Java", Level: "高级"},
		&Skill{Language: "C", Level: "高级"},
		&Skill{Language: "Python", Level: "高级"}}
	err := tmpl.RenderTemplate(w, "templates/skills.tmpl", skillSets)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	//加载模板
	tmpl.LoadTemplates()

}

func main() {
	server := http.Server{
		Addr: ":9090",
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/aboutme", aboutMeHandler)
	http.HandleFunc("/skills", skillSetHandler)
	log.Println("Listening ...")
	server.ListenAndServe()
}
