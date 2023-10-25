package utils

import log "github.com/cihub/seelog"

func InitLogger() {
	logger, err := log.LoggerFromConfigAsFile("./conf/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
}
