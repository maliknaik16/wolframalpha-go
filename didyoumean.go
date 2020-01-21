
package wolfram

import (
  "encoding/xml"
)

// Struct for the <didyoumean> tag.
type DidYouMean struct {
  XMLName             xml.Name                `xml:"didyoumean,omitempty"`
  Score               string                  `xml:"score,attr"`
  Level               string                  `xml:"level,attr"`
  Value               string                  `xml:",chardata"`
}

// The interface for the DidYouMean
type IDidYouMean interface {
  GetScore()          string
  GetLevel()          string
  GetValue()          string
}

// Returns the "score" attribute value from the <didyoumean>.
func (dym *DidYouMean) GetScore() string {
  return dym.Score
}

// Returns the "level" attribute value from the <didyoumean>.
func (dym *DidYouMean) GetLevel() string {
  return dym.Level
}

// Returns the value of the <didyoumean>.
func (dym *DidYouMean) GetValue() string {
  return dym.Value
}
