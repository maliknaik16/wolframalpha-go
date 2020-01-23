
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

type ISounds interface {
	GetCount()						int
	GetSounds()						[]*Sound
	GetSound(int)					*Sound
}

type ISound interface {
	GetUrl()							string
	GetType()							string
}


func (s *Sounds) GetCount() int {
	c, _ := strconv.Atoi(s.Count)

	return c
}

func (s *Sounds) GetSounds() []*Sound {
	return s.Sounds
}

func (s *Sounds) GetSound(index int) *Sound {
	if index >= 0 && index < len(s.Sounds) - 1 {
		return s.Sounds[index]
	}

	return nil
}

func (s *Sound) GetUrl() string {
	return s.Url
}

func (s *Sound) GetType() string {
	return s.Type
}

