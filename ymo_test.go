package ymo

import (
	"testing"
)

func TestYMOClient_GetStatus(t *testing.T) {
	type fields struct {
		token      string
		counter    string
		clientType string
		debug      bool
	}
	type args struct {
		eventID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Test_GetStatus",
			fields{
				token:      "AQAAAAAABT",
				counter:    "01234567",
				clientType: "CLIENT_ID",
				debug:      true,
			},
			args{
				eventID: "1234567890",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := NewYMOClient(tt.fields.counter, tt.fields.token, tt.fields.clientType, tt.fields.debug)
			if result, err := g.GetStatus(tt.args.eventID); (err != nil) != tt.wantErr {
				t.Errorf("YMOClient.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				switch result {
				case `{"errors":[{"error_type":"invalid_token","message":"Invalid oauth_token"}],"code":403,"message":"Invalid oauth_token"}`:
					t.Log("Invalid oauth_token")
				case `{"errors":[{"error_type":"invalid_parameter","message":"No object with specified ID."}],"code":400,"message":"No object with specified ID."}`:
					t.Log("No object with specified ID.")
				default:
					t.Errorf("Result: %v", result)
				}
			}
		})
	}
}
