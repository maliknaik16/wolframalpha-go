
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <states> tag.
type States struct {
  XMLName               xml.Name              `xml:"states,omitempty"`
  Count                 string                `xml:"count,attr"`
  StateLists            []*StateList          `xml:"statelist,omitempty"`
  States                []*State              `xml:"state,omitempty"`
}

// Struct for the <state> tag.
type State struct {
  XMLName               xml.Name              `xml:"state,omitempty"`
  Name                  string                `xml:"name,attr"`
  Input                 string                `xml:"input,attr"`
}

// Struct for the <statelist> tag.
type StateList struct {
  XMLName               xml.Name              `xml:"statelist,omitempty"`
  Count                 string                `xml:"count,attr"`
  Value                 string                `xml:"value,attr"`
  Delimiters            string                `xml:"delimiters,attr"`
  States                []*State              `xml:"state,omitempty"`
}

// The interface for the `States`.
type IStates interface {
  GetCount()            int
  GetStateLists()       []*StateList
  GetStateList(int)     *StateList
  GetStates()           []*State
  GetState(int)         *State
}

// The interface for the `State`.
type IState interface {
  GetName()             string
  GetInput()            string
}

// The interface for the `StateList`.
type IStateList interface {
  GetCount()            int
  GetValue()            string
  GetDelimiters()       string
  GetStates()           []*State
  GetState(int)         *State
}

// Returns the number of <state> elements.
func (states *States) GetCount() int {
  c, _ := strconv.Atoi(states.Count)

  return c
}

// Returns the slice of pointers to `StateList`.
func (states *States) GetStateLists() []*StateList {
  return states.StateLists
}

// Returns the pointer to `StateList` at the given index.
func (states *States) GetStateList(index int) *StateList {
  if index >= 0 && index < len(states.StateLists) - 1 {
    return states.StateLists[index]
  }

  return nil
}

// Returns the slice of pointers to `State`.
func (states *States) GetStates() []*State {
  return states.States
}

// Returns the pointer to `State` at the given index.
func (states *States) GetState(index int) *State {
  if index >= 0 && index < len(states.States) - 1 {
    return states.States[index]
  }

  return nil
}

// Returns the "name" attribute value from the <state>.
func (state *State) GetName() string {
  return state.Name
}

// Returns the "input" attribute value from the <state>.
func (state *State) GetInput() string {
  return state.Input
}

// Returns the number of <state> elements in <statelist>.
func (stateList *StateList) GetCount() int {
  c, _ := strconv.Atoi(stateList.Count)

  return c
}

// Returns the "value" attribute value from the <statelist>.
func (stateList *StateList) GetValue() string {
  return stateList.Value
}

// Returns the "delimiters" attribute value from the <statelist>.
func (stateList *StateList) GetDelimiters() string {
  return stateList.Delimiters
}

// Returns the slice of pointers to `State`.
func (stateList *StateList) GetStates() []*State {
  return stateList.States
}

// Returns the pointer to `State` at the given index.
func (stateList *StateList) GetState(index int) *State {
  if index >= 0 && index < len(stateList.States) - 1 {
    return stateList.States[index]
  }

  return nil
}

