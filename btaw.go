package btaw

import "fmt"

const Version string = "0.0.1"

var AppEnvDefault string = "development"

// gw
const GwPortDefault = 4000
const GwHostDefault = "localhost"

// TODO const
var GwUrlDefault = "http://" + GwHostDefault + ":" + fmt.Sprintf("%d", GwPortDefault)
