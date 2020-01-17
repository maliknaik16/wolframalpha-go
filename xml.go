
package wolfram

import(
	"encoding/xml"
)

type QueryResult struct {
	// Contains the entire API result.
	XMLName xml.Name `xml:"queryresult"`

	// true or false depending on whether the input could be successfully
	// understood. If false, there will be no <pod> subelements.
	Success string `xml:"success,attr"`

	// true or false depending on whether a serious processing error occurred,
	// such as a missing required parameter. If true, there will be no pod
	// content, just an <error> subelement
	Error string `xml:"error,attr"`

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

	// Each pod contains a piece or category of information about the given query.
	Pods []Pod `xml:"pod"`
}

type Pod struct {

	// <pod> elements are the main output of the Full Results API
	XMLName xml.Name `xml:"pod"`

	// The pod title, used to identify the pod and its contents.
	Title  string `xml:"title,attr"`

	//true or false depending on whether a serious processing error occurred
	// with this specific pod. If true, there will be an <error> subelement.
	Error string `xml:"error,attr"`

	// A number indicating the intended position of the pod in a visual display.
	// These numbers are typically multiples of 100, and they form an increasing
	// sequence from top to bottom.
	Position string `xml:"position,attr"`

	// The name of the scanner that produced this pod. A general guide to the
	// type of data it holds.
	Scanner string `xml:"scanner,attr"`

	// A unique identifier for a pod, used for selecting specific pods to include
	// or exclude.
	ID string `xml:"id,attr"`

	// The number of subpod elements present.
	NumSubPods string `xml:"numsubpods,attr"`

	// Subelements of <pod> that contain the results for a single subpod.
	// <subpod> has a title attribute, which is usually an empty string because
	// most subpods have no title.
	SubPods []SubPod `xml:"subpod"`
}

type SubPod struct {
}
