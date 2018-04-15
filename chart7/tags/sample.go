package main

import "net/url"

type Route struct {
	remoteID     string   `json:"remote_id" xml:"remote_id"`
	didSolicit   bool     `json:"did_solicit" xml:"did_solicit"`
	retry        bool     `json:"retry" xml:"retry"`
	url          *url.URL `json:"url" xml:"url"`
	authRequired bool     `json:"auth_required" xml:"auth_required"`
	tlsRequired  bool     `json:"tls_required" xml:"tls_required"`
	closed       bool     `json:"closed" xml:"closed"`
	Credentials  struct {
		Username string `json:"username" xml:"username"`
		Password string `json:"password" xml:"password"`
	} `json:"credentials" xml:"credentials"`
}
