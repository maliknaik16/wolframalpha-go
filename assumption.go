
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <assumption> tag.
type Assumption struct {
  XMLName             xml.Name                `xml:"assumption,omitempty"`
  Type                string                  `xml:"type,attr"`
  Word                string                  `xml:"word,attr"`
  Template            string                  `xml:"template,attr"`
  Current             string                  `xml:"current,attr"`
  Count               string                  `xml:"count,attr"`
  Values              []*Value                `xml:"value,omitempty"`
}

// The interface for the Assumption.
type IAssumption interface {
  GetType()           string
  GetCount()          int
  GetWord()           string
  GetTemplate()       string
  GetCurrent()        int
  GetValue(int)       *Value

  GetValues()         []*Value
  GetNames()          []string
  GetDescriptions()   []string
  GetInputs()         []string
  GetWords()          []string
  GetValidities()     []bool
}

// Returns the "type" attribute value of the assumption(E.g: Clash, MultiClash, .etc).
func (a *Assumption) GetType() string {
  return a.Type
}

// Returns the number of <value> elements.
func (a *Assumption) GetCount() int {
  count, _ := strconv.Atoi(a.Count)

  return count
}

// Returns the "word" attribute value of the assumption.
func (a *Assumption) GetWord() string {
  return a.Word
}

// Returns the "template" attribute value of the assumption.
func (a *Assumption) GetTemplate() string {
  return a.Template
}

// Returns the "current" attribute value of the assumption.
func (a *Assumption) GetCurrent() int {
  current, _ := strconv.Atoi(a.Current)

  return current
}

// Returns the *Value at the given index.
func (a *Assumption) GetValue(index int) *Value {
  if index >= 0 && index < len(a.Values) - 1 {
    return a.Values[index]
  }

  return nil
}

// Returns the slice of *Value
func (a *Assumption) GetValues() []*Value {
  if a.Values != nil {
    return a.Values
  }

  return nil
}

// Returns the slice of values of the "name" attribute of the <value> element
// from the <assumption>.
func (a *Assumption) GetNames() []string {
  length := a.GetCount()
  names := make([]string, length)

  for i := 0; i < length; i++ {
    names[i] = a.GetValue(i).GetName()
  }

  return names
}

// Returns the slice of values of the "desc" attribute of the <value> element
// from the <assumption>.
func (a Assumption) GetDescriptions() []string {
  length := a.GetCount()
  descriptions := make([]string, length)

  for i := 0; i < length; i++ {
    descriptions[i] = a.GetValue(i).GetDescription()
  }

  return descriptions
}

// Returns the slice of values of the "input" attribute of the <value> element
// from the <assumption>.
func (a Assumption) GetInputs() []string {
  length := a.GetCount()
  inputs := make([]string, length)

  for i := 0; i < length; i++ {
    inputs[i] = a.GetValue(i).GetInput()
  }

  return inputs
}
// Returns the slice of values of the "word" attribute of the <value> element
// from the <assumption>.
func (a Assumption) GetWords() []string {
  length := a.GetCount()
  words := make([]string, length)

  for i := 0; i < length; i++ {
    words[i] = a.GetValue(i).GetWord()
  }

  return words
}

// Returns the slice of values of the "valid" attribute of the <value> element
// from the <assumption>.
func (a Assumption) GetValidities() []bool {
  length := a.GetCount()
  valids := make([]bool, length)

  for i := 0; i < length; i++ {
    valids[i] = a.GetValue(i).IsValid()
  }

  return valids
}
