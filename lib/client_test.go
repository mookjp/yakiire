package lib

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/firestore"

	"github.com/mookjp/yakiire/test"
)

var helper *test.Helper

func TestNewClient(t *testing.T) {
	setup()

	type args struct {
		ctx    context.Context
		config *ClientConfig
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		fireStoreErr bool
	}{
		{
			name: "should create Client with the connection to Firestore",
			args: args{
				ctx: context.Background(),
				config: &ClientConfig{
					Credentials: "test",
					ProjectID:   "yakiire",
				},
			},
			wantErr:      false,
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
			wantErr:      false,
			fireStoreErr: false,
		},
		{
			name: "should return error if it can't connect to Firestore",
			args: args{
				ctx: context.Background(),
				config: &ClientConfig{
					Credentials: "test",
					ProjectID:   "yakiire",
				},
			},
			wantErr:      false,
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
				}
				t.Logf("err is expected, ok: error: %+v", err)
			}
		})

		teardown()
	}
}

func TestClient_Get(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "yakiire")
	if err != nil {
		panic(err)
	}

	type fields struct {
		config    *ClientConfig
		firestore *firestore.Client
	}
	type args struct {
		ctx        context.Context
		collection string
		docID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "returns a doc when the ID is matched",
			fields: fields{
				config: &ClientConfig{
					Credentials: "test",
					ProjectID:   "yakiire",
				},
				firestore: client,
			},
			args: args{
				ctx:        context.Background(),
				collection: "products",
				docID:      "1",
			},
			want:    "{\"Attributes\":{\"color\":\"red\",\"size\":\"100\"},\"CategoryIDs\":[\"1\",\"2\",\"3\"],\"ID\":\"1\",\"Name\":\"Test Product\"}",
			wantErr: false,
		},
		{
			name: "returns an error when the ID is not matched",
			fields: fields{
				config: &ClientConfig{
					Credentials: "test",
					ProjectID:   "yakiire",
				},
				firestore: client,
			},
			args: args{
				ctx:        context.Background(),
				collection: "products",
				docID:      "XXX",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:    tt.fields.config,
				firestore: tt.fields.firestore,
			}
			got, err := c.Get(tt.args.ctx, tt.args.collection, tt.args.docID)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			if got.String() != tt.want {
				t.Errorf("Client.Get().String() = %s, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestClient_Query(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "yakiire")
	if err != nil {
		panic(err)
	}

	type fields struct {
		config    *ClientConfig
		firestore *firestore.Client
	}
	type args struct {
		ctx        context.Context
		collection string
		conditions []*Condition
		limit      int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "returns a doc when the Name is matched",
			fields: fields{
				config: &ClientConfig{
					Credentials: "test",
					ProjectID:   "yakiire",
				},
				firestore: client,
			},
			args: args{
				ctx:        context.Background(),
				collection: "products",
				conditions: []*Condition{
					{
						Path:  "Name",
						Op:    "==",
						Value: "Test Product",
					},
				},
				limit: 1,
			},
			want: []string{
				"{\"Attributes\":{\"color\":\"red\",\"size\":\"100\"},\"CategoryIDs\":[\"1\",\"2\",\"3\"],\"ID\":\"1\",\"Name\":\"Test Product\"}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:    tt.fields.config,
				firestore: tt.fields.firestore,
			}
			got, err := c.Query(tt.args.ctx, tt.args.collection, tt.args.conditions, tt.args.limit)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			for i, s := range got {
				if s.String() != tt.want[i] {
					t.Errorf("Client.Query() doc = %s, want %v", s.String(), tt.want[i])
				}
			}
		})
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
