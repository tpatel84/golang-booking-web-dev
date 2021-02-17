package renders

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/tpatel84/golang-booking-web-dev/internal/config"
	"github.com/tpatel84/golang-booking-web-dev/internal/models"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// Store data in session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	//Setup an logger for application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWriter) WriteHeader(statusCode int) {

}

func (mw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}