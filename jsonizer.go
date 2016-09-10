package main

import (
	"log"

	as "github.com/aerospike/aerospike-client-go"
)

//ConvertRecord converts record to proper format
func ConvertRecord(binMap as.BinMap) map[string]interface{} {
	bins := make(map[string]interface{})
	for key, value := range binMap {
		switch value := value.(type) {
		case map[string]interface{}:
			log.Println("converting stringMap ", key)
			val := ConvertStringMap(value)
			bins[key] = val
		case map[interface{}]interface{}:
			log.Println("converting interfaceMap ", key)
			val := ConvertInterfaceMap(value)
			bins[key] = val
		case []interface{}:
			log.Println("converting interfaceArray ", key)
			val := ConvertInterfaceArray(value)
			bins[key] = val
		default:
			log.Println("converting normal ", key)
			bins[key] = value
		}
	}
	return bins
}

//ConvertInterfaceMap converts interfaceMap to stringMap
func ConvertInterfaceMap(interfaceMap map[interface{}]interface{}) map[string]interface{} {
	stringMap := make(map[string]interface{})
	for key, value := range interfaceMap {
		switch key := key.(type) {
		case string:
			switch value := value.(type) {
			case map[string]interface{}:
				val := ConvertStringMap(value)
				stringMap[key] = val
			case map[interface{}]interface{}:
				val := ConvertInterfaceMap(value)
				stringMap[key] = val
			case []interface{}:
				val := ConvertInterfaceArray(value)
				stringMap[key] = val
			default:
				stringMap[key] = value
			}
			// if str, ok := value.(map[interface{}]interface{}); ok {
			//  val := ConvertMap(value)
			//  stringMap[key] = val
			// } else {
			//  stringMap[key] = value
			// }
		}
	}
	return stringMap
}

//ConvertStringMap converts stringMap to proper format
func ConvertStringMap(sMap map[string]interface{}) map[string]interface{} {
	stringMap := make(map[string]interface{})
	for key, value := range sMap {
		switch value := value.(type) {
		case map[string]interface{}:
			val := ConvertStringMap(value)
			stringMap[key] = val
		case map[interface{}]interface{}:
			val := ConvertInterfaceMap(value)
			stringMap[key] = val
		case []interface{}:
			val := ConvertInterfaceArray(value)
			stringMap[key] = val
		default:
			stringMap[key] = value
		}
		// if str, ok := value.(map[interface{}]interface{}); ok {
		//  val := ConvertMap(value)
		//  stringMap[key] = val
		// } else {
		//  stringMap[key] = value
		// }
	}
	return stringMap
}

//ConvertInterfaceArray converts interface array to proper format
func ConvertInterfaceArray(array []interface{}) []interface{} {
	newArray := make([]interface{}, len(array))
	for i, val := range array {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			stringMap := ConvertStringMap(concreteVal)
			newArray[i] = stringMap
		case map[interface{}]interface{}:
			interfaceMap := ConvertInterfaceMap(concreteVal)
			newArray[i] = interfaceMap
		case []interface{}:
			interfaceArray := ConvertInterfaceArray(concreteVal)
			newArray[i] = interfaceArray
		default:
			newArray[i] = concreteVal

		}
	}
	return newArray
}
