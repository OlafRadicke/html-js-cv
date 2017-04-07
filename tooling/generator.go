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
}


type OtherTasks struct {
  Heading string `json:"Heading"`
  Topic string `json:"Topic"`
  Url string `json:"Url"`
}

type Years struct {
    Jobs []Jobs `json:"Jobs"`
    OtherTasks []OtherTasks  `json:"OtherTasks"`
    Date string `json:"Date"`
}

type SkillSections struct {
}

type PersVita struct {
    Years []Years `json:"Years"`
    SkillSections []SkillSections `json:"SkillSections"`
}
  // "Years": [
  //           {
  //             "Jobs": [
  //                   { "Heading": "Job", "Body": "Entwickler"},
  //                   { "Heading": "Job", "Body": "Tester"}
  //             ],
  //             "OtherTasks": [
  //                   { "Heading": "Job", "Topic": "Einfürung in Git", "Uri": "https://exta.org"},
  //                   { "Heading": "Job", "Topic": "Einfürung in Git", "Uri": "https://exta.org"}
  //             ],
  //             "Date": "2016"
  //           },
  //           {
  //             "Jobs": [
  //                   { "Heading": "Job", "Body": "Entwickler"},
  //                   { "Heading": "Job", "Body": "DevOps"}
  //             ],
  //             "OtherTasks": [
  //                   { "Heading": "Job", "Topic": "Einfürung in Git", "Uri": "https://exta.org"},
  //                   { "Heading": "Job", "Topic": "Einfürung in Git", "Uri": "https://exta.org"}
  //             ],
  //             "Date": "2015"}
  //           ],
  // "SkillSections": [
  //                   { "Title": "Entwicklertools",
  //                     "Items": [
  //                       { "Score": "30", "Topic": "Go: Basics" },
  //                       { "Score": "60", "Topic": "Git: Gut" }'
  //                     ]
  //                   },
  //                   { "Title": "CD",
  //                     "Items": [
  //                       { "Score": "30", "Topic": "Jenkins: Basics" },
  //                       { "Score": "60", "Topic": "GitLab-CI: Gut" }'
  //                     ]
  //                   }
  //                 ]
// }


func main() {
    readConfig()
    t := template.New("html-cv.tmpl")
    t.ParseFiles("html-cv.tmpl")
    err := t.Execute(os.Stdout, PersVita)
    if err != nil {
        panic(err)
    }

}


func readConfig(){
  file, e := ioutil.ReadFile("./vita.json")
  if e != nil {
    fmt.Printf("File error: %v\n", e)
    os.Exit(1)
  }
  json.Unmarshal(file, &persVita)
}
