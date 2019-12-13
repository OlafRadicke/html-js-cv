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


type Parts struct {
    Key string `json:"Key"`
    Value string `json:"Value"`
}

type Projects struct {
    Heading string `json:"Heading"`
    Parts []Parts `json:"Parts"`
    Url string `json:"Url"`
    Date string `json:"Date"`
}

type EducationTrainingConferences struct {
    Heading string `json:"Heading"`
    Parts []Parts `json:"Parts"`
    Url string `json:"Url"`
    Date string `json:"Date"`
}

type PublicationTalks struct {
    Heading string `json:"Heading"`
    Parts []Parts `json:"Parts"`
    Url string `json:"Url"`
    Date string `json:"Date"`
}


// type Years struct {
//     Jobs []Jobs `json:"Jobs"`
//     OtherTasks []OtherTasks  `json:"OtherTasks"`
//     Date string `json:"Date"`
// }

type Items struct {
     Score string `json:"Score"`
     Topic string `json:"Topic"`
}

type SkillSections struct {
    Title string `json:"Title"`
    Items []Items `json:"Items"`
}


type PersVita struct {
    Jobs []Jobs `json:"Jobs"`
    Projects []Projects `json:"Projects"`
    EducationTrainingConferences []EducationTrainingConferences `json:"EducationTrainingConferences"`
    PublicationTalks []PublicationTalks `json:"PublicationTalks"`
    SkillSections []SkillSections `json:"SkillSections"`
}


func main() {
    var tempFile = "cv-tmpl.html"
    var persVita = readConfig()
    if _, err := os.Stat( tempFile ); os.IsNotExist(err) {
        fmt.Printf("Error: template file %v not exist!\n",  tempFile)
        os.Exit(1)
    }
    t := template.New(tempFile)
    t.ParseFiles(tempFile)
    err := t.Execute(os.Stdout, persVita)
    if err != nil {
        fmt.Printf("Panic: %v \n",  err)
        panic(err)
    }
}

func readConfig()(persVita  PersVita){
    file, e := ioutil.ReadFile("./vita.json")
    if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
    }
    // fmt.Printf("file: %v\n", string(file))
    e = json.Unmarshal(file, &persVita)
    if e != nil {
      fmt.Printf("json.Unmarshal error: %v\n", e)
      os.Exit(1)
    }
    return persVita
}
