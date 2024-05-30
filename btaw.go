package btaw

import "fmt"

const Version string = "0.0.1"

var AppEnvDefault string = "development"

const GwPortDefault = 4000
const GwHostDefault = "localhost"

var GwURLDefault = "http://" + GwHostDefault + ":" + fmt.Sprintf("%d", GwPortDefault)
