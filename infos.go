
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <infos> tag.
type Infos struct {
	XMLName 							xml.Name 							`xml:"infos,omitempty"`
	Count 								string 								`xml:"count,attr"`
	Infos 								[]*Info 							`xml:"info,omitempty"`
}

// Struct for the <info> tag.
type Info struct {
	XMLName 							xml.Name 							`xml:"info,omitempty"`
	AttrText 							string 								`xml:"text,attr"`
	Images 								[]*Image 							`xml:"img,omitempty"`
	Units 								*Units 								`xml:"units,omitempty"`
	Links 								[]*Link 							`xml:"link,omitempty"`
}

// Struct for the <link> tag.
type Link struct {
	XMLName 							xml.Name							`xml:"link,omitempty"`
	URL     							string 								`xml:"url,attr"`
	AttrText 							string 								`xml:"text,attr"`
	Title    							string 								`xml:"title,attr"`
}

// Struct for the <units> tag.
type Units struct {
	XMLName 							xml.Name							`xml:"units,omitempty"`
	Count 								string 								`xml:"count,attr"`
	Units 								[]*Unit 							`xml:"unit,omitempty"`
	Image 								*Image 								`xml:"img,omitempty"`
}

// Struct for the <unit> tag.
type Unit struct {
	XMLName 							xml.Name 							`xml:"unit,omitempty"`
	Short 								string 								`xml:"short,attr"`
	Long  								string 								`xml:"long,attr"`
}

type InfosInterface interface {
	GetCount()						int
	GetInfos()						[]*Info
	GetInfo(int)					*Info
}

type InfoInterface interface {
	GetTextAttr() 				string
	GetImages() 					[]*Image
	GetImage(int)					*Image
	GetUnits() 						*Units
	GetLinks() 						[]*Link
	GetLink(int)					*Link
}

type Linker interface {
	GetUrl()     					string
	GetTextAttr() 				string
	GetTitle()    				string
}

type IUnits interface {
	GetCount() 						int
	GetUnits() 						[]*Unit
	GetUnit(int)					*Unit
	GetImage() 						*Image
}

type IUnit interface {
	GetShort()						string
	GetLong()							string
}


func (infos *Infos) GetCount() int {
	c, _ := strconv.Atoi(infos.Count)

	return c
}

func (infos *Infos) GetInfos() []*Info {
	return infos.Infos
}

func (infos *Infos) GetInfo(index int) *Info {
	if index >= 0 && index < len(infos.Infos) - 1 {
		return infos.Infos[index]
	}

	return nil
}


func (info *Info) GetTextAttr() string {
	return info.AttrText
}

func (info *Info) GetImages() []*Image {
	return info.Images
}

func (info *Info) GetImage(index int) *Image {
	if index >= 0 && index < len(info.Images) - 1 {
		return info.Images[index]
	}

	return nil
}

func (info *Info) GetUnits() *Units {
	return info.Units
}

func (info *Info) GetLinks() []*Link {
	return info.Links
}

func (info *Info) GetLink(index int) *Link {
	if index >= 0 && index < len(info.Links) - 1 {
		return info.Links[index]
	}

	return nil
}


func (link *Link) GetUrl() string {
	return link.URL
}

func (link *Link) GetTextAttr() string {
	return link.AttrText
}

func (link *Link) GetTitle() string {
	return link.Title
}


func (units *Units) GetCount() int {
	c, _ := strconv.Atoi(units.Count)
	return c
}

func (units *Units) GetUnits() []*Unit {
	return units.Units
}

func (units *Units) GetUnit(index int) *Unit {
	if index >= 0 && index < len(units.Units) - 1 {
		return units.Units[index]
	}

	return nil
}

func (units *Units) GetImage() *Image {
	return units.Image
}


func (unit *Unit) GetShort() string {
	return unit.Short
}

func (unit *Unit) GetLong() string {
	return unit.Long
}

