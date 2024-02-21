package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"one-day-job/config"
	"os"
	"path"
	"runtime"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

//func GetLogger() *Logger {
//	return &Logger{e}
//}

//func init() {
//	l := logrus.New()
//	l.SetReportCaller(true)
//	l.Formatter = &logrus.TextFormatter{
//		CallerPrettyfier: func(frame *runtime.Frame) (funcion string, file string) {
//			filename := path.Base(frame.File)
//			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
//		},
//		DisableColors: false,
//		FullTimestamp: true,
//	}
//
//	err := os.Mkdir("./logs", 0777)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	allFile, err := os.OpenFile("./logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
//	if err != nil {
//		log.Fatal(err)
//	}
//	l.SetOutput(io.Discard)
//
//	l.AddHook(&writeHook{
//		Writer:    []io.Writer{allFile, os.Stdout},
//		LogLevels: logrus.AllLevels,
//	})
//
//	l.SetLevel(logrus.TraceLevel)
//
//	e = logrus.NewEntry(l)
//}

// instead of using Get logger we could do this:
// in config.yml we could save separate field for Logger setup, which could have two subfields: writeToFile: False -> format: text
// then we could create a struct for logger also with two field, and after we could add it to our Config model and we could give these subfield to InitLogger function, which could have look like this
func InitLogger(cfg *config.Logger) *Logger {
	l := logrus.New()
	l.SetReportCaller(true)

	if cfg.Format == "json" {
		l.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			},
		})

	} else if cfg.Format == "text" {
		l.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				filename := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
			},
			DisableColors: false,
			FullTimestamp: true,
		}
	}

	if !cfg.WriteToFile {
		l.SetOutput(io.Discard)

		l.AddHook(&writeHook{
			Writer:    []io.Writer{os.Stdout},
			LogLevels: logrus.AllLevels,
		})

		l.SetLevel(logrus.TraceLevel)

		e = logrus.NewEntry(l)

		return &Logger{Entry: e}
	} else if cfg.WriteToFile {
		err := os.MkdirAll("./logs", 0777)
		if err != nil {
			log.Fatal(err)
		}

		allFile, err := os.OpenFile("./logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}

		l.SetOutput(io.Discard)

		l.AddHook(&writeHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})

		l.SetLevel(logrus.TraceLevel)

		e = logrus.NewEntry(l)

		return &Logger{Entry: e}
	} else {
		log.Fatal(`введите правильный формат("text"/"json")`)
		return nil
	}
}
