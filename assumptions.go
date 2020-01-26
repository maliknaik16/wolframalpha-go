
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <assumptions> tag.
type Assumptions struct {
  XMLName               xml.Name              `xml:"assumptions,omitempty"`
  Count                 string                `xml:"count,attr"`
  Assumptions           []*Assumption         `xml:"assumption,omitempty"`
}

// The interface for the Assumptions
type IAssumptions interface {
  GetCount()            int
  GetAssumptions()      []*Assumption
  GetAssumption(int)    *Assumption
}

// Returns the number of <assumption> elements.
func (a *Assumptions) GetCount() int {
  c, _ := strconv.Atoi(a.Count)

  return c
}

// Returns the slice of pointers to `Assumption`.
func (a *Assumptions) GetAssumptions() []*Assumption {
  if a.Assumptions != nil {
    return a.Assumptions
  }

  return nil
}

// Returns the pointer to `Assumption` at the give index.
func (a *Assumptions) GetAssumption(index int) *Assumption {
  if index >= 0 && index < len(a.Assumptions) - 1 {
    return a.Assumptions[index]
  }

  return nil
}
