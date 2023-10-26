package solr

import (
	"os"

	solr "github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

var SolrClient *solr.Connection

func StartSolr() error {

	// Crea una conexión a Solr
	var err error
	SolrClient, err = solr.Init(os.Getenv("SOLR_SERVICE_URL") , 8983, "hotels")
	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
		return err
	} else {
		log.Info("Connection Established")
	}
	return nil
}

func StartTestSolr() error {

	// Crea una conexión a Solr
	var err error
	SolrClient, err = solr.Init(os.Getenv("SOLR_SERVICE_URL") , 8983, "hotels_test")
	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
		return err
	} else {
		log.Info("Connection Established")
	}
	return nil
}

