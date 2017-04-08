package main

import (
    "fmt"
    "os"
    "text/template"
    "encoding/json"
    "io/ioutil"
)

// var persVita map[string]string

type Jobs struct {
  Heading string `json:"Heading"`
  Body string `json:"Body"`
  Date string `json:"Date"`
}


type OtherTasks struct {
  Heading string `json:"Heading"`
  Topic string `json:"Topic"`
  Url string `json:"Url"`
  Date string `json:"Date"`
}

type Years struct {
    Jobs []Jobs `json:"Jobs"`
    OtherTasks []OtherTasks  `json:"OtherTasks"`
    Date string `json:"Date"`
}

type Items struct {
     Score string `json:"Score"`
     Topic string `json:"Topic"`

}

type SkillSections struct {
    Title string `json:"Title"`
    Items []Items `json:"Items"`
}


type PersVita struct {
    Years []Years `json:"Years"`
    SkillSections []SkillSections `json:"SkillSections"`
}


func main() {
    var persVita = readConfig()

    fmt.Printf("read template...\n")
    t := template.New("html-cv.tmpl")
    t.ParseFiles("html-cv.tmpl")

    fmt.Printf("persVita value: %v\n", persVita)
    fmt.Printf("execute template...\n")
    err := t.Execute(os.Stdout, persVita)
    if err != nil {
        panic(err)
    }

}

func readConfig()(persVita  PersVita){
    fmt.Printf("readConfig...\n")
    // var persVita = PersVita{}
    file, e := ioutil.ReadFile("./vita.json")
    if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
    }

    fmt.Printf("file value: %v\n", string(file))
    json.Unmarshal(file, &persVita)
    fmt.Printf("persVita.Years[0].Date: %v\n", persVita.Years[0].Date)
    return persVita
}
