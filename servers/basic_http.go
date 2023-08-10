package servers

import (
	"fmt"
	"net/http"
)

func test1(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Fprint(w, "test1\n")
}

func test2(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Fprint(w, "test2\n")
}

func BasicHttpServer() {
	fs := http.FileServer(http.Dir("./static"))
	fs2 := http.FileServer(http.Dir("./static/dir"))
	http.Handle("/", fs)
	http.Handle("/dir", fs2)

	http.HandleFunc("/test1", test1)
	http.HandleFunc("/test2", test2)
	if err := http.ListenAndServe(":8899", nil); err != nil {
		fmt.Println("some error")
		fmt.Println(err)
	}
}
