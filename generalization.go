
package wolfram

import (
  "encoding/xml"
)

// Struct for the <generalization> tag.
type Generalization struct {
  XMLName             xml.Name                `xml:"generalization,omitempty"`
  Topic               string                  `xml:"topic,attr"`
  Desc                string                  `xml:"desc,attr"`
  Url                 string                  `xml:"url,attr"`
}

// The interface of the Generalization.
type IGeneralization interface {
  GetTopic()          string
  GetDescription()    string
  GetUrl()            string
}

// Returns the attribute value of 'topic' from the <generalization>.
func (g *Generalization) GetTopic() string {
  return g.Topic
}

// Returns the attribute value of 'desc' from the <generalization>.
func (g *Generalization) GetDescription() string {
  return g.Desc
}

// Returns the attribute value of 'url' from the <generalization>.
func (g *Generalization) GetUrl() string {
  return g.Url
}

