
package std_logger

import (
	"log"

	"github.com/kmcsr/go-logger"
)

var Logger = initStdLogger()

func initStdLogger()(logger.Logger){
	return logger.WrapLogger(log.Default())
}

func Unwrap(l logger.BasicLogger)(*log.Logger){
	return logger.Unwrap(l).(*log.Logger)
}
