package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "description": "Monotremata", "age": 18},
		{"Name": "Quoll",    "description": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string `json:"description"`
		Age   int
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%# v", animals)

	m := DecodeJson(exampleJson)
	fmt.Printf("%# v", m)
	fmt.Println()
	fmt.Println(m.Menu.ID)
	fmt.Println(m.Menu.Popup)

	fmt.Printf("%# v\n", DecodeXML(exampleXML))

	fmt.Println("---------------")
	fmt.Println(DecodeJson(exampleJson))
	fmt.Println(DecodeXML(exampleXML))
	fmt.Println("---------------")

	x := DecodeXML(exampleXML)
	s := encodeJson(Menu{Menu: x})
	fmt.Println(s)
}

var exampleJson = `{"menu": {
  "id": "file",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}`

type Menu struct {
	Menu MenuX
}

type MenuX struct {
	XMLName xml.Name `xml:"menu" json:"-"`
	ID      string   `xml:"id,attr" json:"id"`
	Value   string   `xml:"value,attr" json:"value"`
	Popup   struct {
		Menuitems []struct {
			Value   string `xml:"value,attr" json:"value"`
			Onclick string `xml:"onclick,attr" json:"onclick"`
		} `xml:"menuitem" json:"menuitem"`
	} `xml:"popup" json:"popup"`
}

func DecodeJson(s string) Menu {
	m := Menu{}
	err := json.Unmarshal([]byte(s), &m)

	if err != nil {
		fmt.Println(err)
	}
	return m
}

var exampleXML = `<menu id="file" value="File">
  <popup>
    <menuitem value="New" onclick="CreateNewDoc()" />
    <menuitem value="Open" onclick="OpenDoc()" />
    <menuitem value="Close" onclick="CloseDoc()" />
  </popup>
</menu>`

func DecodeXML(s string) MenuX {
	x := MenuX{}
	err := xml.Unmarshal([]byte(s), &x)
	if err != nil {
		fmt.Println(err)
	}
	return x
}

func encodeJson(m Menu) string {
	b, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
