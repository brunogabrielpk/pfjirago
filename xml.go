package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Workflow struct {
	XMLName        xml.Name `xml:"workflow"`
	Meta           []Meta   `xml:"meta"`
	InitialActions []Action `xml:"initial-actions>action"`
	GlobalActions  []Action `xml:"global-actions>action"`
	CommonActions  []Action `xml:"common-actions>action"`
  Steps          []Step   `xml:"steps>step"`
}

type Step struct {
  XMLName xml.Name `xml:"step"`
  Name string `xml:"name,attr"`
  Actions []Action `xml:"actions>action"`
}

type Meta struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Action struct {
	ID         string      `xml:"id,attr"`
	Name       string      `xml:"name,attr"`
	View       string      `xml:"view,attr"`
	Meta       []Meta      `xml:"meta"`
	Validators []Validator `xml:"validators>validator"`
	Results    []Result    `xml:"results>unconditional-result"`
}

type Validator struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Args []Arg  `xml:"arg"`
}

type Arg struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Result struct {
	OldStatus string `xml:"old-status,attr"`
	Status    string `xml:"status,attr"`
	Step      string `xml:"step,attr"`
	PostFuncs []Func `xml:"post-functions>function"`
}

type Func struct {
	Type string `xml:"type,attr"`
	Args []Arg  `xml:"arg"`
}

func main() {
	// Read the XML code from a file
	xmlFile, err := os.Open("./OMNI-JSM.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	// Read the XML data into a byte slice
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the XML data into a Workflow struct
	var workflow Workflow
	if err := xml.Unmarshal(xmlData, &workflow); err != nil {
		log.Fatal(err)
	}

	// Print the workflow data
	//for _, meta := range workflow.Meta {
	//fmt.Println(meta.Name, meta.Value)
	//}

	//for _, action := range workflow.InitialActions {
	//fmt.Println(action.ID, action.Name, action.View)
	//for _, meta := range action.Meta {
	//fmt.Println(meta.Name, meta.Value)
	//}
	//for _, validator := range action.Validators {
	//fmt.Println(validator.Name, validator.Type)
	//for _, arg := range validator.Args {
	//fmt.Println(arg.Name, arg.Value)
	//}
	//}
	//for _, result := range action.Results {
	//fmt.Println(result.OldStatus, result.Status, result.Step)
	//for _, postFunc := range result.PostFuncs {
	//fmt.Println(postFunc.Type)
	//for _, arg := range postFunc.Args {
	//fmt.Println(arg.Name, arg.Value)
	//}
	//}
	//}
	//}

	//printing the postfunctions from InitialActions
	for _, action := range workflow.InitialActions {
    fmt.Println("#######################")
    fmt.Println("ACTION NAME : ", action.Name)
		for _, result := range action.Results {
			for _, postFunc := range result.PostFuncs {
				fmt.Println(postFunc.Type)
				for _, arg := range postFunc.Args {
					fmt.Println(arg.Name, arg.Value)
				}
			}
		}
	}

  //printing the postfunctions from GlobalActions
  for _, action := range workflow.GlobalActions {
    fmt.Println("#######################")
    fmt.Println("ACTION NAME : ", action.Name)
    for _, result := range action.Results {
      for _, postFunc := range result.PostFuncs {
        fmt.Println(postFunc.Type)
        for _, arg := range postFunc.Args {
          fmt.Println(arg.Name, arg.Value)
        }
      }
    }
  }

  //printing the postfunctions from CommonActions
  for _, action := range workflow.CommonActions {
    fmt.Println("#######################")
    fmt.Println("ACTION NAME : ", action.Name)
    for _, result := range action.Results {
      for _, postFunc := range result.PostFuncs {
        fmt.Println(postFunc.Type)
        for _, arg := range postFunc.Args {
          fmt.Println(arg.Name, arg.Value)
        }
      }
    }
  }

  //printing the postfunctions from steps
  for _, step := range workflow.Steps {
    fmt.Println("#######################")
    fmt.Println("STEP NAME : ", step.Name)
    for _, action := range step.Actions {
      fmt.Println("ACTION NAME : ", action.Name)
      for _, result := range action.Results {
        for _, postFunc := range result.PostFuncs {
          fmt.Println(postFunc.Type)
          for _, arg := range postFunc.Args {
            fmt.Println(arg.Name, arg.Value)
          }
        }
      }
    }
  }

}
