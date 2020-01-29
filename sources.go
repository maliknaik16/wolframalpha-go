
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <sources> tag.
type Sources struct {
  XMLName               xml.Name              `xml:"sources,omitempty"`
  Count                 string                `xml:"count,attr"`
  Sources               []*Source             `xml:"source,omitempty"`
}

// Struct for the <source> tag.
type Source struct {
  XMLName               xml.Name              `xml:"source,omitempty"`
  Url                   string                `xml:"url,attr"`
  Text                  string                `xml:"text,attr"`
}

// The interface for the `Sources`.
type ISources interface {
  GetCount()            int
  GetSources()          []*Source
  GetSource(int)        *Source
}

// The interface for the `Source`.
type ISource interface {
  GetUrl()              string
  GetText()             string
}

// Returns the number of <source> elements.
func (sources *Sources) GetCount() int {
  c, _ := strconv.Atoi(sources.Count)

  return c
}

// Returns the slice of pointers to `Source`.
func (sources *Sources) GetSources() []*Source {
  return sources.Sources
}

// Returns the pointer to `Source` at the given index.
func (sources *Sources) GetSource(index int) *Source {
  if index >= 0 && index < len(sources.Sources) - 1 {
    return sources.Sources[index]
  }

  return nil
}

// Returns the "url" attribute value from the <source>.
func (source *Source) GetUrl() string {
  return source.Url
}

// Returns the "text" attribute value from the <source>.
func (source *Source) GetText() string {
  return source.Text
}

