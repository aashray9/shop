package middleware

import (
	"bytes"
	"io/ioutil"
	"lms/common"

	"github.com/gin-gonic/gin"
)

var log = common.Loggers()

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ResponseLogger(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	log.Info().
		Int("statusCode", c.Writer.Status()).
		Msg(blw.body.String())
}

func RequestLogger(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Info().
		Str("method", c.Request.Method).
		Str("url", c.Request.RequestURI).
		Str("user_agent", c.Request.UserAgent()).
		Msg(string(jsonData))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	c.Next()
}
