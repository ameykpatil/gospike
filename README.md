# gospike
A simple ui-client to get, put, delete records from aerospike database.

#### How to start :
 - Assuming you already have GO enviroment set i.e. `$GOROOT`, `$GOPATH` etc. 
 - Get the **gospike** - `go get github.com/ameykpatil/gospike`
 - Go to the directory - `cd $GOPATH/src/github.com/ameykpatil/gospike`
 - Install the dependecies - `go get ./...` (_gin_, _aerospike-client-go_ etc.)
 - Build **gospike** - `go install`
 - Run **gospike** - `gospike` (assuming you have set PATH properly `export PATH=$PATH:$GOPATH/bin`)

#### How to use : 
- Once **gospike** server is up, go to `http://localhost:4848`
- **gospike** will first ask to connect to `aerospike` server. Enter `host` & `port` on which your `aerospike` server is up.
- On successful connection, you will be directed to a page where you can carry out simple operations to `aerospike` database.
- For now, basic `GET`, `PUT`, `DELETE` is supported.
- For every operation you will need to enter `namespace`, `set` & `key`.
- For `namespaces` available values from `aerospike` cluster will be pre-populated.
