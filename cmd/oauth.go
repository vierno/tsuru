// Copyright 2014 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func clientID() string {
	return os.Getenv("TSURU_AUTH_CLIENTID")
}

func localServer() {
	func() {
		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Test")
		})
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=http://localhost:4242/callback&scope=user:email", clientID())
			http.Redirect(w, r, url, 302)
		})
		http.ListenAndServe(":4242", nil)
	}()
	exec.Command("open", "http://localhost:4242").Start()
}

func oauthLogin(context *Context, client *Client) error {
	localServer()
	return nil
}
