package handlers

import (
	"html/template"
	"log"
	"net/http"
	"slices"

	"github.com/agentio/atiquette/api/app/bsky"
	"github.com/agentio/tokens/internal/clients"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").Parse(home_template)
	if err != nil {
		return
	}
	// When a user is signed-in, IO puts their did in this header.
	did := r.Header.Get("user-did")
	if did == "" {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	profile, err := bsky.ActorGetProfile(r.Context(), clients.AnonymousClient, did)
	if err != nil {
		log.Printf("%+v", err)
	}
	enrolled := slices.Contains(handles(), profile.Handle)
	err = t.ExecuteTemplate(w, "home", map[string]any{
		"Did":      did,
		"Name":     *profile.DisplayName,
		"Handle":   profile.Handle,
		"Enrolled": enrolled,
	})
	if err != nil {
		return
	}
}

const home_template = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Home</title>
<link rel="stylesheet" href="./public/styles.css">
</head>
<body>
<div id="root">
<div class="error"></div>
<div id="header">
<img src="/public/icon.svg" width="144px" style="margin: 0 auto"/>
</div>
<div class="container">

{{ if .Enrolled }}
<div class="card">
<h3 style="text-align: center;">Welcome to the IO Identity Server!</h3>
<br/>
<p  style="text-align: center;">You are signed in as {{ .Name }}</p>
<p  style="text-align: center;">(@{{ .Handle }})</p>
<br/>
<hr/>
<br/>
<div class="session-form" style="display:flex; align-items: flex-end;">
<div>
<form action="/generate" method="post"><button type="submit">Download a JWT</button></form>
</div>
<div>
<a href="/@/signout" class="button">Sign out</a>
</div>
</div>
</div>
<div class="signin-cta">
To learn more, visit <a href="https://agent.io/posts/identity">agent.io</a>
</div>
{{ else }}
<div class="card" style="text-align: center;">
<h3>Thanks for your interest!</h3>
<br/>
<p>We don't have anything for you at this time.</p>
<br/>
<a href="/@/signout" class="button">Sign out</a>
</div>
{{ end}}

</div>

</body>
</html>
`
