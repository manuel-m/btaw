package main

import (
	"btaw"
	"btaw/cmd/filestore/app"
	"flag"
	"fmt"
)

/*
- day lists to proceed
- call api endpoint
- convert normalized response
- filename: data/ts/symbol-tf.json
- mkdir /target/dir
- save file
*/

func main() {
	// cli parse
	flag.StringVar(&app.Env, "env", btaw.AppEnvDefault, "Environment (development|staging|production)")
	flag.StringVar(&app.GwUrl, "gw", btaw.GwUrlDefault, "Gateway url (http://localhost:4000/)")
	flag.Parse()

	fmt.Println(app.GwUrl)

}
