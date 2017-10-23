# simple_kvdbmysql


## read 

curl 'http://localhost:8080/kv?limit=30&offset=0&sortby=Id&order=desc' -H 'Origin: http://localhost:8081' -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-US,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.63 Safari/537.36' -H 'Accept: application/json, text/plain, */*' -H 'Referer: http://localhost:8081/' -H 'Connection: keep-alive' --compressed

## insert 
curl 'http://localhost:8080/kv/new' -H 'Origin: http://localhost:8081' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.63 Safari/537.36' -H 'Content-Type: text/plain' -H 'Accept: application/json, text/plain, */*' -H 'Referer: http://localhost:8081/' -H 'Connection: keep-alive' --data-binary '"Url":"url","K":"key","V":"value","T":123}' --compressed

## update
curl 'http://localhost:8080/kv/769403753574932480' -X PUT -H 'Origin: http://localhost:8081' -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-US,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.63 Safari/537.36' -H 'Content-Type: text/plain' -H 'Accept: application/json, text/plain, */*' -H 'Referer: http://localhost:8081/' -H 'Connection: keep-alive' --data-binary '{"IdStr":"769403753574932480","K":"key","Url":"url","V":"valueupdate","T":123}}' --compressed

## delete
curl 'http://localhost:8080/kv/769403753574932480' -X DELETE -H 'Origin: http://localhost:8081' -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-US,en;q=0.8' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.63 Safari/537.36' -H 'Content-Type: text/plain' -H 'Accept: application/json, text/plain, */*' -H 'Referer: http://localhost:8081/' -H 'Connection: keep-alive' --data-binary '{"IdStr":"769403753574932480","K":"key","Url":"url","V":"valueupdate","T":123}' --compressed