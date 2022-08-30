package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

    //example of simple http router, fased out in favor of gorilla mux
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello World!")
	//})
	//http.ListenAndServe(":8000", nil)



    r := mux.NewRouter()
    r.HandleFunc("/example/{first}/{second:[0-9]+}", exampleHandler)
    //r.HandleFunc("/", )
    //r.HandleFunc("/access", )
    //r.HandleFunc("/stats", )
    //r.HandleFunc("/topology", )
    //r.HandleFunc("/health", )
    //r.HandleFunc("/free", )//probably do this with server side or sockets
    //r.HandleFunc("/flowdel", )//this is an external request
    //r.HandleFunc("/mpstat", )
    //r.HandleFunc("/ifstat", )
    //r.HandleFunc("/showtemp", )
    //r.HandleFunc("/gettopo", )
    //r.HandleFunc("/net", )
    //r.HandleFunc("/rpiping", )//workerping
    //r.HandleFunc("/nodes", )
    //r.HandleFunc("/statusnodes", )
    //r.HandleFunc("/intfs", )
    //r.HandleFunc("/iperf", )
    //r.HandleFunc("/pingall", )
    //r.HandleFunc("/placement", )
    //r.HandleFunc("/getvsorcdata", )
    //r.HandleFunc("/getcontrollerdata", )
    //r.HandleFunc("/resetflows", )
    //r.HandleFunc("/listswitch", )
    //r.HandleFunc("/status", )
    //r.HandleFunc("/tablestatus", )
    //r.HandleFunc("/portsdesc", )
    //r.HandleFunc("/portsstat", )
    //r.HandleFunc("/startcontroller", )
    //r.HandleFunc("/startcontrollerrouter", )
    //r.HandleFunc("/stopcontroller", )
    //r.HandleFunc("/sendcommand", )
    //r.HandleFunc("/cancel", )
    //r.HandleFunc("/startvsorc", )
    //r.HandleFunc("/stopvsorc", )
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)



}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "First: %v\nSecond: %v", vars["first"], vars["second"])
}
