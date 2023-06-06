package app

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"vkbot/internal/db"
	vkjson "vkbot/internal/presenters"
	"vkbot/internal/users"
)

const api = "https://api.vk.com/method/"

type App interface {
	GetUserFriends(id int64) ([]int64, error)
	SendMessageToUser(id int64, msg string) error
	AddUser(id int64, name string, description string) error
	EditUser(id int64, name string, description string) error
	DeleteUser(id int64) error
}

type BotApp struct {
	extAppKey  string
	groupKey   string
	apiVersion string

	storage db.DB
	client  Client
}

func (a *BotApp) AddAppTokenAndVersion(values *url.Values) *url.Values {
	values.Add("access_token", a.extAppKey)
	values.Add("v", a.apiVersion)
	return values
}

func (a *BotApp) AddGroupTokenAndVersion(values *url.Values) *url.Values {
	values.Add("access_token", a.groupKey)
	values.Add("v", a.apiVersion)
	return values
}

func (a *BotApp) SendMessageToUser(id int64, msg string) error {

	u, _ := url.Parse(api + "messages.send")

	v := url.Values{}
	strUserID := strconv.FormatInt(id, 10)

	v.Add("user_id", strUserID)
	strRandID := strconv.FormatInt(int64(rand.Int31()), 10)
	v.Add("random_id", strRandID)
	v.Add("message", msg)
	a.AddGroupTokenAndVersion(&v)

	u.RawQuery = v.Encode()
	_, err := a.client.GetResponse(http.MethodPost, u, nil)
	return err
}
func (a *BotApp) GetUserFriends(id int64) ([]int64, error) {

	u, _ := url.Parse(api + "friends.get")
	v := url.Values{}
	strUserID := strconv.FormatInt(id, 10)
	strCount := strconv.FormatInt(10000, 10)

	v.Add("user_id", strUserID)
	v.Add("count", strCount)
	a.AddAppTokenAndVersion(&v)

	u.RawQuery = v.Encode()
	rb, err := a.client.GetResponseBytes(http.MethodPost, u, nil)

	if err != nil {
		return nil, err
	}
	var friends vkjson.FriendsResponse

	err = json.Unmarshal(rb, &friends)

	return friends.Response.Items, err

}

// doesnt work

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

func NewApp(eaKey string, gKey string, apiVersion string) App {
	return &BotApp{
		extAppKey:  eaKey,
		groupKey:   gKey,
		apiVersion: apiVersion,
		storage:    db.NewDB(),
		client: Client{
			HTTPClient: http.DefaultClient,
		},
	}
}
