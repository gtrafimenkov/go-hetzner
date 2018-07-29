// SPDX-License-Identifier: Apache-2.0

package hetzner

type Rescue struct {
	ServerIP      string           `json:"server_ip"`
	ServerNumber  int              `json:"server_number"`
	Os            []string         `json:"os"`
	Arch          []int            `json:"arch"`
	Active        bool             `json:"active"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

type Linux struct {
	ServerIP      string           `json:"server_ip"`
	ServerNumber  int              `json:"server_number"`
	Dist          []string         `json:"dist"`
	Arch          []int            `json:"arch"`
	Lang          []string         `json:"lang"`
	Active        bool             `json:"active"`
	Password      *string          `json:"password"`
	AuthorizedKey []*AuthorizedKey `json:"authorized_key"`
	HostKey       []*HostKey       `json:"host_key"`
}

type Vnc struct {
	ServerIP     string   `json:"server_ip"`
	ServerNumber int      `json:"server_number"`
	Dist         []string `json:"dist"`
	Arch         []int    `json:"arch"`
	Lang         []string `json:"lang"`
	Active       bool     `json:"active"`
	Password     *string  `json:"password"`
}

type Windows struct {
	ServerIP     string   `json:"server_ip"`
	ServerNumber int      `json:"server_number"`
	Dist         []string `json:"dist"`
	Lang         []string `json:"lang"`
	Active       bool     `json:"active"`
	Password     *string  `json:"password"`
}

type Plesk struct {
	ServerIP     string   `json:"server_ip"`
	ServerNumber int      `json:"server_number"`
	Dist         []string `json:"dist"`
	Arch         []int    `json:"arch"`
	Lang         []string `json:"lang"`
	Active       bool     `json:"active"`
	Password     *string  `json:"password"`
	Hostname     *string  `json:"hostname"`
}

type Cpanel struct {
	ServerIP     string   `json:"server_ip"`
	ServerNumber int      `json:"server_number"`
	Dist         []string `json:"dist"`
	Arch         []int    `json:"arch"`
	Lang         []string `json:"lang"`
	Active       bool     `json:"active"`
	Password     *string  `json:"password"`
	Hostname     *string  `json:"hostname"`
}

type Boot struct {
	Rescue  *Rescue  `json:"rescue"`
	Linux   *Linux   `json:"linux"`
	Vnc     *Vnc     `json:"vnc"`
	Windows *Windows `json:"windows"`
	Plesk   *Plesk   `json:"plesk"`
	Cpanel  *Cpanel  `json:"cpanel"`
}

type ActivateLinuxRequest struct {
	ServerIP      string
	Dist          string   `url:"dist"`
	Arch          int      `url:"arch"`
	Lang          string   `url:"lang"`
	AuthorizedKey []string `url:"authorized_key,brackets"`
}
