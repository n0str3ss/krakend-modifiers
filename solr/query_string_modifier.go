package solr

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/martian"
	"github.com/google/martian/parse"
)

func init() {
	parse.Register("solr.QueryModifier", queryModifierFromJSON)
}

// QueryModifier contains the private and public Marvel API key
type QueryModifier struct{}

// QueryModifierJSON to Unmarshal the JSON configuration
type QueryModifierJSON struct {
	Scope []parse.ModifierType `json:"scope"`
}

// ModifyRequest modifies the query string of the request with the given key and value.
func (m *QueryModifier) ModifyRequest(req *http.Request) error {
	/*	query := req.URL.Query()
		ts := strconv.FormatInt(time.Now().Unix(), 10)
		hash := GetMD5Hash(ts + m.private + m.public)
		query.Set("apikey", m.public)
		query.Set("ts", ts)
		query.Set("hash", hash)
		req.URL.RawQuery = query.Encode()
	*/
	var r map[string]interface{}

	msh, _ := json.Marshal(req.URL)

	_ = json.Unmarshal(msh, &r)

	return errors.New("bam")
}

// NewQueryModifier returns a request modifier that will set the query string
func NewQueryModifier() martian.RequestModifier {
	return &QueryModifier{}
}

// queryModifierFromJSON takes a JSON message as a byte slice and returns
// a modifier and an error.
//
// Example JSON:
// {
//  "scope": ["request", "response"]
// }
func queryModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &QueryModifierJSON{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(NewQueryModifier(), msg.Scope)
}
