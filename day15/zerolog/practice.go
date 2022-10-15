package main

import (
	"os"
	"errors"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
	_ "github.com/rs/zerolog/pkgerrors"
)


func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	appPid := os.Getpid()

	log.Print("Hello world")

	log.Info().Str("Note", "App running").Int("PID", appPid).Send()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("Prueba", "Hello world").Int("PID", appPid).Send()

	file, err := os.OpenFile("./test.log", os.O_RDWR, 0755)
	if err != nil && !os.IsExist(err){
		file, err = os.Create("./test.log")
	}
	defer file.Close()

	// send log to file output
	log.Logger = log.Output(file)

	err = errors.New("Test error log")
	log.Error().Err(err).Send()
	log.Info().Str("message", "Test after file log").Send()


	newLogger := zerolog.New(file)
	newLogger.Info().Str("Message", "Mensaje info").Int("PID", appPid).Send()
}