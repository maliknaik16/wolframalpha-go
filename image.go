
package wolfram

import (
  "encoding/xml"
  "strconv"
  "strings"
)

// Struct for the <img> tag.
type Image struct {
  XMLName               xml.Name              `xml:"img,omitempty"`
  Src                   string                `xml:"src,attr"`
  Alt                   string                `xml:"alt,attr"`
  Title                 string                `xml:"title,attr"`
  Width                 string                `xml:"width,attr"`
  Height                string                `xml:"height,attr"`
  Type                  string                `xml:"type,attr,omitempty"`
  Themes                string                `xml:"themes,attr,omitempty"`
  ColorInvertable       string                `xml:"colorinvertable,attr,omitempty"`
}

// The Image format constants
const (
  FORMAT_UNKNOWN int = 0
  FORMAT_GIF int = 1
  FORMAT_PNG int = 2
)

// The dimensions struct
type Dimension struct {
  Width  int
  Height int
}

// Returns the value of the "src" attribute from <img>.
func (i *Image) GetSrc() string {
  return i.Src
}

// Returns the value of the "alt" attribute from <img>.
func (i *Image) GetAlt() string {
  return i.Alt
}

// Returns the value of the "title" attribute from <img>.
func (i *Image) GetTitle() string {
  return i.Title
}

// Returns the dimensions of the image as `Dimension` struct.
func (i *Image) GetDimensions() Dimension {
  width, _ := strconv.Atoi(i.Width)
  height, _ := strconv.Atoi(i.Height)

  return Dimension{width, height}
}

// Returns the value of the "type" attribute from <img>.
func (i *Image) GetType() string {
  return i.Type
}

// Returns the value of the "themes" attribute from <img>.
func (i *Image) GetThemes() string {
  return i.Themes
}

// Returns the value of the "colorinvertable" attribute as boolean.
func (i *Image) IsColorInvertable() bool {
  ci, _ := strconv.ParseBool(i.ColorInvertable)

  return ci
}

// Returns the format of the image from the url.
//
// Returns the FORMAT_GIF integer constant if the format of the image is 'gif',
// FORMAT_PNG if the format of the image is 'png', and FORMAT_UNKNOWN for any
// other image formats.
func (i *Image) GetFormat() int {
  url := i.Src
  index := strings.LastIndex(url, "MSPStoreType=image/")

  if index > 0 {
    format := url[index + 19: index + 22]

    if format == "gif" {
      return FORMAT_GIF
    }

    if format == "png" {
      return FORMAT_PNG
    }
  }

  if strings.HasSuffix(url, "gif") {
    return FORMAT_GIF
  }

  if strings.HasSuffix(url, "png") {
    return FORMAT_PNG
  }

  return FORMAT_UNKNOWN
}
