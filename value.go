
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <value> tag.
type Value struct {
  XMLName             xml.Name                `xml:"value,omitempty"`
  Name                string                  `xml:"name,attr"`
  Desc                string                  `xml:"desc,attr"`
  Input               string                  `xml:"input,attr"`
  Valid               string                  `xml:"valid,attr"`
  Word                string                  `xml:"word,attr"`
}

// The interface for the Value.
type IValue interface {
  GetName()           string
  GetDescription()    string
  GetInput()          string
  IsValid()           bool
  GetWord()           string
}

// Returns the "name" attribute value from the <value>.
func (v *Value) GetName() string {
  return v.Name
}

// Returns the "desc" attribute value from the <value>.
func (v *Value) GetDescription() string {
  return v.Desc
}

// Returns the "input" attribute value from the <value>.
func (v *Value) GetInput() string {
  return v.Input
}

// Returns the "valid" attribute value from the <value> as boolean.
func (v *Value) IsValid() bool {
  if v.Valid == "" {
    return true
  }

  valid, _ := strconv.ParseBool(v.Valid)

  return valid
}

// Returns the "word" attribute value from the <value>.
func (v *Value) GetWord() string {
  return v.Word
}

