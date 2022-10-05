package main

import (
	"fmt"
	"net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
//create a new server and mount the handlers:
    s := CreateNewServer()
    s.MountHandlers()
//staring up the server:
    http.ListenAndServe(":8000", s.Router)
}


type server struct {
    Router *chi.Mux
    //suggestion to add config settings or DB in here
}

func CreateNewServer() *server {
    s := &server{}
    s.Router = chi.NewRouter()
    return s
}

func (s *server) MountHandlers() {

//Creation of the router:
    //logging enables as a middleware
    s.Router.Use(middleware.Logger)
    //recover from panics and send a 500 internal error
    s.Router.Use(middleware.Recoverer)
    //a middleware to check if the server is alive
    s.Router.Use(middleware.Heartbeat("/ping"))
    //a profiler to check healt of server
    s.Router.Mount("/debug", middleware.Profiler())
    //personalized 404
    s.Router.NotFound(GenericHandler404)
    //personalized 405
    s.Router.MethodNotAllowed(GenericHandler405)

//example:
    //example of a get function and its handler
    s.Router.Get("/example/{first}/{second:[0-9]+}", ExampleHandler)

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
    //s.Router.Connect(pattern string, h http.HandlerFunc)
    //s.Router.Delete(pattern string, h http.HandlerFunc)
    //s.Router.Get(pattern string, h http.HandlerFunc)
    //s.Router.Head(pattern string, h http.HandlerFunc)
    //s.Router.Options(pattern string, h http.HandlerFunc)
    //s.Router.Patch(pattern string, h http.HandlerFunc)
    //s.Router.Post(pattern string, h http.HandlerFunc)
    //s.Router.Put(pattern string, h http.HandlerFunc)
    //s.Router.Trace(pattern string, h http.HandlerFunc)
///////////////////////////////////////////////////////////////////////////////

//generic ryu version at first...
    //s.Router.Get("/", )
    //s.Router.Get("/access", )
    //s.Router.Get("/stats", )
    //s.Router.Get("/topology", )
    //s.Router.Get("/mpstat", )//DEPRECATED see: /masterHealth
    //s.Router.Get("/ifstat", )//DEPRECATED see: /masterHealth
    //s.Router.Get("/showtemp", )//DEPRECATED see: /masterHealth
    //healthRouter.Get("/", )//this renders the health page "/health"
    //healthRouter.Get("/masterHealth", )//return a JSON with healt of the master "/free"
    //controllerRouter.Delete("/flow", )//this is an external request "/flowdel"
    //controllerRouter.Get("/topology", )// "/gettopo"
    //mininetApiRouter.Get("/net", )
    //healthRouter.Get("/ping", )// "/rpiping"
    //mininetApiRouter.Get("/nodes", )
    //mininetApiRouter.Get("/status", )// "/statusnodes"
    //mininetApiRouter.Get("/intfs", )
    //mininetApiRouter.Get("/iperf", )
    //mininetApiRouter.Get("/pingall", )
    //mininetApiRouter.Get("/placement", )
    //managementRouter.Get("/vscpData", )// "/getvsorcdata"
    //managementRouter.Get("/controllerData", ) // "/getcontrollerdata"
    //controllerRouter.Get("/resetflows", )//not sure if here or managementRouter
    //controllerRouter.Get("/listswitch", )
    //controllerRouter.Get("/status", )
    //controllerRouter.Get("/tablestatus", )
    //controllerRouter.Get("/portDescription", )// "/portsdesc"
    //controllerRouter.Get("/portStatus", )// "/portsstat"
    //managementRouter.Get("/startController", )// "/startcontroller"
    //managementRouter.Get("/startcontrollerAPI", )// "/startcontrollerrouter"
    //managementRouter.Get("/stopController", )// "/stopcontroller"
    //mininetApiRouter.Get("/sendCommand", )// "/sendcommand"
    //managementRouter.Get("/cancel", )// "/cancel"
    //managementRouter.Get("/startVsorc", )// "/startvsorc"
    //managementRouter.Get("/stopVsorc", )// "/stopvsorc"
    //http.Handle("/", r)

//Mounting all of the subrouters:
    s.Router.Mount("/health", healthRouter)
    s.Router.Mount("/controller", controllerRouter)
    s.Router.Mount("/management", managementRouter)
    s.Router.Mount("/virtualAPI", mininetApiRouter)

}



//GenericHandler404 is the universal 404 response of this front end
func GenericHandler404(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(404)
    w.Write([]byte("route does not exist"))
}
//GenericHandler405 is the universal 405 response of this front end
func GenericHandler405(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(405)
    w.Write([]byte("Method not valid"))
}
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    var1 := chi.URLParam(r, "first")
    var2 := chi.URLParam(r, "second")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "First: %v\nSecond: %v", var1, var2)
}
