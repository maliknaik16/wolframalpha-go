
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <sources> tag.
type Sources struct {
	XMLName 							xml.Name 							`xml:"sources,omitempty"`
	Count 								string 								`xml:"count,attr"`
	Sources 							[]*Source 						`xml:"source,omitempty"`
}

// Struct for the <source> tag.
type Source struct {
	XMLName 							xml.Name 							`xml:"source,omitempty"`
	Url 									string 								`xml:"url,attr"`
	Text 									string 								`xml:"text,attr"`
}

type ISources interface {
	GetCount() 						int
	GetSources() 					[]*Source
	GetSource(int)				*Source
}

type ISource interface {
	GetUrl() 							string
	GetText() 						string
}


func (sources *Sources) GetCount() int {
	c, _ := strconv.Atoi(sources.Count)
}

func (sources *Sources) GetSources() []*Source {
	return sources.Sources
}

func (sources *Sources) GetSource(int) *Source {
  if index >= 0 && index < len(sources.Sources) - 1 {
    return sources.Sources[index]
  }

  return nil
}

func (source *Source) GetUrl() string {
	return source.Url
}

func (source *Source) GetText() string {
	return source.Text
}

