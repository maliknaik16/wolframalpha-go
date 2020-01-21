
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <didyoumeans> tag.
type DidYouMeans struct {
  XMLName             xml.Name                `xml:"didyoumeans,omitempty"`
  Count               string                  `xml:"count,attr"`
  DidYouMean          []*DidYouMean           `xml:"didyoumean,omitempty"`
}

// The interface for the DidYouMeans
type IDidYouMeans interface {
	GetCount()					int
	GetDidYouMeans()		[]*DidYouMean
	GetDidYouMean()			*DidYouMean
}

// Returns the number of <didyoumean> elements.
func (dyms *DidYouMeans) GetCount() int {
	c, _ := strconv.Atoi(dyms.Count)

	return c
}

// Returns the slice of <didyoumean> elements.
func (dyms *DidYouMeans) GetDidYouMeans() []*DidYouMean {
	return dyms.DidYouMean
}

// Returns the <didyoumean> at the given index.
func (dyms *DidYouMeans) GetDidYouMean(index int) *DidYouMean {
	if index >= 0 && index < len(dyms.DidYouMean) - 1 {
		return dyms.DidYouMean[index]
	}

	return nil
}
