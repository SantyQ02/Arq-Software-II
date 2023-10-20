package solr

import (
	"fmt"

	solr "github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

var SolrClient *solr.Connection

func StartSolr() error {

	// Crea una conexión a Solr
	var err error
	SolrClient, err = solr.Init("localhost", 8983, "hotels")
	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
		return err
	} else {
		log.Info("Connection Established")
	}
	return nil
}

func Test() error {
	// Realiza la consulta a Solr
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"}, // Consulta que selecciona todos los documentos
		},
		Rows: 10, // Número de filas a recuperar (ajusta según tus necesidades)
	}
	resp, err := SolrClient.Select(query)
	if err != nil {
		return err
	}

	// Itera a través de los resultados e imprímelos
	for _, doc := range resp.Results.Collection {
		// Aquí puedes acceder a los campos del documento y mostrar la información que necesites
		// Por ejemplo, si tienes un campo "nombre" en tus documentos Solr, puedes imprimirlo así:
		fmt.Printf("Nombre: %s\n", doc.Fields["test"])
		// Repite esto para otros campos que desees imprimir
	}
	return nil
}

/*
Error de permisos contenedor de solr -> hacer solr_data publico
sudo chmod -R 777 ./solr_data
*/
