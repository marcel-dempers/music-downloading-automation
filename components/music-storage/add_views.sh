#/bin/bash

curl -x PUT http://localhost:5984/mydatabase/_design/songitembyurl_view -d @songitem.json