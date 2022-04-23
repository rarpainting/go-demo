package document

import (
	"context"
	"log"
)

func specFunc(specArgs int32, anyT interface{}) {
	a := 3
	log.Println("specString", 1234, a, specArgs)
}

func specFunc2(appId int) {
	log.Println("specString", 1234)
}

func specFunc3(ctx context.Context) {
	log.Println("specString", 1234)
}
