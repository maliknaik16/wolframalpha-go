
package wolfram

import(
	"encoding/xml"
)

type QueryResult struct {
	// Contains the entire API result.
	XMLName xml.Name `xml:"queryresult"`

	Text string `xml:",chardata"`

	LanguageMsg LanguageMsg `xml:"languagemsg"`

	Tips Tips `xml:"tips"`

	FutureTopic FutureTopic `xml:"futuretopic"`

	// true or false depending on whether the input could be successfully
	// understood. If false, there will be no <pod> subelements.
	Success string `xml:"success,attr"`

	// true or false depending on whether a serious processing error occurred,
	// such as a missing required parameter. If true, there will be no pod
	// content, just an <error> subelement
	ErrorAttr string `xml:"error,attr"`

	// The number of pods.
	Numpods  string `xml:"numpods,attr"`

	// The version specification of the API on the server that produced this result.
	Version  string `xml:"version,attr"`

	// Categories and types of data represented in the results (e.g. "Financial").
	DataTypes  string `xml:"datatypes,attr"`

	// The wall-clock time in seconds required to generate the output.
	Timing  string `xml:"timing,attr"`

	// The number of pods that are missing because they timed out.
	TimedOut  string `xml:"timedout,attr"`

	// The time in seconds required by the parsing phase.
	ParseTiming  string `xml:"parsetiming,attr"`

	// Whether the parsing stage timed out.
	ParseTimedOut  string `xml:"parsetimedout,attr"`

	// A URL to use to recalculate the query and get more pods.
	Recalculate  string `xml:"recalculate,attr"`

	TimedOutPods  string   `xml:"timedoutpods,attr"`

	ID string `xml:"id,attr"`

	Host string `xml:"host,attr"`

	Server string `xml:"server,attr"`

	Related string `xml:"related,attr"`

	// Each pod contains a piece or category of information about the given query.
	Pods []Pod `xml:"pod"`

	Assumptions Assumptions `xml:"assumptions"`

	States States `xml:"states"`

	ExamplePage string `xml:"examplepage"`

	Warnings Warnings `xml:"warnings"`

	Sources Sources `xml:"sources"`

	Generalizaton Generalizaton `xml:"generalization"`

	DidYouMeans DidYouMeans `xml:"didyoumeans"`

	ErrorTag ErrorTag `xml:"error"`

	Sounds Sounds `xml:"sounds"`
}

type Generalizaton struct {
	XMLName xml.Name `xml:"generalization"`
	Text string `xml:",chardata"`
	Topic string `xml:"topic,attr"`
	Desc string `xml:"desc,attr"`
	Url string `xml:"url,attr"`
}

type DidYouMeans struct {
	XMLName xml.Name `xml:"didyoumeans"`
	Text string `xml:",chardata"`
	Count string `xml:"count,attr"`
	DidYouMean []DidYouMean `xml:"didyoumean"`
}

type DidYouMean struct {
	XMLName xml.Name `xml:"didyoumean"`
	Text string `xml:",chardata"`
	Score string `xml:"score,attr"`
	Level string `xml:"level,attr"`
}

type ErrorTag struct {
	XMLName xml.Name `xml:"error"`
	Text string `xml:",chardata"`
	Code string `xml:"code"`
	Msg string `xml:"msg"`
}

type LanguageMsg struct {
	XMLName xml.Name `xml:"languagemsg"`
	Text string `xml:",chardata"`
	English string `xml:"english,attr"`
	Other string `xml:"other,attr"`
}

type Tips struct {
	XMLName xml.Name `xml:"tips"`
	Text string `xml:",chardata"`
	Count string `xml:"count,attr"`
	Tip []Tip `xml:"tip"`
}

type Tip struct {
	XMLName xml.Name `xml:"tip"`
	Text string `xml:",chardata"`
	AttrText string `xml:"text,attr"`
}

type FutureTopic struct {
	XMLName xml.Name `xml:"futuretopic"`
	Text string `xml:",chardata"`
	Topic string `xml:"topic,attr"`
	Msg string `xml:"msg,attr"`
}
