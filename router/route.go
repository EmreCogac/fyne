package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micmonay/keybd_event"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/data/left", left)
	r.GET("/data/right", right)
	r.GET("/test", test)

	return r

}
func test(c *gin.Context) {

	c.JSON(200, gin.H{
		"test": "test başarılı",
	})

}

func left(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": "left",
	})
	automate("0")
}
func right(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": "right",
	})
	automate("1")
}

func automate(direction string) error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	defer kb.Clear()

	if direction == "1" {
		kb.SetKeys(keybd_event.VK_RIGHT)
	} else if direction == "0" {
		kb.SetKeys(keybd_event.VK_LEFT)
	}

	err = kb.Launching()
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)
	return nil

}
