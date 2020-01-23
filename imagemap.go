
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <imagemap> tag.
type ImageMap struct {
	XMLName 							xml.Name 							`xml:"imagemap,omitempty"`
	Rects 								[]*Rect 							`xml:"rect,omitempty"`
}

// Struct for the <rect> tag.
type Rect struct {
	XMLName 							xml.Name 							`xml:"rect,omitempty"`
	Left 									string 								`xml:"left,attr"`
	Top 									string 								`xml:"top,attr"`
	Right 								string 								`xml:"right,attr"`
	Bottom 								string 								`xml:"bottom,attr"`
	Query 								string 								`xml:"query,attr"`
	Assumptions 					string 								`xml:"assumptions,attr"`
	Title 								string 								`xml:"title,attr"`
}

type ImageMaper interface {
	GetRects()						[]*Rect
	GetRect(int)					*Rect
}

type IRect interface {
	GetLeft()							int
	GetTop() 							int
	GetRight() 						int
	GetBottom() 					int
	GetQuery() 						string
	GetAssumptions() 			string
	GetTitle() 						string
}

func (i *ImageMap) GetRects() []*Rect {
	return i.Rects
}

func (i *ImageMap) GetRect(index int) *Rect {
	if index >= 0 && index < len(i.Rects) - 1 {
		return i.Rects[index]
	}

	return nil
}

// Rect


func (r *Rect) GetLeft() int {
	l, _ := strconv.Atoi(r.Left)

	return l
}

func (r *Rect) GetTop() int {
	t, _ := strconv.Atoi(r.Top)

	return t
}

func (r *Rect) GetRight() int {
	rt, _ := strconv.Atoi(r.Right)

	return rt
}

func (r *Rect) GetBottom() int {
	b, _ := strconv.Atoi(r.Bottom)

	return b
}

func (r *Rect) GetQuery() string {
	return r.Query
}

func (r *Rect) GetAssumptions() string {
	return r.Assumptions
}

func (r *Rect) GetTitle() string {
	return r.Title
}

