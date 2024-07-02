package mandrill

import (
	"context"
	"os"
	"testing"
)

func TestMessageApi_Send(t *testing.T) {
	resp, err := New(os.Getenv("apiKey")).MessageApi.Send(context.Background(), &SendMessageReq{
		Message: SendMessageItem{
			Text:      "testapi",
			Subject:   "testapi",
			FromEmail: os.Getenv("fromEmail"),
			FromName:  os.Getenv("fromName"),
			To: []SendMessageToItem{
				{
					Email: os.Getenv("toEmail"),
					Name:  os.Getenv("toName"),
					Type:  "to",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", resp)
}
