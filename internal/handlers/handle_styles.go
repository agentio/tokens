package handlers

import "net/http"

func StylesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/css; charset=UTF-8")
	w.Write([]byte(styles))
}

const styles = `

body {
 


    --font-sans: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji",
      "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    --font-mono: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono",
      "Courier New", monospace;

  font-family: var(--font-sans);

  --color-neutral:  rgb(255, 255, 255)
  /* Stone */
  --color-neutral-50:  rgb(250, 250, 249);
  --color-neutral-100:  rgb(245, 245, 244);
  --color-neutral-200:  rgb(231, 229, 228);
  --color-neutral-300:  rgb(214, 211, 209);
  --color-neutral-400:  rgb(168, 162, 158);
  --color-neutral-500:  rgb(120, 113, 108);
  --color-neutral-600:  rgb(87, 83, 78);
  --color-neutral-700: rgb( 68, 64, 60);
  --color-neutral-800:  rgb(41, 37, 36);
  --color-neutral-900:  rgb(28, 25, 23);
  /* Orange */
  --color-primary-50:  rgb(255, 247, 237);
  --color-primary-100:  rgb(255, 237, 213);
  --color-primary-200:  rgb(254, 215, 170);
  --color-primary-300:  rgb(253, 186, 116);
  --color-primary-400:  rgb(251, 146, 60);
  --color-primary-500: rgb(249, 115, 22);
  --color-primary-600:  rgb(234, 88, 12);
  --color-primary-700:  rgb(194, 65, 12);
  --color-primary-800:  rgb(154, 52, 18);
  --color-primary-900:  rgb(124, 45, 18);
  /* Rose */
  --color-secondary-50:  rgb(255, 241, 242);
  --color-secondary-100:  rgb(255, 228, 230);
  --color-secondary-200:  rgb(254, 205, 211);
  --color-secondary-300:  rgb(253, 164, 175);
  --color-secondary-400:  rgb(251, 113, 133);
  --color-secondary-500:  rgb(244, 63, 94);
  --color-secondary-600:  rgb(225, 29, 72);
  --color-secondary-700:  rgb(190, 18, 60);
  --color-secondary-800:  rgb(159, 18, 57);
  --color-secondary-900:  rgb(136, 19, 55);
  
  --border-color: #ddd;
  --gray-100: #fafafa;
  --gray-500: #666;
  --gray-700: #333;
  --primary-100: #d2e7ff;
  --primary-200: #b1d3fa;
  --primary-400: #2e8fff;
  --primary-500: #0078ff;
  --primary-600: #0066db;
  --error-500: #f00;
  --error-100: #fee;
}

/*
  Josh's Custom CSS Reset
  https://www.joshwcomeau.com/css/custom-css-reset/
*/
*, *::before, *::after {
  box-sizing: border-box;
}
* {
  margin: 0;
}
body {
  background-color: var(--color-neutral-200);
  line-height: 1.5;
  -webkit-font-smoothing: antialiased;
}
img, picture, video, canvas, svg {
  display: block;
  max-width: 100%;
}
input, button, textarea, select {
  font: inherit;
}
p, h1, h2, h3, h4, h5, h6 {
  overflow-wrap: break-word;
}
#root, #__next {
  isolation: isolate;
}

/*
  Common components
*/
button, .button {
  display: inline-block;
  border: 0;
  background-color: var(--color-primary-500);
  border-radius: 50px;
  color: #fff;
  padding: 2px 10px;
  cursor: pointer;
  text-decoration: none;
}
button:hover, .button:hover {
  background: var(--color-primary-400);
}

/*
  Custom components
*/
.error {
  background-color: var(--error-100);
  color: var(--error-500);
  text-align: center;
  padding: 1rem;
  display: none;
}
.error.visible {
  display: block;
}

#header {
  text-align: center;
  padding: 0.5rem 0 0;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin: 0 auto;
  max-width: 600px;
  padding: 20px;
}

.card {
 
  /* border: 1px solid var(--border-color); */
  border-radius: 6px;
  padding: 10px 16px;
  background-color: #fff;
}
.card > :first-child {
  margin-top: 0;
}
.card > :last-child {
  margin-bottom: 0;
}

.session-form {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.signin-form {
  display: flex;
  flex-direction: row;
  gap: 6px;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  padding: 10px 16px;
  background-color: #fff;
}

.signin-form input {
  flex: 1;
  border: 0;
}

.status-options {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 8px;
  margin: 10px 0;
}

.status-option {
  font-size: 2rem;
  width: 3rem;
  height: 3rem;
  padding: 0;
  background-color: #fff;
  border: 1px solid var(--border-color);
  border-radius: 3rem;
  text-align: center;
  box-shadow: 0 1px 4px #0001;
  cursor: pointer;
}

.status-option:hover {
  background-color: var(--primary-100);
  box-shadow: 0 0 0 1px var(--primary-400);
}

.status-option.selected {
  box-shadow: 0 0 0 1px var(--primary-500);
  background-color: var(--primary-100);
}

.status-option.selected:hover {
  background-color: var(--primary-200);
}

.status-line {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
  position: relative;
  margin-top: 15px;
}

.status-line:not(.no-line)::before {
  content: '';
  position: absolute;
  width: 2px;
  background-color: var(--border-color);
  left: 1.45rem;
  bottom: calc(100% + 2px);
  height: 15px;
}

.status-line .status {
  font-size: 2rem;
  background-color: #fff;
  width: 3rem;
  height: 3rem;
  border-radius: 1.5rem;
  text-align: center;
  border: 1px solid var(--border-color);
}

.status-line .desc {
  color: var(--gray-500);
}

.status-line .author {
  color: var(--gray-700);
  font-weight: 600;
  text-decoration: none;
}

.status-line .author:hover {
  text-decoration: underline;
}

.signin-cta {
  margin-top: 1em;
  text-align: center;
  text-wrap: balance;
}

`
