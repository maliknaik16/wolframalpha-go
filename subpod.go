
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <subpod> tag.
type SubPod struct {
  XMLName               xml.Name              `xml:"subpod,omitempty"`
  Title                 string                `xml:"title,attr"`
  Image                 *Image                `xml:"img,omitempty"`
  ImageMap              *ImageMap             `xml:"imagemap,omitempty"`
  PlainText             *PlainText            `xml:"plaintext,omitempty"`
  Sound                 *Sound                `xml:"sound,omitempty"`
  Minput                *Minput               `xml:"minput,omitempty"`
  Moutput               *Moutput              `xml:"moutput,omitempty"`
  Cell                  *Cell                 `xml:"cell,omitempty"`
  States                *States               `xml:"states,omitempty"`
}

// Struct for the <plaintext> tag.
type PlainText struct {
  XMLName               xml.Name              `xml:"plaintext,omitempty"`
  PlainText             string                `xml:",chardata"`
}

// Struct for the <minput> tag.
type Minput struct {
  XMLName               xml.Name              `xml:"minput,omitempty"`
  Minput                string                `xml:",chardata"`
}

// Struct for the <moutput> tag.
type Moutput struct {
  XMLName               xml.Name              `xml:"moutput,omitempty"`
  Moutput               string                `xml:",chardata"`
}

// Struct for the <cell> tag.
type Cell struct {
  XMLName               xml.Name              `xml:"cell,omitempty"`
  Compressed            string                `xml:"compressed,attr"`
  CellExpr              string                `xml:",chardata"`
}

// The interface for the `SubPod`.
type ISubPod interface {
  GetTitle()            string
  GetImage()            *Image
  GetImageMap()         *ImageMap
  GetPlainText()        string
  GetSound()            *Sound
  GetMinput()           string
  GetMoutput()          string
  GetCell()             *Cell
  GetStates()           *States
  HasCell()             bool
  GetCellExpression()   string
}

// The interface for the `Cell`.
type ICell interface {
  IsCompressed()        bool
  GetValue()            string
}

// Returns the "title" attribute value from the <subpod>.
func (s *SubPod) GetTitle() string {
  return s.Title
}

// Returns the pointer to `Image`.
func (s *SubPod) GetImage() *Image {
  if s.Image != nil {
    return s.Image
  }

  return nil
}

// Returns the pointer to `ImageMap`.
func (s *SubPod) GetImageMap() *ImageMap {
  if s.ImageMap != nil {
    return s.ImageMap
  }

  return nil
}

// Returns the chardata from <plaintext> as string.
func (s *SubPod) GetPlainText() string {
  return s.PlainText.PlainText
}

// Returns the pointer to `Sound`.
func (s *SubPod) GetSound() *Sound {
  if s.Sound != nil {
    return s.Sound
  }

  return nil
}

// Returns the chardata from <minput> as string.
func (s *SubPod) GetMinput() string {
  return s.Minput.Minput
}

// Returns the chardata from <moutput> as string.
func (s *SubPod) GetMoutput() string {
  return s.Moutput.Moutput
}

// Returns the pointer to `Cell`.
func (s *SubPod) GetCell() *Cell {
  if s.Cell != nil {
    return s.Cell
  }

  return nil
}

// Returns the pointer to `States`.
func (s *SubPod) GetStates() *States {
  if s.States != nil {
    return s.States
  }

  return nil
}

// Returns whether the <cell> element exists or not.
func (s *SubPod) HasCell() bool {
  if s.Cell != nil {
    return true
  }

  return false
}

// Returns the chardata from the <cell> as string.
func (s *SubPod) GetCellExpression() string {
  if s.Cell != nil {
    return s.Cell.GetValue()
  }

  return ""
}

// Returns the "compressed" attribute value from <cell> as boolean.
func (c *Cell) IsCompressed() bool {
  b, _ := strconv.ParseBool(c.Compressed)

  return b
}

// Returns the chardata from the <cell> as string.
func (c *Cell) GetValue() string {
  return c.CellExpr
}
