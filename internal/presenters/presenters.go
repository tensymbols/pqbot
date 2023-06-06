package presenters

type Message struct {
	Date                  int64         `json:"date"`
	FromId                int64         `json:"from_id"`
	Id                    int64         `json:"id"`
	Out                   int64         `json:"out"`
	AdminAuthorId         int64         `json:"admin_author_id"`
	Attachments           []interface{} `json:"attachments"`
	ConversationMessageId int           `json:"conversation_message_id"`
	FwdMessages           []interface{} `json:"fwd_messages"`
	Important             bool          `json:"important"`
	IsHidden              bool          `json:"is_hidden"`
	PeerId                int64         `json:"peer_id"`
	RandomId              int64         `json:"random_id"`
	Text                  string        `json:"text"`
}

type MessageNewEvent struct {
	GroupId int64  `json:"group_id"`
	Type    string `json:"type"`
	EventId string `json:"event_id"`
	V       string `json:"v"`
	Object  struct {
		Message Message `json:"Message"`
	} `json:"object"`
	Secret string `json:"secret"`
}
type FriendsResponse struct {
	Response struct {
		Count int64   `json:"count"`
		Items []int64 `json:"items"`
	} `json:"response"`
}
