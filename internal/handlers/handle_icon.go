package handlers

import "net/http"

func IconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "image/svg+xml")
	w.Write([]byte(icon))
}

const icon = `
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 2200 2200">
  <circle cx="1100" cy="1100" r="1100" fill="#222" />
  <circle cx="1100" cy="1100" r="990" fill="#ccc" />
  <path fill="#222" 
    d="M100,1100 a1000,1000 0 0 0 2000,0 a500,500 0 0 0 -1000,0 a500,500 0 0 1 -1000,0"/>
  <rect x="520" y="800" width="160" height="600" fill="#222"/>
  <circle r="330" cx="1600" cy="1100" fill="#ccc"/>
  <circle r="200" cx="1600" cy="1100" fill="#222"/>
</svg>
`
