// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package hetzner

// FailoverSummary contains information about failover ip
type FailoverSummary struct {
	IP             string `json:"ip"`
	Netmask        string `json:"netmask"`
	ServerIP       string `json:"server_ip"`
	ServerNumber   int    `json:"server_number"`
	ActiveServerIP string `json:"active_server_ip"`
}

// failoverSwitchRequest describes parameters of POST /failover API call
type failoverSwitchRequest struct {
	ActiveServerIP string `url:"active_server_ip"`
}
