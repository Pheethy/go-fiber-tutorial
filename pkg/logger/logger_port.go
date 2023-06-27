package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pheethy/go-fiber-tutorial/pkg/utils"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResp(resp any)
}
/* entity */
type Logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func (l *Logger) Print() ILogger {
	utils.Debug(l)
	return l
}
func (l *Logger) Save() {
	var data = utils.OutPut(l)
	fileName := fmt.Sprintf("./assets/logs/logger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-10-15"), "-", ""))
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Opening file failed:%v", err)
	}
	defer file.Close()

	file.WriteString(string(data) + "\n")
}
func (l *Logger) SetQuery(c *fiber.Ctx) {}
func (l *Logger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(body); err != nil {
		log.Printf("body parser error: %v", err)
	}
}
func (l *Logger) SetResp(resp any) {
	l.Response = resp
}

func InitLogger(c *fiber.Ctx) ILogger {
	log := &Logger{
		Time: time.Now().Local().Format("2006-10-02 18:00:00"),
		Ip: c.IP(),
		Method: c.Method(),
		StatusCode: c.Response().StatusCode(),
	}

	log.SetQuery(c)
	log.SetBody(c)
	log.SetResp(c)

	return log
}
