
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <tips> tag.
type Tips struct {
  XMLName             xml.Name                `xml:"tips,omitempty"`
  Count               string                  `xml:"count,attr"`
  Tips                 []*Tip                  `xml:"tip,omitempty"`
}

// Struct for the <tip> tag.
type Tip struct {
  XMLName             xml.Name                `xml:"tip,omitempty"`
  TextAttr            string                  `xml:"text,attr"`
}

// The interface for the `Tips`.
type ITips interface {
  GetCount()          int
  GetTips()           []*Tip
  GetTip(int)         *Tip
}

// The interface for the `Tip`.
type ITip interface {
  GetText()           string
}

// Returns the number of <tip> elements.
func (tips *Tips) GetCount() int {
  c, _ := strconv.Atoi(tips.Count)

  return c
}

// Returns the slice of <tip> elements.
func (tips *Tips) GetTips() []*Tip {
  return tips.Tips
}

// Returns the <tip> element at the given index.
func (tips *Tips) GetTip(index int) *Tip {
  if index >=0 && index < len(tips.Tips) {
    return tips.Tips[index]
  }

  return nil
}

// Returns the "text" attribute value from <tip>.
func (tip *Tip) GetText() string {
  return tip.TextAttr
}
