// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2018 Gennady Trafimenkov <gennady.trafimenkov@gmail.com>

package hetzner

// ServerIPMap is a mapping of server ip addresses to server instances
type ServerIPMap map[string]*ServerSummary

// ServerListToIPMap creates map of servers from the list
func ServerListToIPMap(servers []*ServerSummary) ServerIPMap {
	m := ServerIPMap{}
	for _, s := range servers {
		m[s.ServerIP] = s
	}
	return m
}
