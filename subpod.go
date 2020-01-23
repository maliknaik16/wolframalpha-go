
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <subpod> tag.
type SubPod struct {
	XMLName 							xml.Name 							`xml:"subpod,omitempty"`
	Title 								string 								`xml:"title,attr"`
	Image 								*Image 								`xml:"img,omitempty"`
	ImageMap 							*ImageMap 						`xml:"imagemap,omitempty"`
	PlainText 						*PlainText 						`xml:"plaintext,omitempty"`
	Sound 								*Sound 								`xml:"sound,omitempty"`
	Minput 								*Minput 							`xml:"minput,omitempty"`
	Moutput 							*Moutput 							`xml:"moutput,omitempty"`
	Cell 									*Cell 								`xml:"cell,omitempty"`
	States 								*States 							`xml:"states,omitempty"`
}

// Struct for the <plaintext> tag.
type PlainText struct {
  XMLName 							xml.Name 							`xml:"plaintext,omitempty"`
  PlainText 						string 								`xml:",chardata"`
}

// Struct for the <minput> tag.
type Minput struct {
  XMLName 							xml.Name 							`xml:"minput,omitempty"`
  Minput 								string 								`xml:",chardata"`
}

// Struct for the <moutput> tag.
type Moutput struct {
  XMLName 							xml.Name 							`xml:"moutput,omitempty"`
  Moutput 							string 								`xml:",chardata"`
}

// Struct for the <cell> tag.
type Cell struct {
	XMLName 							xml.Name 							`xml:"cell,omitempty"`
	Compressed 						string 								`xml:"compressed,attr"`
  CellExpr 							string 								`xml:",chardata"`
}

type ISubPod interface {
	GetTitle() 						string
	GetImage() 						*Image
	GetImageMap() 				*ImageMap
	GetPlainText() 				string
	GetSound() 						*Sound
	GetMinput() 					string
	GetMoutput() 					string
	GetCell() 						*Cell
	GetStates() 					*States
	HasCell()							bool
	GetCellExpression()		string
}

type ICell interface {
	IsCompressed()				bool
	GetValue()						string
}

func (s *SubPod) GetTitle() string {
	return s.Title
}

func (s *SubPod) GetImage() *Image {
	if s.Image != nil {
		return s.Image
	}

	return nil
}

func (s *SubPod) GetImageMap() *ImageMap {
	if s.ImageMap != nil {
		return s.ImageMap
	}

	return nil
}

func (s *SubPod) GetPlainText() string {
	return s.PlainText.PlainText
}

func (s *SubPod) GetSound() *Sound {
	if s.Sound != nil {
		return s.Sound
	}

	return nil
}

func (s *SubPod) GetMinput() string {
	return s.Minput.Minput
}

func (s *SubPod) GetMoutput() string {
	return s.Moutput.Moutput
}

func (s *SubPod) GetCell() *Cell {
	if s.Cell != nil {
		return s.Cell
	}

	return nil
}

func (s *SubPod) GetStates() *States {
	if s.States != nil {
		return s.States
	}

	return nil
}

func (s *SubPod) HasCell() bool {
	if s.Cell != nil {
		return true
	}

	return false
}

func (s *SubPod) GetCellExpression() string {
	if s.Cell != nil {
		return s.Cell.GetValue()
	}

	return ""
}

func (c *Cell) IsCompressed() bool {
	b, _ := strconv.ParseBool(c.Compressed)

	return b
}

func (c *Cell) GetValue() string {
	return c.CellExpr
}
