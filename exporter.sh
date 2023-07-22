#/usr/bin/env bash
#
#
#This is a bit of a setup to export the server
export port=8001
#export API_SERVICE="${port}-${HOSTNAME}.cluster-lknrrkkitbcdsvoir6wqg4mwt6.cloudworkstations.dev"
export API_SERVICE="${port}-${WEB_HOST}"
echo $API_SERVICE
