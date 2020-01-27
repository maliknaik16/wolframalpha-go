
package wolfram

import (
  "encoding/xml"
  "strconv"
)

// Struct for the <warnings> tag.
type Warnings struct {
  XMLName               xml.Name              `xml:"warnings,omitempty"`
  Count                 string                `xml:"count,attr"`
  Reinterpret           *Reinterpret          `xml:"reinterpret,omitempty"`
  SpellCheck            *SpellCheck           `xml:"spellcheck,omitempty"`
  Delimiters            *Delimiters           `xml:"delimiters,omitempty"`
  Translation           *Translation          `xml:"translation,omitempty"`
}

// Struct for the <reinterpret> tag.
type Reinterpret struct {
  XMLName               xml.Name              `xml:"reinterpret,omitempty"`
  Text                  string                `xml:"text,attr"`
  New                   string                `xml:"new,attr"`
  Score                 string                `xml:"score,attr"`
  Level                 string                `xml:"level,attr"`
  Alternatives          []*Alternative        `xml:"alternative,omitempty"`
}

// Struct for the <alternative> tag.
type Alternative struct {
  XMLName               xml.Name              `xml:"alternative,omitempty"`
  Score                 string                `xml:"score,attr"`
  Level                 string                `xml:"level,attr"`
  Value                 string                `xml:",chardata"`
}

// Struct for the <spellcheck> tag.
type SpellCheck struct {
  XMLName               xml.Name              `xml:"spellcheck,omitempty"`
  Word                  string                `xml:"word,attr"`
  Suggestion            string                `xml:"suggestion,attr"`
  Text                  string                `xml:"text,attr"`
}

// Struct for the <delimiters> tag.
type Delimiters struct {
  XMLName               xml.Name              `xml:"delimiters,omitempty"`
  Text                  string                `xml:"text,attr"`
}

// Struct for the <translation> tag.
type Translation struct {
  XMLName               xml.Name              `xml:"translation,omitempty"`
  Phrase                string                `xml:"phrase,attr"`
  Trans                 string                `xml:"trans,attr"`
  Lang                  string                `xml:"lang,attr"`
  Text                  string                `xml:"text,attr"`
}

// The interface for the `Warnings`.
type IWarnings interface {
  GetCount()            int
  GetReinterpret()      *Reinterpret
  GetSpellCheck()       *SpellCheck
  GetDelimiters()       *Delimiters
  GetTranslation()      *Translation
}

// The interface for the `Reinterpret`.
type IReinterpret interface {
  GetText()               string
  GetNew()                string
  GetScore()              string
  GetLevel()              string
  GetAlternatives()       []*Alternative
  GetAlternative(int)     *Alternative
}

// The interface for the `Alternative`.
type IAlternative interface {
  GetScore()              string
  GetLevel()              string
  GetValue()              string
}

// The interface for the `SpellCheck`.
type ISpellCheck interface {
  GetWord()               string
  GetSuggestion()         string
  GetText()               string
}

// The interface for the `Delimiters`.
type IDelimiters interface {
  GetText()               string
}

// The interface for the `Translation`.
type ITranslation interface {
  GetPhrase()             string
  GetTranslation()        string
  GetLanguage()           string
  GetText()               string
}

// Returns the number of warnings.
func (w *Warnings) GetCount() int {
  c, _ := strconv.Atoi(w.Count)

  return c
}

// Returns the pointer to `Reinterpret`.
func (w *Warnings) GetReinterpret() *Reinterpret {
  return w.Reinterpret
}

// Returns the pointer to `SpellCheck`.
func (w *Warnings) GetSpellCheck() *SpellCheck {
  return w.SpellCheck
}

// Returns the pointer to `Delimiters`.
func (w *Warnings) GetDelimiters() *Delimiters {
  return w.Delimiters
}

// Returns the pointer to `Translation`.
func (w *Warnings) GetTranslation() *Translation {
  return w.Translation
}

// Returns the "text" attribute value from the <reinterpret>.
func (r *Reinterpret) GetText() string {
  return r.Text
}

// Returns the "new" attribute value from the <reinterpret>.
func (r *Reinterpret) GetNew() string {
  return r.New
}

// Returns the "score" attribute value from the <reinterpret>.
func (r *Reinterpret) GetScore() string {
  return r.Score
}

// Returns the "level" attribute value from the <reinterpret>.
func (r *Reinterpret) GetLevel() string {
  return r.Level
}

// Returns the slice of pointers to `Alternative`.
func (r *Reinterpret) GetAlternatives() []*Alternative {
  return r.Alternatives
}

// Returns the pointer to `Alternative` at the given index.
func (r *Reinterpret) GetAlternative(index int) *Alternative {
  if index >= 0 && index < len(r.Alternatives) {
    return r.Alternatives[index]
  }
  return nil
}

// Returns the "score" attribute value from the <alternative>.
func (a *Alternative) GetScore() string {
  return a.Score
}

// Returns the "level" attribute value from the <alternative>.
func (a *Alternative) GetLevel() string {
  return a.Level
}

// Returns the chardata from <alternative> as string.
func (a *Alternative) GetValue() string {
  return a.Value
}

// Returns the "word" attribute value from the <spellcheck>.
func (s *SpellCheck) GetWord() string {
  return s.Word
}

// Returns the "suggestion" attribute value from the <spellcheck>.
func (s *SpellCheck) GetSuggestion() string {
  return s.Suggestion
}

// Returns the "text" attribute value from the <spellcheck>.
func (s *SpellCheck) GetText() string {
  return s.Text
}

// Returns the "text" attribute value from the <delimiters>.
func (d *Delimiters) GetText() string {
  return d.Text
}

// Returns the "phrase" attribute value from the <translation>.
func (t *Translation) GetPhrase() string {
  return t.Phrase
}

// Returns the "trans" attribute value from the <translation>.
func (t *Translation) GetTranslation() string {
  return t.Trans
}

// Returns the "lang" attribute value from the <translation>.
func (t *Translation) GetLanguage() string {
  return t.Lang
}

// Returns the "text" attribute value from the <translation>.
func (t *Translation) GetText() string {
  return t.Text
}
