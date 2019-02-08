package main

import (
	"github.com/jacky-htg/shonet-frontend/libraries"
	"log"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) <= 0 {
		log.Println("please insert valid argument or type help example : go run elastic.go help")
		return
	}

	if arguments[0] == "help" {
		help := "\n This is for input database on SQL to elasticsearch database\n"+
		        " 1.\tFor insert bulk argument you can include name:<table_name> example:go run elastic.go bulking name:articles\n"+
				" 2.\tFor migration data mapping to ElasticSearch you can include <prefix_name>:<file_name> example: go run elastic.go mapping articles:articles_mapping\n"+
				" 3.\tFor delete elastic you just include delete <your_table_name> example: go run elastic.go delete articles\n"+
				" 4.\tFor check if your elastic is exists or not you can include exists <table_name> example: go run elastic.go exists articles\n"+
				"\n NOTE : please create mapping file on config folder and just make for json file \n\n"+
			    " Just for this time you just inputing 3 tables: 'articles', 'products`, and 'users'\n"+
		        " ..."
		log.Printf(help)
		return
	}

	if arguments[0] == "mapping" {
		log.Println("\nPut your request data ...")
		resp, err := libraries.InsertElasticMapping(arguments[1])
		if resp {log.Println("\nInserting mapping is success ...")} else {log.Println("\nERROR : <inserting mapper> : ", err.Error())}
		return
	}

	if arguments[0] == "delete" {
		log.Println("\nDelete your request data ...")
		resp := libraries.DeleteElastic(arguments[1])
		if resp {log.Println("\nDeleting elastic is success ...")} else {log.Println("\n ERROR ON OCUURED, please check your error")}
		return
	}

	if arguments[0] == "bulking" {
		log.Println("\nBulking your request data ...")
		var tables map[string]string
		tables = map[string]string{}

		instructions := strings.Split(arguments[1], ":")
		if 1 >= len(instructions) {
			log.Println("\n please insert valid argumen <name>:<value> example => table:tablename for more type : help")
			return
		}
		tables[instructions[0]] = instructions[1]

		log.Printf("\n bulking table : %s..\n please wait..\n\n", tables["name"])
		status, err := libraries.BulkingAllDataFromSQL(tables)
		if !status {
			if err != nil {
				panic(err)
				return
			}
		}

		log.Printf("Success to insert SQL data %s to elastic search..", tables["name"])
		return
	}

	if arguments[0] == "exists" {
		log.Println("\nChecking exsisting your request data ...")
		stat := libraries.CheckExistElastic(arguments[1])
		if stat {log.Println("\n"+"Finish..")} else {log.Println("\n"+"ERROR <exists> : something wrong")}

		return
	}
}