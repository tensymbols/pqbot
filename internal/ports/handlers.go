package ports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"vkbot/internal/app"
)

type message struct {
	Date                  int           `json:"date"`
	FromId                int           `json:"from_id"`
	Id                    int           `json:"id"`
	Out                   int           `json:"out"`
	AdminAuthorId         int           `json:"admin_author_id"`
	Attachments           []interface{} `json:"attachments"`
	ConversationMessageId int           `json:"conversation_message_id"`
	FwdMessages           []interface{} `json:"fwd_messages"`
	Important             bool          `json:"important"`
	IsHidden              bool          `json:"is_hidden"`
	PeerId                int           `json:"peer_id"`
	RandomId              int           `json:"random_id"`
	Text                  string        `json:"text"`
}

type messageNewEvent struct {
	GroupId int    `json:"group_id"`
	Type    string `json:"type"`
	EventId string `json:"event_id"`
	V       string `json:"v"`
	Object  struct {
		Message message `json:"message"`
	} `json:"object"`
	Secret string `json:"secret"`
}

func hello(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteString("Hello!")
		c.Writer.Flush()
	}
}
func eventHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		var eventBody messageNewEvent
		err := c.BindJSON(&eventBody)

		if err != nil {
			return
		}
		fmt.Println(eventBody.Type)
		switch eventBody.Type {
		case "message_new":
			fmt.Println(eventBody)
			fromID := eventBody.Object.Message.FromId
			fmt.Println(fromID)
			err = app.SendMessageToUser(int64(fromID), getRandomPrikol())
			if err != nil {
				c.Status(http.StatusInternalServerError)

			} else {
				c.Status(http.StatusOK)
			}
			c.Writer.WriteString("ok")
			c.Writer.Flush()
		case "confirmation":
			c.Writer.Write([]byte("abe03ede"))
			c.Status(http.StatusOK)
		}

	}

}

func getRandomPrikol() string {
	opt := rand.Int31n(10)
	switch opt {
	case 0:
		return "НИКТО НЕ СМЕЕТ МНЕ ПРИКАЗЫВАТЬ"
	case 1:
		return "НУ ЧТО ЕЩЕ"
	case 2:
		return "ОТЛЕТАЕШЬ ОЧЕРЕДНЯРА"
	case 3:
		return "ДА ЭТО ЖЕСТКО"
	case 4:
		return "ЧТО"
	case 5:
		return "РЕБЯТА НЕ СТОИТ ВСКРЫВАТЬ ЭТУ ТЕМУ"
	case 6:
		return "ДА-ДА НЕ УДИВЛЯЙТЕСЬ"
	default:
		return "ТЫ КТО ТАКОЙ ЧТОБЫ ЭТО ДЕЛАТЬ"
	}
}
