package snippets

import (
	"net/http"
)

type empty struct{}
type semaphore chan empty

func main() {
	http.HandleFunc("/check-window-status", func(rw http.ResponseWriter, _ *http.Request) {
		_, _ = rw.Write([]byte(`{"isWindowOpen":false}`))
	})

	_ = http.ListenAndServe(":8000", nil)

	//time.Date()
	//time.Now().Before()
	//time.Now().After()
}
