
package wolfram

import (
  "encoding/xml"
)

// Struct for the <examplepage> tag.
type ExamplePage struct {
  XMLName             xml.Name                `xml:"examplepage,omitempty"`
  Category            string                  `xml:"category,attr"`
  Url                 string                  `xml:"url,attr"`
}

// Returns the value of the "category" attribute from <examplepage>.
func (e *ExamplePage) GetCategory() string {
  return e.Category
}

// Returns the value of the "url" attribute from <examplepage>.
func (e *ExamplePage) GetUrl() string {
  return e.Url
}
