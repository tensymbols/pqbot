package app

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"vkbot/internal/db"
	"vkbot/internal/users"
)

const api = "https://api.vk.com/method/"
const apiVersion = "5.131"

type App interface {
	SendMessageToUser(id int64, msg string) error
	AddUser(id int64, name string, description string) error
	EditUser(id int64, name string, description string) error
	DeleteUser(id int64) error
}

type BotApp struct {
	accessKey string

	storage db.DB
	client  *http.Client
}

func (a *BotApp) SendMessageToUser(id int64, msg string) error {

	u, err := url.Parse(api + "messages.send")
	if err != nil {
		return err
	}
	v := url.Values{}
	strUserID := strconv.FormatInt(id, 10)
	//v.Add("peer_id", strUserID)
	v.Add("user_id", strUserID)
	strRandID := strconv.FormatInt(int64(rand.Int31()), 10)
	v.Add("random_id", strRandID)
	v.Add("message", msg)
	v.Add("access_token", a.accessKey)
	v.Add("v", "5.131")

	u.RawQuery = v.Encode()

	fmt.Println(u.String())

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	fmt.Println(err)
	_, err = a.client.Do(req)
	fmt.Println(err)

	return nil
}

func (a *BotApp) AddUser(id int64, name string, description string) error {
	return a.storage.Add(id, users.User{
		ID:          id,
		Name:        name,
		Description: description,
	})
}
func (a *BotApp) EditUser(id int64, name string, description string) error {
	if e, err := a.storage.Get(id); err != nil {
		return err
	} else {
		u := e.(users.User)
		a.storage.Add(id, users.User{
			ID:          u.ID,
			Name:        name,
			Description: description,
		})
	}
	return nil
}
func (a *BotApp) DeleteUser(id int64) error {
	if _, err := a.storage.Get(id); err != nil {
		return err
	}
	return a.storage.Delete(id)
}

func NewApp(accessKey string) App {
	return &BotApp{
		accessKey: accessKey,
		storage:   db.NewDB(),
		client:    http.DefaultClient,
	}
}
