package main

import (
	"fmt"
	"net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {


//Creation of the router:
    //r main router for VSCP web application
    r := chi.NewRouter()
    //logging enables as a middleware
    r.Use(middleware.Logger)
    //a middleware to check if the server is alive
    r.Use(middleware.Heartbeat("/ping"))
    //a profiler to check healt of server
    r.Mount("/debug", middleware.Profiler())
    //recover from panics and send a 500 internal error
    r.Use(middleware.Recoverer)
    //personalized 404
    r.NotFound(genericHandler404)
    //personalized 405
    r.MethodNotAllowed(genericHandler405)

//example:
    //example of a get function and its handler
    r.Get("/example/{first}/{second:[0-9]+}", exampleHandler)

//Creating subrouters:
    //healthRouter check on the health of nodes or main server
    healthRouter := chi.NewRouter()
    //controllerRouter Agnostic openVswitch controller bridge
    controllerRouter := chi.NewRouter()
    //managementRouter Start, stop, monitoring of the core (mininet)
    managementRouter := chi.NewRouter()
    //mininetApiRouter Interface with the virtual environment inside the core (mininet)
    mininetApiRouter := chi.NewRouter()
    //might want to check https://go-chi.io/#/pages/routing?id=routing-groups
    //in order to make groups where you have other middleware like authentication

//////////////////////////////////possible routing/////////////////////////////
    //r.Connect(pattern string, h http.HandlerFunc)
    //r.Delete(pattern string, h http.HandlerFunc)
    //r.Get(pattern string, h http.HandlerFunc)
    //r.Head(pattern string, h http.HandlerFunc)
    //r.Options(pattern string, h http.HandlerFunc)
    //r.Patch(pattern string, h http.HandlerFunc)
    //r.Post(pattern string, h http.HandlerFunc)
    //r.Put(pattern string, h http.HandlerFunc)
    //r.Trace(pattern string, h http.HandlerFunc)
///////////////////////////////////////////////////////////////////////////////

    //r.Get("/", )
    //r.Get("/access", )
    //r.Get("/stats", )
    //r.Get("/topology", )
    //r.Get("/health", )
    //r.Get("/free", )//probably do this with server side or sockets
    //r.Get("/flowdel", )//this is an external request
    //r.Get("/mpstat", )
    //r.Get("/ifstat", )
    //r.Get("/showtemp", )
    //r.Get("/gettopo", )
    //r.Get("/net", )
    //r.Get("/rpiping", )//workerping
    //r.Get("/nodes", )
    //r.Get("/statusnodes", )
    //r.Get("/intfs", )
    //r.Get("/iperf", )
    //r.Get("/pingall", )
    //r.Get("/placement", )
    //r.Get("/getvsorcdata", )
    //r.Get("/getcontrollerdata", )
    //r.Get("/resetflows", )
    //r.Get("/listswitch", )
    //r.Get("/status", )
    //r.Get("/tablestatus", )
    //r.Get("/portsdesc", )
    //r.Get("/portsstat", )
    //r.Get("/startcontroller", )
    //r.Get("/startcontrollerrouter", )
    //r.Get("/stopcontroller", )
    //r.Get("/sendcommand", )
    //r.Get("/cancel", )
    //r.Get("/startvsorc", )
    //r.Get("/stopvsorc", )
    //http.Handle("/", r)

//Mounting all of the subrouters:
    r.Mount("/health", healthRouter)
    r.Mount("/controller", controllerRouter)
    r.Mount("/management", managementRouter)
    r.Mount("/virtualAPI", mininetApiRouter)

//staring up the server
    http.ListenAndServe(":8000", r)
}









//genericHandler404 is the universal 404 response of this front end
func genericHandler404(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(404)
    w.Write([]byte("route does not exist"))
}
//genericHandler405 is the universal 405 response of this front end
func genericHandler405(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(405)
    w.Write([]byte("Method not valid"))
}
func exampleHandler(w http.ResponseWriter, r *http.Request) {
    var1 := chi.URLParam(r, "first")
    var2 := chi.URLParam(r, "second")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "First: %v\nSecond: %v", var1, var2)
}
