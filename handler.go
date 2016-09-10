package main

import (
	"log"
	"net/http"
	"strconv"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
)

const (
	defaultNS  = "test"
	defaultSet = "test_set"
)

//SpikeObject would have a body of POST APIs
type SpikeObject struct {
	Namespace string    `form:"namespace" json:"namespace"`
	Set       string    `form:"set" json:"set"`
	Key       string    `form:"key" json:"key" binding:"required"`
	Record    as.BinMap `form:"record" json:"record" binding:"required"`
}

//Connect creates a connection to aerospike
func Connect(c *gin.Context) {
	host := c.DefaultQuery("host", "127.0.0.1")
	portString := c.Query("port")
	port, _ := strconv.Atoi(portString)
	namespaces, err := Init(host, port)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "index.tmpl", gin.H{
			"message": err.Error(),
		})
	} else {
		log.Printf("host: %v port: %v", host, port)
		c.HTML(http.StatusOK, "operation.tmpl", gin.H{
			"message":    "Connected",
			"namespaces": namespaces,
		})
	}
}

//GetRecord gets a record from aerospike
func GetRecord(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", defaultNS)
	set := c.DefaultQuery("set", defaultSet)
	key := c.Param("key")

	record, err := GetRec(namespace, set, key)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else if record == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "record not found",
		})
	} else {
		bins := ConvertRecord(record.Bins)
		c.JSON(http.StatusOK, gin.H{
			"key":    key,
			"record": bins,
		})
	}
}

//DeleteRecord deletes a record from aerospike
func DeleteRecord(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", defaultNS)
	set := c.DefaultQuery("set", defaultSet)
	key := c.Param("key")

	existed, err := DeleteRec(namespace, set, key)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else if existed == true {
		c.JSON(http.StatusOK, gin.H{
			"message": "record deleted",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "record didn't exist",
		})
	}
}

//AddRecord adds a record to aerospike
func AddRecord(c *gin.Context) {
	var form SpikeObject
	bindErr := c.Bind(&form)
	if bindErr != nil {
		log.Println("bind error", bindErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": bindErr.Error(),
		})
	} else {
		log.Printf("namespace: %v set: %v", form.Namespace, form.Set)
		if form.Namespace == "" {
			form.Namespace = defaultNS
		}
		if form.Set == "" {
			form.Set = defaultSet
		}
		err := PutRec(form.Namespace, form.Set, form.Key, form.Record)
		if err != nil {
			log.Println("error while adding", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			log.Printf("key: %v record: %v", form.Key, form.Record)
			c.JSON(http.StatusOK, gin.H{
				"key":    form.Key,
				"record": form.Record,
			})
		}
	}
}
