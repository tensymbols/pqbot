package main

import (
	"log"
	"vkbot/internal/app"
	"vkbot/internal/ports"
)

const groupAccessKey = "vk1.a.giHcHwo3bIYME861JvsvSSTZn5N63Bx_NAJasg9EYuvlYy089HmAtbpLC1q1extZxz4yoHd-uu1sYgyv10Qud5BMPk_2Kuzs9ml2wk0xCxmdQ17QoxoBJugfG3lWrlOWBCX4B0LvCIq6L8QQ32gH3j_Wmhs44eH5IXDolC8QREcvpApnB1jTIkl41kQ9-PXKLTuRWP-eLJuCXbd7Zahr0Q"
const externalAppAccessKey = "61166f6861166f6861166f6884620264346611661166f6805769b3eb221de80cb59bdb5"
const apiVersion = "5.131"

func main() {
	logger := log.Default()
	logger.Println("Starting app...")

	server := ports.NewServer(":8080", app.NewApp(externalAppAccessKey, groupAccessKey, apiVersion))
	server.Run()
}
