package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eirsyl/statuspage/pkg"
	"io/ioutil"
)

func (a *API) ListServices() ([]pkg.Service, error) {
	request := a.CreateRequest("/api/services", "GET", nil)
	response, err := a.client.Do(request)
	if err == nil && response.StatusCode == 200 {
		var result []pkg.Service
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []pkg.Service{}, err
		}
		json.Unmarshal(body, &result)
		return result, nil
	}

	if err == nil && response.StatusCode != 200 {
		return []pkg.Service{}, errors.New(fmt.Sprintf("Response failed %s", response.StatusCode))
	}

	return []pkg.Service{}, err
}
