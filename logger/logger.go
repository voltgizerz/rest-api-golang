package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	logrus "github.com/sirupsen/logrus"
)

// Log - logrus logging.
var Log *logrus.Logger

// SetupLog - return logrus.
func init() {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	Log = log
}
