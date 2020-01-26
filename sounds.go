
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <sounds> tag.
type Sounds struct {
	XMLName 							xml.Name 							`xml:"sounds,omitempty"`
	Count 								string 								`xml:"count,attr"`
	Sounds 								[]*Sound 							`xml:"sound,omitempty"`
}

// Struct for the <sound> tag.
type Sound struct {
	XMLName 							xml.Name 							`xml:"sound,omitempty"`
	Url 									string 								`xml:"url,attr"`
	Type 									string 								`xml:"type,attr"`
}

// The interface for the `Sounds`.
type ISounds interface {
	GetCount()						int
	GetSounds()						[]*Sound
	GetSound(int)					*Sound
}

// The interface for the `Sound`.
type ISound interface {
	GetUrl()							string
	GetType()							string
}

// Returns the number of <sound> elements.
func (s *Sounds) GetCount() int {
	c, _ := strconv.Atoi(s.Count)

	return c
}

// Returns the slice of pointers to `Sound`.
func (s *Sounds) GetSounds() []*Sound {
	return s.Sounds
}

// Returns the pointer to `Sound` at the given index.
func (s *Sounds) GetSound(index int) *Sound {
	if index >= 0 && index < len(s.Sounds) - 1 {
		return s.Sounds[index]
	}

	return nil
}

// Returns the "url" attribute value from the <rect>.
func (s *Sound) GetUrl() string {
	return s.Url
}

// Returns the "type" attribute value from the <rect>.
func (s *Sound) GetType() string {
	return s.Type
}

