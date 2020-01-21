
package wolfram

import(
  "encoding/xml"
  "strconv"
)

// Struct for the <queryresult> tag.
type QueryResult struct {
  XMLName             xml.Name                `xml:"queryresult,omitempty"`
  LanguageMsg         *LanguageMsg            `xml:"languagemsg,omitempty"`
  Tips                *Tips                   `xml:"tips,omitempty"`
  FutureTopic         *FutureTopic            `xml:"futuretopic,omitempty"`
  Success             string                  `xml:"success,attr"`
  ErrorAttr           string                  `xml:"error,attr"`
  NumPods             string                  `xml:"numpods,attr"`
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
  Generalization      *Generalization          `xml:"generalization,omitempty"`
  DidYouMeans         *DidYouMeans            `xml:"didyoumeans,omitempty"`
  ErrorTag            *ErrorTag               `xml:"error,omitempty"`
  Sounds              *Sounds                 `xml:"sounds,omitempty"`
}

// The interface for the QueryResult.
type QueryResulter interface {
  // Elements
  GetPods()             []*Pod
  GetPod(int)           *Pod
  GetAssumptions()      *Assumptions
  GetStates()           *States
  GetExamplePage()      *ExamplePage
  GetWarnings()         *Warnings
  GetSources()          *Sources
  GetGeneralization()   *Generalization
  GetDidYouMeans()      *DidYouMeans
  GetError()            *ErrorTag
  GetSounds()           *Sounds
  GetLanguageMsg()      *LanguageMsg
  GetTips()             *Tips
  GetFutureTopic()      *FutureTopic

  // Attributes
  GetSuccess()          bool
  GetErrorAttr()        bool
  GetNumPods()          int
  GetVersion()          string
  GetDataTypes()        string
  GetTiming()           float64
  GetTimedOut()         string
  GetParseTiming()      float64
  GetParseTimedOut()    bool
  GetRecalculate()      string
  GetTimedOutPods()     string
  GetId()               string
  GetHost()             string
  GetServer()           string
  GetRelated()          string
}

// Returns the 'Pod' slice.
func (q *QueryResult) GetPods() []*Pod {
  if q.Pods != nil {
    return q.Pods
  }
  return nil
}

// Returns the 'Pod' at the given index.
func (q *QueryResult) GetPod(index int) *Pod {
  if index >= 0 && index > len(q.Pods) - 1 {
    return nil
  }
  return q.Pods[index]
}

// Returns the 'Assumptions' slice.
func (q *QueryResult) GetAssumptions() *Assumptions {
  if q.Assumptions != nil {
    return q.Assumptions
  }
  return nil
}

// Returns the 'States'.
func (q *QueryResult) GetStates() *States {
  if q.States != nil {
    return q.States
  }
  return nil
}

// Returns the 'ExamplePage'.
func (q *QueryResult) GetExamplePage() *ExamplePage {
  if q.ExamplePage != nil {
    return q.ExamplePage
  }
  return nil
}

// Returns the 'Warnings'.
func (q *QueryResult) GetWarnings() *Warnings {
  if q.Warnings != nil {
    return q.Warnings
  }
  return nil
}

// Returns the 'Sources'.
func (q *QueryResult) GetSources() *Sources {
  if q.Sources != nil {
    return q.Sources
  }
  return nil
}

// Returns the 'Generalization'.
func (q *QueryResult) GetGeneralization() *Generalization {
  if q.Generalization != nil {
    return q.Generalization
  }
  return nil
}

// Returns the 'DidYouMeans'.
func (q *QueryResult) GetDidYouMeans() *DidYouMeans {
  if q.DidYouMeans != nil {
    return q.DidYouMeans
  }
  return nil
}

// Returns the 'ErrorTag'.
func (q *QueryResult) GetError() *ErrorTag {
  if q.ErrorTag != nil {
    return q.ErrorTag
  }
  return nil
}

// Returns the 'Sounds'.
func (q *QueryResult) GetSounds() *Sounds {
  if q.Sounds != nil {
    return q.Sounds
  }
  return nil
}

// Returns the 'LanguageMsg'.
func (q *QueryResult) GetLanguageMsg() *LanguageMsg {
  if q.LanguageMsg != nil {
    return q.LanguageMsg
  }
  return nil
}

// Returns the 'FutureTopic'.
func (q *QueryResult) GetFutureTopic() *FutureTopic {
  if q.FutureTopic != nil {
    return q.FutureTopic
  }
  return nil
}

// Returns the 'Tips'.
func (q *QueryResult) GetTips() *Tips {
  if q.Tips != nil {
    return q.Tips
  }
  return nil
}

// Returns the attribute value of 'success' from the <queryresult>.
func (q *QueryResult) GetSuccess() bool {
  s, _ := strconv.ParseBool(q.Success)

  return s
}

// Returns the attribute value of 'error' from the <queryresult>.
func (q *QueryResult) GetErrorAttr() bool {
  ea, _ := strconv.ParseBool(q.ErrorAttr)

  return ea
}

// Returns the attribute value of 'numpods' from the <queryresult>.
func (q *QueryResult) GetNumPods() int {
  np, _ := strconv.Atoi(q.NumPods)

  return np
}

// Returns the attribute value of 'version' from the <queryresult>.
func (q *QueryResult) GetVersion() string {
  return q.Version
}

// Returns the attribute value of 'datatypes' from the <queryresult>.
func (q *QueryResult) GetDataTypes() string {
  return q.DataTypes
}

// Returns the attribute value of 'timing' from the <queryresult>.
func (q *QueryResult) GetTiming() float64 {
  t, _ := strconv.ParseFloat(q.Timing, 64)

  return t
}

// Returns the attribute value of 'timedout' from the <queryresult>.
func (q *QueryResult) GetTimedOut() string {
  return q.TimedOut
}

// Returns the attribute value of 'parsetiming' from the <queryresult>.
func (q *QueryResult) GetParseTiming() float64 {
  pt, _ := strconv.ParseFloat(q.ParseTiming, 64)

  return pt
}

// Returns the attribute value of 'parsetimedout' from the <queryresult>.
func (q *QueryResult) GetParseTimedOut() bool {
  pto, _ := strconv.ParseBool(q.ParseTimedOut)

  return pto
}

// Returns the attribute value of 'recalculate' from the <queryresult>.
func (q *QueryResult) GetRecalculate() string {
  return q.Recalculate
}

// Returns the attribute value of 'timedoutpods' from the <queryresult>.
func (q *QueryResult) GetTimedOutPods() string {
  return q.TimedOutPods
}

// Returns the attribute value of 'id' from the <queryresult>.
func (q *QueryResult) GetId() string {
  return q.ID
}

// Returns the attribute value of 'host' from the <queryresult>.
func (q *QueryResult) GetHost() string {
  return q.Host
}

// Returns the attribute value of 'server' from the <queryresult>.
func (q *QueryResult) GetServer() string {
  return q.Server
}

// Returns the attribute value of 'related' from the <queryresult>.
func (q *QueryResult) GetRelated() string {
  return q.Related
}
