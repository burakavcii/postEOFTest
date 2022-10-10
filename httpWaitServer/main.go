package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/PostEOF", postEOF)
	http.HandleFunc("/UnexpectedEOF", unexpectedEOF)

	http.ListenAndServe(":54872", nil)
}

func postEOF(w http.ResponseWriter, _ *http.Request) {
	//w.Write([]byte("{\"test\":\"t\"}"))

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	conn.Close()
}

func unexpectedEOF(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"test\":\"t\"}"))

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	conn.Close()
}

//func getRoot(w http.ResponseWriter, r *http.Request) {
//	//log.Println("received request")
//	//w.Header().Set("Connection", "Keep-Alive")
//	//time.Sleep(time.Second * 3)
//	//panic("test")
//	hj, ok := w.(http.Hijacker)
//	if !ok {
//		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
//		return
//	}
//	conn, _, err := hj.Hijack()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//r.Close = true
//	//for i := 0; i < 255; i++ {
//	//	w.Write([]byte("{\"test\":\"t\"}"))
//	//}
//	//panic("test")
//
//	conn.Close()
//}
