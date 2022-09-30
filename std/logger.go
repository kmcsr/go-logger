
package std_logger

import (
	"os"
	"log"

	"github.com/kmcsr/go-logger"
)

var Logger = initStdLogger()

func initStdLogger()(logger.Logger){
	return logger.WrapLogger(log.Default())
}

func New()(logger.Logger){
	return logger.WrapLogger(log.New(os.Stderr, "", log.LstdFlags))
}

func Unwrap(l logger.BasicLogger)(*log.Logger){
	return logger.Unwrap(l).(*log.Logger)
}
