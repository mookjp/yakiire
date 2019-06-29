package lib

import (
	"context"
	"github.com/mookjp/yakiire/test"
	"reflect"
	"testing"

	"cloud.google.com/go/firestore"
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
		want    *Client
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		setup()

		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.ctx, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})

		teardown()
	}
}

func TestClient_Get(t *testing.T) {

	type fields struct {
		config    *ClientConfig
		firestore *firestore.Client
		ctx       context.Context
	}
	type args struct {
		collection string
		docID      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Doc
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		setup()

		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:    tt.fields.config,
				firestore: tt.fields.firestore,
				ctx:       tt.fields.ctx,
			}
			got, err := c.Get(tt.args.collection, tt.args.docID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Get() = %v, want %v", got, tt.want)
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
	helper.DeleteAll()

}
