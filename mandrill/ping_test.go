package mandrill

import (
	"context"
	"os"
	"testing"
)

func TestPingApi_Ping(t *testing.T) {
	if err := New(os.Getenv("apiKey")).PingApi.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}
}
