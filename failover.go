// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package hetzner

import (
	"fmt"
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#failover

// FailoverService represents a service to work with failover ips.
type FailoverService interface {
	List() ([]*FailoverSummary, *http.Response, error)
	Switch(failoverIP string, newActiveIP string) (*FailoverSummary, *http.Response, error)
}

type failoverServiceImpl struct {
	client *Client
}

// List returns the list of failover ips
func (s *failoverServiceImpl) List() ([]*FailoverSummary, *http.Response, error) {
	path := "/failover"

	type Data struct {
		Failover *FailoverSummary `json:"failover"`
	}
	data := make([]Data, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data, true)

	a := make([]*FailoverSummary, len(data))
	for i, d := range data {
		a[i] = d.Failover
	}
	return a, resp, err
}

// Switch switches traffic of failoverIP to the server with ip newActiveIP.
// When successful, returns updated information about the failover IP
func (s *failoverServiceImpl) Switch(failoverIP string, newActiveIP string) (*FailoverSummary, *http.Response, error) {
	path := fmt.Sprintf("/failover/%v", failoverIP)

	req := failoverSwitchRequest{ActiveServerIP: newActiveIP}

	type Data struct {
		Failover *FailoverSummary `json:"failover"`
	}
	data := Data{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data, true)
	return data.Failover, resp, err
}
