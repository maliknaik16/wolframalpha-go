
package wolfram

import (
	"testing"
	"fmt"
)

// Tests the GetSrc() method.
func TestGetSrc(t *testing.T) {
	img := queryResult.GetPod(1).SubPods[0].Image

	if url := img.GetSrc(); url == "" {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", "", url))
	}
}

// Tests the GetAlt() method.
func TestGetAlt(t *testing.T) {
	img := queryResult.GetPod(0).SubPods[0].Image
	expected := "tides | Seattle, Washington, United States"

	if alt := img.GetAlt(); alt != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, alt))
	}
}

// Tests the GetTitle() method.
func TestGetTitle(t *testing.T) {
	img := queryResult.Pods[0].SubPods[0].Image
	expected := "tides | Seattle, Washington, United States"

	if title := img.GetTitle(); title != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, title))
	}
}

// Tests the GetDimensions() method.
func TestGetDimensions(t *testing.T) {
	img := queryResult.Pods[1].SubPods[0].Image
	expected := Dimension{400, 396}

	if dimensions := img.GetDimensions(); dimensions.Width != expected.Width || dimensions.Height != expected.Height {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, dimensions))
	}
}

// Tests the GetType() method.
func TestGetType(t *testing.T) {
	img := queryResult.Pods[1].SubPods[0].Image
	expected := "Grid"

	if itype := img.GetType(); itype != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, itype))
	}
}

// Tests the GetThemes() method.
func TestGetThemes(t *testing.T) {
	img := queryResult.Pods[1].SubPods[0].Image
	expected := "1,2,3,4,5,6,7,8,9,10,11,12"

	if themes := img.GetThemes(); themes != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, themes))
	}
}

// Tests the IsColorInvertable() method.
func TestIsColorInvertable(t *testing.T) {
	img := queryResult.Pods[1].SubPods[0].Image
	expected := true

	if colorInvertable := img.IsColorInvertable(); colorInvertable != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, colorInvertable))
	}
}
// Tests the GetFormat() method.
func TestGetFormat(t *testing.T) {
	img := queryResult.Pods[0].SubPods[0].Image
	expected := FORMAT_GIF

	if format := img.GetFormat(); format != expected {
		t.Error(fmt.Sprintf("Test Failed: \"%v\" Expected, \"%v\" Received.\n", expected, format))
	}
}
