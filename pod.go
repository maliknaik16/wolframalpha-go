
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

type IExpressionTypes interface {
  GetCount()             int
  GetExpressionTypes()   []*ExpressionType
  GetExpressionType(int) *ExpressionType
}

type IExpressionType interface {
  GetName()             string
}

func (pod *Pod) GetExpressionTypes() []*ExpressionType {
  return pod.ExpressionTypes.ExpressionTypes
}


func (pod *Pod) GetTitle() string {
  return pod.Title
}

func (pod *Pod) GetErrorAttr() string {
  return pod.ErrorAttr
}

func (pod *Pod) GetPosition() string {
  return pod.Position
}

func (pod *Pod) GetScanner() string {
  return pod.Scanner
}

func (pod *Pod) GetId() string {
  return pod.ID
}

func (pod *Pod) GetNumSubPods() string {
  return pod.NumSubPods
}

func (pod *Pod) GetPrimary() string {
  return pod.Primary
}

func (pod *Pod) GetSubPods() []*SubPod {
  return pod.SubPods
}

func (pod *Pod) GetSubPod(index  int) *SubPod {
  if index >= 0 && index < len(pod.SubPods) - 1 {
    return pod.SubPods[index]
  }

  return nil
}

func (pod *Pod) GetErrorTag() *ErrorTag {
  return pod.ErrorTag
}

func (pod *Pod) GetExpressionTypes() []*ExpressionType {
  return pod.ExpressionTypes
}

func (pod *Pod) GetStates() *States {
  return pod.States
}

func (pod *Pod) GetInfos() *Infos {
  return pod.Infos
}




func (ets *ExpressionTypes) GetCount() int {
  c, _ := strconv.Atoi(ets.Count)

  return c
}

func (ets *ExpressionTypes) GetExpressionTypes() []*ExpressionType {
  return ets.ExpressionTypes
}

func (ets *ExpressionTypes) GetExpressionType(index int) *ExpressionType {
  if index >= 0 && index < len(ets.ExpressionTypes) - 1 {
    return ets.ExpressionTypes[index]
  }

  return nil
}


func (et *ExpressionType) GetName() string {
  return et.Name
}
