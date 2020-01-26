
package wolfram

import (
  "encoding/xml"
)

// Struct for the <pod> tag.
type Pod struct {
  XMLName               xml.Name              `xml:"pod,omitempty"`
  Title                 string                `xml:"title,attr"`
  ErrorAttr             string                `xml:"error,attr"`
  Position              string                `xml:"position,attr"`
  Scanner               string                `xml:"scanner,attr"`
  ID                    string                `xml:"id,attr"`
  NumSubPods            string                `xml:"numsubpods,attr"`
  Primary               string                `xml:"primary,attr"`
  SubPods               []*SubPod             `xml:"subpod,omitempty"`
  ErrorTag              *ErrorTag             `xml:"error,omitempty"`
  ExpressionTypes       *ExpressionTypes      `xml:"expressiontypes,omitempty"`
  States                *States               `xml:"states,omitempty"`
  Infos                 *Infos                `xml:"infos,omitempty"`
}

// Struct for the <expressiontypes> tag.
type ExpressionTypes struct {
  XMLName               xml.Name              `xml:"expressiontypes,omitempty"`
  Count                 string                `xml:"count,attr"`
  ExpressionTypes       []*ExpressionType     `xml:"expressiontype,omitempty"`
}

// Struct for the <expressiontype> tag.
type ExpressionType struct {
  XMLName               xml.Name              `xml:"expressiontype,omitempty"`
  Name                  string                `xml:"name,attr"`
}

// The interface for the `Pod`.
type IPod interface {
  GetTitle()            string
  GetErrorAttr()        string
  GetPosition()         string
  GetScanner()          string
  GetId()               string
  GetNumSubPods()       string
  GetPrimary()          string
  GetSubPods()          []*SubPod
  GetSubPod(int)        *SubPod
  GetErrorTag()         *ErrorTag
  GetExpressionTypes()  []*ExpressionType
  GetStates()           *States
  GetInfos()            *Infos
}

// The interface for the `ExpressionTypes`.
type IExpressionTypes interface {
  GetCount()             int
  GetExpressionTypes()   []*ExpressionType
  GetExpressionType(int) *ExpressionType
}

// The interface for the `ExpressionType`.
type IExpressionType interface {
  GetName()             string
}

// Returns the slice of pointers to `ExpressionType`.
func (pod *Pod) GetExpressionTypes() []*ExpressionType {
  return pod.ExpressionTypes.ExpressionTypes
}

// Returns the "title" attribute value from the <pod>.
func (pod *Pod) GetTitle() string {
  return pod.Title
}

// Returns the "error" attribute value from the <pod>.
func (pod *Pod) GetErrorAttr() string {
  return pod.ErrorAttr
}

// Returns the "position" attribute value from the <pod>.
func (pod *Pod) GetPosition() string {
  return pod.Position
}

// Returns the "scanner" attribute value from the <pod>.
func (pod *Pod) GetScanner() string {
  return pod.Scanner
}

// Returns the "id" attribute value from the <pod>.
func (pod *Pod) GetId() string {
  return pod.ID
}

// Returns the "numsubpods" attribute value from the <pod>.
func (pod *Pod) GetNumSubPods() string {
  return pod.NumSubPods
}

// Returns the "primary" attribute value from the <pod>.
func (pod *Pod) GetPrimary() string {
  return pod.Primary
}

// Returns the slice of pointers to `SubPod`.
func (pod *Pod) GetSubPods() []*SubPod {
  return pod.SubPods
}

// Returns the pointer to `SubPod` at the given index.
func (pod *Pod) GetSubPod(index  int) *SubPod {
  if index >= 0 && index < len(pod.SubPods) - 1 {
    return pod.SubPods[index]
  }

  return nil
}

// Returns the pointer to `ErrorTag`.
func (pod *Pod) GetErrorTag() *ErrorTag {
  return pod.ErrorTag
}

// Returns the slice of pointers to `ExpressionType`.
func (pod *Pod) GetExpressionTypes() []*ExpressionType {
  return pod.ExpressionTypes
}

// Returns the pointer to `States`.
func (pod *Pod) GetStates() *States {
  return pod.States
}
// Returns the pointer to `Infos`.
func (pod *Pod) GetInfos() *Infos {
  return pod.Infos
}

// Returns the number of <expressiontype> elements.
func (ets *ExpressionTypes) GetCount() int {
  c, _ := strconv.Atoi(ets.Count)

  return c
}

// Returns the slice of pointers to `ExpressionType`.
func (ets *ExpressionTypes) GetExpressionTypes() []*ExpressionType {
  return ets.ExpressionTypes
}

// Returns the pointer to `ExpressionType` at the given index.
func (ets *ExpressionTypes) GetExpressionType(index int) *ExpressionType {
  if index >= 0 && index < len(ets.ExpressionTypes) - 1 {
    return ets.ExpressionTypes[index]
  }

  return nil
}

// Returns the "name" attribute value from the <expressiontype>.
func (et *ExpressionType) GetName() string {
  return et.Name
}
