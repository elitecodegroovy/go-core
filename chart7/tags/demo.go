package main

import "net/url"

type Route struct {
	remoteID     string
	didSolicit   bool
	retry        bool
	url          *url.URL
	authRequired bool
	tlsRequired  bool
	closed       bool
	Credentials struct {
		Username string
		Password string
	}
}