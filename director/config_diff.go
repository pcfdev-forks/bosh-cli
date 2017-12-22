package director

import (
	"encoding/json"
	"net/http"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type ConfigDiff struct {
	Diff [][]interface{}
}

type ConfigDiffResponse struct {
	Diff [][]interface{} `json:"diff"`
}

func NewConfigDiff(diff [][]interface{}) ConfigDiff {
	return ConfigDiff{
		Diff: diff,
	}
}

type DiffConfigError struct {
	Code        int    `json:code`
	Description string `json:description`
}

func (c Client) postConfigDiff(path string, manifest []byte, setHeaders func(*http.Request)) (ConfigDiffResponse, error) {
	var resp ConfigDiffResponse

	respBody, response, err := c.clientRequest.RawPost(path, manifest, setHeaders)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			// return empty diff, just for compatibility with directors which don't have the endpoint
			return resp, nil
		} else {
			//diffConfigError := DiffConfigError{}

			//subMessageErr := json.Unmarshal([]byte(err.Error()), &diffConfigError)
			//if subMessageErr == nil {
			//	return resp, bosherr.Error(diffConfigError.Description)
			//} else {
			//	return resp, bosherr.Error(subMessageErr.Error())
			//}
			return resp, bosherr.Error(err.Error())

			return resp, bosherr.WrapError(err, "Fetching diff result")
		}
	}

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return resp, bosherr.WrapError(err, "Unmarshaling Director response")
	}

	return resp, nil
}
