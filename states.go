
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <states> tag.
type States struct {
	XMLName 							xml.Name 							`xml:"states,omitempty"`
	Count 								string 								`xml:"count,attr"`
	StateLists 						[]*StateList 					`xml:"statelist,omitempty"`
	States 								[]*State 							`xml:"state,omitempty"`
}

// Struct for the <state> tag.
type State struct {
	XMLName 							xml.Name 							`xml:"state,omitempty"`
	Name  								string 								`xml:"name,attr"`
	Input 								string 								`xml:"input,attr"`
}

// Struct for the <statelist> tag.
type StateList struct {
	XMLName 							xml.Name 							`xml:"statelist,omitempty"`
	Count  								string 								`xml:"count,attr"`
	Value 								string 								`xml:"value,attr"`
	Delimiters 						string 								`xml:"delimiters,attr"`
	States 								[]*State 							`xml:"state,omitempty"`
}

// Struct for the <states> tag.
type IStates interface {
	GetCount() 						int
	GetStateLists() 			[]*StateList
	GetStateList(int) 		*StateList
	GetStates() 					[]*State
	GetState(int) 				*State
}

// Struct for the <state> tag.
type IState interface {
	GetName()  						string
	GetInput() 						string
}

// Struct for the <statelist> tag.
type IStateList interface {
	GetCount()  					int
	GetValue() 						string
	GetDelimiters() 			string
	GetStates() 					[]*State
	GetState(int) 				*State
}


func (states *States) GetCount() int {
	c, _ := strconv.Atoi(states.Count)

	return c
}

func (states *States) GetStateLists() []*StateList {
	return states.StateLists
}

func (states *States) GetStateList(index int) *StateList {
  if index >= 0 && index < len(states.StateLists) - 1 {
    return states.StateLists[index]
  }

  return nil
}

func (states *States) GetStates() []*State {
	return states.States
}

func (states *States) GetState(index int) *State {
  if index >= 0 && index < len(states.States) - 1 {
    return states.States[index]
  }

  return nil
}


func (state *State) GetName() string {
	return state.Name
}

func (state *State) GetInput() string {
	return state.Input
}

func (stateList *StateList) GetCount() int {
	c, _ := strconv.Atoi(statelist.Count)

	return c
}

func (stateList *StateList) GetValue() string {
	return stateList.Value
}

func (stateList *StateList) GetDelimiters() string {
	return stateList.Delimiters
}

func (stateList *StateList) GetStates() []*State {
	return stateList.States
}

func (stateList *StateList) GetState(index int) *State {
  if index >= 0 && index < len(stateList.States) - 1 {
    return stateList.States[index]
  }

  return nil
}

