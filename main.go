package main

import "net/http"

var formPage = `
<!docype html>
<html>
<head>
</head>
<body>
<form method="post" action="/submit">
<input type="text" placeholder="favorite color" name="color" />
<input type="submit" value="Submit" />
</form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(formPage))
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Request method: " + r.Method))
	})

	http.ListenAndServe(":8080", nil)
}
