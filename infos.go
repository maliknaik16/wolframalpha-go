
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <infos> tag.
type Infos struct {
  XMLName               xml.Name              `xml:"infos,omitempty"`
  Count                 string                `xml:"count,attr"`
  Infos                 []*Info               `xml:"info,omitempty"`
}

// Struct for the <info> tag.
type Info struct {
  XMLName               xml.Name              `xml:"info,omitempty"`
  AttrText              string                `xml:"text,attr"`
  Images                []*Image              `xml:"img,omitempty"`
  Units                 *Units                `xml:"units,omitempty"`
  Links                 []*Link               `xml:"link,omitempty"`
}

// Struct for the <link> tag.
type Link struct {
  XMLName               xml.Name              `xml:"link,omitempty"`
  URL                   string                `xml:"url,attr"`
  AttrText              string                `xml:"text,attr"`
  Title                 string                `xml:"title,attr"`
}

// Struct for the <units> tag.
type Units struct {
  XMLName               xml.Name              `xml:"units,omitempty"`
  Count                 string                `xml:"count,attr"`
  Units                 []*Unit               `xml:"unit,omitempty"`
  Image                 *Image                `xml:"img,omitempty"`
}

// Struct for the <unit> tag.
type Unit struct {
  XMLName               xml.Name              `xml:"unit,omitempty"`
  Short                 string                `xml:"short,attr"`
  Long                  string                `xml:"long,attr"`
}

// The interface for the Infos.
type InfosInterface interface {
  GetCount()            int
  GetInfos()            []*Info
  GetInfo(int)          *Info
}

// The interface for the Info.
type InfoInterface interface {
  GetTextAttr()         string
  GetImages()           []*Image
  GetImage(int)         *Image
  GetUnits()            *Units
  GetLinks()            []*Link
  GetLink(int)          *Link
}

// The interface for the Link.
type Linker interface {
  GetUrl()              string
  GetTextAttr()         string
  GetTitle()            string
}

// The interface for the Units.
type IUnits interface {
  GetCount()            int
  GetUnits()            []*Unit
  GetUnit(int)          *Unit
  GetImage()            *Image
}

// The interface for the Unit.
type IUnit interface {
  GetShort()            string
  GetLong()             string
}

// Returns the number of <info> elements.
func (infos *Infos) GetCount() int {
  c, _ := strconv.Atoi(infos.Count)

  return c
}

// Returns the slice of pointers to `Info`.
func (infos *Infos) GetInfos() []*Info {
  return infos.Infos
}

// Returns the pointer to `Info` at the given index.
func (infos *Infos) GetInfo(index int) *Info {
  if index >= 0 && index < len(infos.Infos) - 1 {
    return infos.Infos[index]
  }

  return nil
}

// Returns the "text" attribute value from the <info>.
func (info *Info) GetTextAttr() string {
  return info.AttrText
}

// Returns the slice of pointers to `Image`.
func (info *Info) GetImages() []*Image {
  return info.Images
}

// Returns the pointer to `Image` at the given index.
func (info *Info) GetImage(index int) *Image {
  if index >= 0 && index < len(info.Images) - 1 {
    return info.Images[index]
  }

  return nil
}

// Returns the pointer to `Units`.
func (info *Info) GetUnits() *Units {
  return info.Units
}

// Returns the slice of pointers to `Link`.
func (info *Info) GetLinks() []*Link {
  return info.Links
}

// Returns the pointer to `Link` at the given index.
func (info *Info) GetLink(index int) *Link {
  if index >= 0 && index < len(info.Links) - 1 {
    return info.Links[index]
  }

  return nil
}

// Returns the "url" attribute value from the <link>.
func (link *Link) GetUrl() string {
  return link.URL
}

// Returns the "text" attribute value from the <link>.
func (link *Link) GetTextAttr() string {
  return link.AttrText
}

// Returns the "title" attribute value from the <link>.
func (link *Link) GetTitle() string {
  return link.Title
}

// Returns the number of <unit> elements.
func (units *Units) GetCount() int {
  c, _ := strconv.Atoi(units.Count)
  return c
}

// Returns the slice of pointers to the `Unit`.
func (units *Units) GetUnits() []*Unit {
  return units.Units
}

// Returns the pointer to `Unit` at the given index.
func (units *Units) GetUnit(index int) *Unit {
  if index >= 0 && index < len(units.Units) - 1 {
    return units.Units[index]
  }

  return nil
}

// Returns the pointer to the `Image`.
func (units *Units) GetImage() *Image {
  return units.Image
}

// Returns the "short" attribute value from the <unit>.
func (unit *Unit) GetShort() string {
  return unit.Short
}

// Returns the "long" attribute value from the <unit>.
func (unit *Unit) GetLong() string {
  return unit.Long
}

