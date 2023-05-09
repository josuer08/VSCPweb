// healthHandlers package is the package with all of the handlers for the healt
// route of the system, it gets info from all of the workers of the cluster as
// well as the master server
package healthHandlers

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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ëverything should be okay my boi")
}
