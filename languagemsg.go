
package wolfram

import (
  "encoding/xml"
)

// Struct for the <languagemsg> tag.
type LanguageMsg struct {
  XMLName             xml.Name                `xml:"languagemsg,omitempty"`
  English             string                  `xml:"english,attr"`
  Other               string                  `xml:"other,attr"`
}

// The interface for the `LanguageMsg`.
type ILanguageMsg interface {
  GetEnglish()        string
  GetOther()          string
}

// Returns the "english" attribute value from <languagemsg>.
func (l *LanguageMsg) GetEnglish() string {
  return l.English
}

// Returns the "other" attribute value from <languagemsg>.
func (l *LanguageMsg) GetOther() string {
  return l.Other
}
