package ports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"vkbot/internal/app"
	vkjson "vkbot/internal/presenters"
)

func hello(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteString("Hello!")
		c.Writer.Flush()
	}
}

func getFriends(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		friends, err := a.GetUserFriends(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.Writer.WriteString(fmt.Sprint(friends))
		c.Writer.Flush()
		fmt.Println(err)
	}
}

func eventHandler(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		var eventBody vkjson.MessageNewEvent
		err := c.BindJSON(&eventBody)

		if err != nil {
			return
		}
		fmt.Println(eventBody.Type)
		switch eventBody.Type {
		case "message_new":
			text := eventBody.Object.Message.Text
			fromID := eventBody.Object.Message.FromId
			switch {
			case strings.HasPrefix(text, "/chance"):
				text = strings.TrimLeft(text, "/chance")
				if text == "" {
					err = a.SendMessageToUser(fromID, "чтобы использовать chance, введи /chance [событие]")
					break
				}
				p := strconv.Itoa(rand.Intn(101))
				err = a.SendMessageToUser(fromID, "👁 Вероятность того, что "+text+" — "+p+" %")
			case strings.HasPrefix(text, "/true"):
				text = strings.TrimLeft(text, "/true")
				if text == "" {
					err = a.SendMessageToUser(fromID, "чтобы использовать true, введи /true [событие]")
					break
				}
				p := rand.Intn(2)
				if p == 1 {
					err = a.SendMessageToUser(fromID, "🧢 Информация о том, что "+text+" — ложь")
				} else {
					err = a.SendMessageToUser(fromID, "💯 Информация о том, что "+text+" — правда")
				}

			case strings.HasPrefix(text, "/help"):
				err = a.SendMessageToUser(fromID,
					"Доступные команды: \n"+
						"/chance [событие] - вероятность события\n"+
						"/true [событие] - правда или неправда\n"+
						"/dice - бросить кубик")
			case strings.HasPrefix(text, "/dice"):
				roll := strconv.Itoa(rand.Intn(6) + 1)
				err = a.SendMessageToUser(fromID, "🎲 Вам выпало "+roll+" очков")
			default:
				err = a.SendMessageToUser(fromID, []string{
					"Я тебя не понимаю, используй /help, чтобы посмотреть список команд",
					"Каво",
					"Ай донт андестенд",
				}[rand.Intn(3)])
			}

			if err != nil {
				c.Status(http.StatusInternalServerError)

			} else {
				c.Status(http.StatusOK)
			}
			c.Writer.WriteString("ok")
			c.Writer.Flush()
		case "confirmation":
			c.Writer.Write([]byte("c09aed4d"))
			c.Status(http.StatusOK)
		}

	}

}

/*func getRandomPrikol() string {

	opts := []string{
		"НИКТО НЕ СМЕЕТ МНЕ ПРИКАЗЫВАТЬ",
		"НУ ЧТО ЕЩЕ",
		"ОТЛЕТАЕШЬ ОЧЕРЕДНЯРА",
		"ДА ЭТО ЖЕСТКО",
		"ЧТО",
		"РЕБЯТА НЕ СТОИТ ВСКРЫВАТЬ ЭТУ ТЕМУ",
		"ДА-ДА НЕ УДИВЛЯЙТЕСЬ",
		"ТЫ КТО ТАКОЙ ЧТОБЫ ЭТО ДЕЛАТЬ",
	}
	return opts[rand.Intn(len(opts))]
}*/
