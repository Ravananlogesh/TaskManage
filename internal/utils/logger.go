package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Context string

const RequestIDKey Context = "reqid"

var (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type Logger struct {
	Sid string
	Ref string
}

func (l *Logger) SetSid(r *http.Request) {
	requestID, ok := r.Context().Value(RequestIDKey).(string)
	if !ok {
		KeyValue := ""
		session := uuid.NewV4()
		sessionSHA256 := session.String()
		KeyValue = strings.ReplaceAll(sessionSHA256, "-", "")
		l.Sid = KeyValue
	} else {
		l.Sid = requestID
	}
}

func (l *Logger) SetRef(ref any) {
	l.Ref = fmt.Sprintf("%v", ref)
}
func (l *Logger) RemoveRef() {
	l.Ref = ""
}

func (l *Logger) Log(level string, message ...any) {

	log.Printf("%s - (%s) - (%s) - %v\n", l.Sid, l.Ref, level, message)
}
