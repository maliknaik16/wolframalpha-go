
package wolfram

import (
  "encoding/xml"
)

// Struct for the <futuretopic> tag.
type FutureTopic struct {
  XMLName             xml.Name                `xml:"futuretopic,omitempty"`
  Topic               string                  `xml:"topic,attr"`
  Msg                 string                  `xml:"msg,attr"`
}

// Returns the value of the "msg" attribute from <futuretopic>.
func (f *FutureTopic) GetMessage() string {
  return f.Msg
}

// Returns the value of the "topic" attribute from <futuretopic>.
func (f *FutureTopic) GetTopic() string {
  return f.Topic
}

