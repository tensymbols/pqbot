package main

import (
	"log"
	"vkbot/internal/app"
	"vkbot/internal/ports"
)

const accessKey = "vk1.a.giHcHwo3bIYME861JvsvSSTZn5N63Bx_NAJasg9EYuvlYy089HmAtbpLC1q1extZxz4yoHd-uu1sYgyv10Qud5BMPk_2Kuzs9ml2wk0xCxmdQ17QoxoBJugfG3lWrlOWBCX4B0LvCIq6L8QQ32gH3j_Wmhs44eH5IXDolC8QREcvpApnB1jTIkl41kQ9-PXKLTuRWP-eLJuCXbd7Zahr0Q"

func main() {
	logger := log.Default()
	logger.Println("Starting app...")

	server := ports.NewServer(":8080", app.NewApp(accessKey))
	server.Run()
}
