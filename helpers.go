package pantsuAPI

import (
	"log"
)

// shorthand wrapper function for non fatal errors
func logError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

// shorthand wrapper function for fatal errors
func logErrorFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
