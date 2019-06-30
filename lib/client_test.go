package lib

import (
	"context"
	"fmt"
	"github.com/mookjp/yakiire/test"
	"testing"
)

var helper *test.Helper

func TestNewClient(t *testing.T) {
	setup()

	type args struct {
		ctx    context.Context
		config *ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fireStoreErr bool
	}{
		{
			name: "should create Client with the connection to Firestore",
			args: args{
				ctx: context.Background(),
				config: &ClientConfig{
					Credentials: "test",
					ProjectID: "yakiire",
				},
			},
			wantErr: false,
			fireStoreErr: false,
		},
		{
			name: "should create Client with the connection to Firestore without credentials",
			args: args{
				ctx: context.Background(),
				config: &ClientConfig{
					ProjectID: "yakiire",
				},
			},
			wantErr: false,
			fireStoreErr: false,
		},
		{
			name: "should return error if it can't connect to Firestore",
			args: args{
				ctx: context.Background(),
				config: &ClientConfig{
					Credentials: "test",
					ProjectID: "yakiire",
				},
			},
			wantErr: false,
			fireStoreErr: true,
		},
	}
	for _, tt := range tests {
		setup()

		if tt.fireStoreErr {
			helper.Close()
		}

		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.ctx, tt.args.config)
			if err != nil {
				if tt.wantErr != true {
					t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
					return
				} else {
					t.Logf("err is expected, ok: error: %+v", err)
				}
			}
		})

		teardown()
	}
}

func setup() {
	helper = test.NewHelper()
	helper.CreateData()
}

func teardown() {
	if err := helper.DeleteAll(); err != nil {
		if err.Error() == "rpc error: code = Canceled desc = grpc: the client connection is closing" {
			fmt.Printf("couldn't delete docs as the connection was closed intentionally, cause: %+v\n", err)
			return
		}
		panic(err)
	}
	err := helper.Close()
	if err != nil {
		panic(err)
	}
}
