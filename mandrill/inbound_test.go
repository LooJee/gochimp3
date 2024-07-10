package mandrill

import (
	"context"
	"os"
	"testing"
)

func Test_inboundApi_Domains(t *testing.T) {
	result, err := New(os.Getenv("apiKey")).InboundApi.Domains(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}
