package web

import (
	"fmt"
	"net/http"
	"testing"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是用户")
}

func Test_web(t *testing.T) {
	http.HandleFunc("/home", home)
	http.HandleFunc("/user", user)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
	}
}
