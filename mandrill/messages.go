package mandrill

import "context"

const (
	sendMsgPath = "/messages/send"
	msgInfoPath = "/messages/info"
)

type SendMessageToItem struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type"` //the header type to use for the recipient, defaults to "to" if not provided Possible values: "to", "cc", or "bcc".
}

type SendMessageItem struct {
	Text        string              `json:"text,omitempty"`
	HTML        string              `json:"html,omitempty"`
	Subject     string              `json:"subject"`
	FromEmail   string              `json:"from_email"`
	FromName    string              `json:"from_name,omitempty"`
	To          []SendMessageToItem `json:"to"`
	TrackOpens  bool                `json:"track_opens"`  // 是否跟踪邮件打开
	TrackClicks bool                `json:"track_clicks"` // 是否跟踪邮件点击
	Headers     map[string]string   `json:"headers"`      //
	Tags        []string            `json:"tags"`
}

type SendMessageReq struct {
	Key     string          `json:"key"`
	Message SendMessageItem `json:"message"`
	Async   bool            `json:"async,omitempty"`
	IpPool  string          `json:"ip_pool,omitempty"`
	SendAt  string          `json:"send_at,omitempty"`
}

type SendMessageRespItem struct {
	Email        string `json:"email"`
	Status       string `json:"status"`
	RejectReason string `json:"reject_reason"`
	QueuedReason string `json:"queued_reason"`
	Id           string `json:"_id"`
}

type SendMessageResp []SendMessageRespItem

type MessageApi struct {
	api *API
}

func newMessageApi(api *API) *MessageApi {
	return &MessageApi{api: api}
}

func (api *MessageApi) Send(ctx context.Context, req *SendMessageReq) (resp SendMessageResp, err error) {
	req.Key = api.api.Key

	err = api.api.Post(ctx, sendMsgPath, req, &resp)

	return resp, err
}

type MessageInfoReq struct {
	Key string `json:"key"`
	Id  string `json:"id"`
}

type MessageInfoOpensDetail struct {
	Ts       int    `json:"ts"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
	Ua       string `json:"ua"`
}

type MessageInfoClicksDetail struct {
	MessageInfoOpensDetail
	Url string `json:"url"`
}

type MessageInfoSmtpEvent struct {
	Ts   int    `json:"ts"`
	Type string `json:"type"`
	Diag string `json:"diag"`
}

type MessageInfoResp struct {
	Ts           int                       `json:"ts"`
	Id           string                    `json:"_id"`
	Sender       string                    `json:"sender"`
	Template     string                    `json:"template"`
	Subject      string                    `json:"subject"`
	Email        string                    `json:"email"`
	Tags         []string                  `json:"tags"`
	Opens        int                       `json:"opens"`
	OpensDetail  []MessageInfoOpensDetail  `json:"opens_detail"`
	Clicks       int                       `json:"clicks"`
	ClicksDetail []MessageInfoClicksDetail `json:"clicks_detail"`
	State        string                    `json:"state"`
	Metadata     any                       `json:"metadata"`
	SmtpEvents   []MessageInfoSmtpEvent    `json:"smtp_events"`
}

func (api *MessageApi) Info(ctx context.Context, req *MessageInfoReq) (resp *MessageInfoResp, err error) {
	req.Key = api.api.Key

	err = api.api.Post(ctx, msgInfoPath, req, &resp)

	return resp, err
}
