package main
//
//import (
//	"fmt"
//	"log"
//	"markdown/utils/router"
//	"net/http"
//	"regexp"
//)
//
//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Welcome!\n")
//}
//
//func Hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello")
//}
//
//func main() {
//	router := router.New()
//	router.Handle("/", Index)
//	router.Handle("/hello/:name", Hello)
//
//	match("/", "/hhh")
//	match("/", "/")
//	match("/hello/:name", "/hello/faceair")
//	match("/hello/:name/question", "/hello/faceair")
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}
//
//func match(pattern, path string) bool {
//	tmpPattern := regexp.MustCompile(`:\w+`).ReplaceAllString(pattern, "\\w+")
//	matched, _ := regexp.MatchString("^"+tmpPattern+"$", path)
//	return matched
//}
