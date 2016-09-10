package main

import (
	"strings"
	"time"

	as "github.com/aerospike/aerospike-client-go"
)

var client *as.Client

//Init connects to the aerospike server
func Init(host string, port int) ([]string, error) {
	hosts := []*as.Host{
		as.NewHost(host, port),
	}
	cli, err := as.NewClientWithPolicyAndHost(nil, hosts...)
	if err != nil {
		return nil, err
	}

	client = cli
	node := client.GetNodes()[0]
	conn, _ := node.GetConnection(time.Second)
	infoMap, _ := as.RequestInfo(conn, "namespaces")
	nsString := infoMap["namespaces"]
	namespaces := strings.Split(nsString, ";")
	return namespaces, err
}

//GetRec returns a record for given key
func GetRec(ns, set, key string) (*as.Record, error) {
	spikeKey, _ := as.NewKey(ns, set, key)
	rec, err := client.Get(nil, spikeKey)
	return rec, err
}

//DeleteRec returns a record for given key
func DeleteRec(ns, set, key string) (bool, error) {
	spikeKey, _ := as.NewKey(ns, set, key)
	existed, err := client.Delete(nil, spikeKey)
	return existed, err
}

//PutRec returns a record for given key
func PutRec(ns, set, key string, record as.BinMap) error {
	spikeKey, _ := as.NewKey(ns, set, key)
	err := client.Put(nil, spikeKey, record)
	return err
}
