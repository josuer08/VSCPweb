package staticHandlers
import (
        "fmt"

                "net/http"
                )

//TODO(josuer08): get mpstat idle stat for every machine, as well as the ifstat
//and also the temperatures of every one of the servers working for the cluster.

// This might include a textproto where we have all of the info of every node and
// then with their addresses we can go ahead and query that info from them with
// the special agent that each of them will already have installed in order to
// connect to the cluster.
func MasterHealthHandler(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf("Ëverything should be okay my boi")
}

