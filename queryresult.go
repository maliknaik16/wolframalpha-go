
package wolfram

import(
  "encoding/xml"
)

// Struct for the <queryresult> tag.
type QueryResult struct {
  XMLName             xml.Name                `xml:"queryresult,omitempty"`
  LanguageMsg         *LanguageMsg            `xml:"languagemsg,omitempty"`
  Tips                *Tips                   `xml:"tips,omitempty"`
  FutureTopic         *FutureTopic            `xml:"futuretopic,omitempty"`
  Success             string                  `xml:"success,attr"`
  ErrorAttr           string                  `xml:"error,attr"`
  Numpods             string                  `xml:"numpods,attr"`
  Version             string                  `xml:"version,attr"`
  DataTypes           string                  `xml:"datatypes,attr"`
  Timing              string                  `xml:"timing,attr"`
  TimedOut            string                  `xml:"timedout,attr"`
  ParseTiming         string                  `xml:"parsetiming,attr"`
  ParseTimedOut       string                  `xml:"parsetimedout,attr"`
  Recalculate         string                  `xml:"recalculate,attr"`
  TimedOutPods        string                  `xml:"timedoutpods,attr"`
  ID                  string                  `xml:"id,attr"`
  Host                string                  `xml:"host,attr"`
  Server              string                  `xml:"server,attr"`
  Related             string                  `xml:"related,attr"`
  Pods                []*Pod                  `xml:"pod,omitempty"`
  Assumptions         *Assumptions            `xml:"assumptions,omitempty"`
  States              *States                 `xml:"states,omitempty"`
  ExamplePage         *ExamplePage            `xml:"examplepage,omitempty"`
  Warnings            *Warnings               `xml:"warnings,omitempty"`
  Sources             *Sources                `xml:"sources,omitempty"`
  Generalizaton       *Generalizaton          `xml:"generalization,omitempty"`
  DidYouMeans         *DidYouMeans            `xml:"didyoumeans,omitempty"`
  ErrorTag            *ErrorTag               `xml:"error,omitempty"`
  Sounds              *Sounds                 `xml:"sounds,omitempty"`
}

// Struct for the <generalization> tag.
type Generalizaton struct {
  XMLName             xml.Name                `xml:"generalization,omitempty"`
  Topic               string                  `xml:"topic,attr"`
  Desc                string                  `xml:"desc,attr"`
  Url                 string                  `xml:"url,attr"`
}

// Struct for the <didyoumeans> tag.
type DidYouMeans struct {
  XMLName             xml.Name                `xml:"didyoumeans,omitempty"`
  Count               string                  `xml:"count,attr"`
  DidYouMean          []*DidYouMean           `xml:"didyoumean,omitempty"`
}

// Struct for the <didyoumean> tag.
type DidYouMean struct {
  XMLName             xml.Name                `xml:"didyoumean,omitempty"`
  Score               string                  `xml:"score,attr"`
  Level               string                  `xml:"level,attr"`
  Data                string                  `xml:",chardata"`
}

// Struct for the <error> tag.
type ErrorTag struct {
  XMLName             xml.Name                `xml:"error,omitempty"`
  Code                *Code                   `xml:"code,omitempty"`
  Msg                 *Msg                    `xml:"msg,omitempty"`
}

// Struct for the <code> tag.
type Code struct {
  XMLName             xml.Name                `xml:"code,omitempty"`
  Code                string                  `xml:",chardata"`
}

// Struct for the <msg> tag.
type Msg struct {
  XMLName             xml.Name                `xml:"msg,omitempty"`
  Msg                 string                  `xml:",chardata"`
}

// Struct for the <languagemsg> tag.
type LanguageMsg struct {
  XMLName             xml.Name                `xml:"languagemsg,omitempty"`
  English             string                  `xml:"english,attr"`
  Other               string                  `xml:"other,attr"`
}

// Struct for the <tips> tag.
type Tips struct {
  XMLName             xml.Name                `xml:"tips,omitempty"`
  Count               string                  `xml:"count,attr"`
  Tip                 []*Tip                  `xml:"tip,omitempty"`
}

// Struct for the <tip> tag.
type Tip struct {
  XMLName             xml.Name                `xml:"tip,omitempty"`
  AttrText            string                  `xml:"text,attr"`
}
