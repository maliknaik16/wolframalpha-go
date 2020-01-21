
package wolfram

import (
	"encoding/xml"
	"strconv"
)

// Struct for the <warnings> tag.
type Warnings struct {
	XMLName 							xml.Name 							`xml:"warnings,omitempty"`
	Count 								string 								`xml:"count,attr"`
	Reinterpret 					*Reinterpret 					`xml:"reinterpret,omitempty"`
	SpellCheck 						*SpellCheck 					`xml:"spellcheck,omitempty"`
	Delimiters 						*Delimiters 					`xml:"delimiters,omitempty"`
	Translation 					*Translation 					`xml:"translation,omitempty"`
}

// Struct for the <reinterpret> tag.
type Reinterpret struct {
	XMLName 							xml.Name 							`xml:"reinterpret,omitempty"`
	Text 									string 								`xml:"text,attr"`
	New 									string 								`xml:"new,attr"`
	Score 								string 								`xml:"score,attr"`
	Level 								string 								`xml:"level,attr"`
	Alternatives 					[]*Alternative 				`xml:"alternative,omitempty"`
}

// Struct for the <alternative> tag.
type Alternative struct {
	XMLName 							xml.Name 							`xml:"alternative,omitempty"`
	Score 								string 								`xml:"score,attr"`
	Level 								string 								`xml:"level,attr"`
	Value 								string 								`xml:",chardata"`
}

// Struct for the <spellcheck> tag.
type SpellCheck struct {
	XMLName 							xml.Name 							`xml:"spellcheck,omitempty"`
	Word 									string 								`xml:"word,attr"`
	Suggestion 						string 								`xml:"suggestion,attr"`
	Text 									string 								`xml:"text,attr"`
}

// Struct for the <delimiters> tag.
type Delimiters struct {
	XMLName 							xml.Name 							`xml:"delimiters,omitempty"`
	Text 									string 								`xml:"text,attr"`
}

// Struct for the <translation> tag.
type Translation struct {
	XMLName 							xml.Name 							`xml:"translation,omitempty"`
	Phrase 								string 								`xml:"phrase,attr"`
	Trans 								string 								`xml:"trans,attr"`
	Lang 									string 								`xml:"lang,attr"`
	Text 									string 								`xml:"text,attr"`
}

type IWarnings interface {
	GetCount()						int
	GetReinterpret()			*Reinterpret
	GetSpellCheck() 			*SpellCheck
	GetDelimiters() 			*Delimiters
	GetTranslation() 			*Translation
}

func (w *Warnings) GetCount() int {
	c, _ := strconv.Atoi(w.Count)

	return c
}

func (w *Warnings) GetReinterpret() *Reinterpret {
	return w.Reinterpret
}

func (w *Warnings) GetSpellCheck() *SpellCheck {
	return w.SpellCheck
}

func (w *Warnings) GetDelimiters() *Delimiters {
	return w.Delimiters
}

func (w *Warnings) GetTranslation() *Translation {
	return w.Translation
}

type IReinterpret interface {
	GetText() 							string
	GetNew() 								string
	GetScore() 							string
	GetLevel() 							string
	GetAlternatives() 			[]*Alternative
	GetAlternative(int) 		*Alternative
}


func (r *Reinterpret) GetText() string {
	return r.Text
}

func (r *Reinterpret) GetNew() string {
	return r.New
}

func (r *Reinterpret) GetScore() string {
	return r.Score
}

func (r *Reinterpret) GetLevel() string {
	return r.Level
}

func (r *Reinterpret) GetAlternatives() []*Alternative {
	return r.Alternatives
}

func (r *Reinterpret) GetAlternative(index int) *Alternative {
	if index >= 0 && index < len(r.Alternatives) {
		return r.Alternatives[index]
	}
	return nil
}

type IAlternative interface {
	GetScore() 							string
	GetLevel() 							string
	GetValue() 							string
}

func (a *Alternative) GetScore() string {
	return a.Score
}

func (a *Alternative) GetLevel() string {
	return a.Level
}

func (a *Alternative) GetValue() string {
	return a.Value
}

type ISpellCheck interface {
	GetWord() 							string
	GetSuggestion() 				string
	GetText() 							string
}

func (s *SpellCheck) GetWord() string {
	return s.Word
}

func (s *SpellCheck) GetSuggestion() string {
	return s.Suggestion
}

func (s *SpellCheck) GetText() string {
	return s.Text
}

type IDelimiters interface {
	GetText()								string
}

func (d *Delimiters) GetText() string {
	return d.Text
}

type ITranslation interface {
	GetPhrase() 						string
	GetTranslation() 				string
	GetLanguage() 					string
	GetText() 							string
}

func (t *Translation) GetPhrase() string {
	return t.Phrase
}

func (t *Translation) GetTranslation() string {
	return t.Trans
}

func (t *Translation) GetLanguage() string {
	return t.Lang
}

func (t *Translation) GetText() string {
	return t.Text
}
