package handlers

import (
	"html/template"
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("signin").Parse(signin_template)
	if err != nil {
		return
	}
	err = t.ExecuteTemplate(w, "signin", map[string]any{})
	if err != nil {
		return
	}
}

const signin_template = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Sign in</title>
<link rel="stylesheet" href="/public/styles.css">
</head>
<body>
<div id="root">
<div id="header">
<img src="/public/icon.svg" width="144px" style="margin: 0 auto"/>
</div>
<div class="container">

<form action="/@/signin" method="post" class="signin-form">
<input type="text" name="handle" placeholder="Sign in with your Bluesky handle (eg alice.bsky.social)" required="">
<button type="submit">Sign in</button>
</form>

<div class="signin-cta">
To learn more, visit <a href="https://agent.io/posts/identity">agent.io</a>
</div>
</div>
</div>
</body>
</html>
`
