#/bin/bash

curl -X PUT http://user:password@localhost:5984/mydatabase/_design/songitembyurl_view -d @songitembyurl_view.json
curl -X PUT http://user:password@localhost:5984/mydatabase/_design/songlist_view -d @songlist_view.json
