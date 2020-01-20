
package wolfram

import (
	"fmt"
	"testing"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

var queryResult QueryResult = QueryResult{}

func init() {

	resp, err := http.Get("http://api.wolframalpha.com/v2/query?appid=DEMO&input=tides%20seattle")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(byteValue, &queryResult)
}

func TestQueryResult(t *testing.T) {
	i := 1
	if i == 2 {
		t.Error("Error")
	}
}
