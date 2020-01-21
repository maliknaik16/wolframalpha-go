
package wolfram

import (
  "encoding/xml"
)

// Struct for the <error> tag.
type ErrorTag struct {
  XMLName             xml.Name                `xml:"error,omitempty"`
  Code                *Code                   `xml:"code,omitempty"`
  Msg                 *Msg                    `xml:"msg,omitempty"`
}

// Struct for the <code> tag.
type Code struct {
  XMLName             xml.Name                `xml:"code,omitempty"`
  Code                string                  `xml:",chardata"`
}

// Struct for the <msg> tag.
type Msg struct {
  XMLName             xml.Name                `xml:"msg,omitempty"`
  Msg                 string                  `xml:",chardata"`
}

// The interface for the ErrorTag.
type IErroTag interface {
  GetCode()           int
  GetMsg()            string
}

// Returns the error code from the <queryresult>.
func (e *ErrorTag) GetCode() int {
  c, _ := strconv.Atoi(e.Code.Code)

  return c
}

// Returns the error message from the <queryresult>.
func (e *ErrorTag) GetMsg() *Msg {
  return e.Msg.Msg
}

