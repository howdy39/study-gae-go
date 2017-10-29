package auth

import (
	"net/http"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/welcome", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	ctx := appengine.NewContext(r)

	u := user.Current(ctx)

	if u == nil {
		loginUrl, _ := user.LoginURL(ctx, "welcome")
		fmt.Fprintf(w,  `<a href="%s">Sign in</a>`, loginUrl)
		return
	}

	logoutUrl, _ := user.LogoutURL(ctx, "welcome")
	fmt.Fprintf(w, `Welcome, %s!(<a href="%s">Sign out</a>))`, u, logoutUrl)
}