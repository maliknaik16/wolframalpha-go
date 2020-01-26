
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

// The interface for the `ImageMap`.
type ImageMaper interface {
	GetRects()						[]*Rect
	GetRect(int)					*Rect
}

// The interface for the `Rect`.
type IRect interface {
	GetLeft()							int
	GetTop() 							int
	GetRight() 						int
	GetBottom() 					int
	GetQuery() 						string
	GetAssumptions() 			string
	GetTitle() 						string
}

// Returns the slice of pointers to `Rect`.
func (imagemap *ImageMap) GetRects() []*Rect {
	return imagemap.Rects
}

// Returns the pointer to the `Rect` at the given index.
func (imagemap *ImageMap) GetRect(index int) *Rect {
	if index >= 0 && index < len(imagemap.Rects) - 1 {
		return imagemap.Rects[index]
	}

	return nil
}

// Returns the "left" attribute value from the <rect>.
func (rect *Rect) GetLeft() int {
	l, _ := strconv.Atoi(rect.Left)

	return l
}

// Returns the "top" attribute value from the <rect>.
func (rect *Rect) GetTop() int {
	t, _ := strconv.Atoi(rect.Top)

	return t
}

// Returns the "right" attribute value from the <rect>.
func (rect *Rect) GetRight() int {
	rt, _ := strconv.Atoi(rect.Right)

	return rt
}

// Returns the "bottom" attribute value from the <rect>.
func (rect *Rect) GetBottom() int {
	b, _ := strconv.Atoi(rect.Bottom)

	return b
}

// Returns the "query" attribute value from the <rect>.
func (rect *Rect) GetQuery() string {
	return rect.Query
}

// Returns the "assumptions" attribute value from the <rect>.
func (rect *Rect) GetAssumptions() string {
	return rect.Assumptions
}

// Returns the "title" attribute value from the <rect>.
func (rect *Rect) GetTitle() string {
	return rect.Title
}

