package lib

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"

	"github.com/mookjp/yakiire/test"
)

var helper *test.Helper

func TestNewClient(t *testing.T) {
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
			teardown()
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
	setup()

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
			want:    "{\"Attributes\":{\"color\":\"red\",\"size\":100},\"CategoryIDs\":[\"1\",\"2\",\"3\"],\"ID\":\"1\",\"Name\":\"Test Product\"}",
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

	teardown()
}

func TestClient_Query(t *testing.T) {
	setup()

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
				"{\"Attributes\":{\"color\":\"red\",\"size\":100},\"CategoryIDs\":[\"1\",\"2\",\"3\"],\"ID\":\"1\",\"Name\":\"Test Product\"}",
			},
			wantErr: false,
		},
		{
			name: "returns all docs when the CategoryIDs are matched",
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
						Path:  "CategoryIDs",
						Op:    "array-contains",
						Value: "5",
					},
				},
			},
			want: []string{
				"{\"Attributes\":{\"color\":\"blue\",\"size\":200},\"CategoryIDs\":[\"3\",\"4\",\"5\"],\"ID\":\"2\",\"Name\":\"Another Test Product\"}",
				"{\"Attributes\":{\"color\":\"yellow\",\"size\":300},\"CategoryIDs\":[\"5\",\"6\",\"7\"],\"ID\":\"3\",\"Name\":\"Another Great Test Product\"}",
			},
			wantErr: false,
		},
		{
			name: "limits docs to 1 doc when the limit is 1",
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
						Path:  "CategoryIDs",
						Op:    "array-contains",
						Value: "5",
					},
				},
				limit: 1,
			},
			want: []string{
				"{\"Attributes\":{\"color\":\"blue\",\"size\":200},\"CategoryIDs\":[\"3\",\"4\",\"5\"],\"ID\":\"2\",\"Name\":\"Another Test Product\"}",
			},
			wantErr: false,
		},
		{
			name: "returns nothing when the Name is not matched",
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
						Value: "Not Matched",
					},
				},
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "returns error when the operation is not supported",
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
						Op:    "AAA",
						Value: "Not Matched",
					},
				},
			},
			want:    []string{},
			wantErr: true,
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
					t.Errorf("Client.Query() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			if len(tt.want) != len(got) {
				t.Errorf("Client.Query() error len(res) = %v, want %v", len(got), len(tt.want))
			}
			for i, s := range got {
				if s.String() != tt.want[i] {
					t.Errorf("Client.Query() doc = %s, want %v", s.String(), tt.want[i])
				}
			}
		})
	}

	teardown()
}

func TestClient_Add(t *testing.T) {
	setup()

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
		document   map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Adds a doc to the collection",
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
				document: map[string]interface{}{
					"Attributes": map[string]interface{}{
						"color": "blue", "size": 200,
					},
					"CategoryIDs": [...]string{"3", "4", "5"},
					"ID":          "2",
					"Name":        "Another Test Product",
				},
			},
			want:    "{\"Attributes\":{\"color\":\"blue\",\"size\":200},\"CategoryIDs\":[\"3\",\"4\",\"5\"],\"ID\":\"2\",\"Name\":\"Another Test Product\"}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:    tt.fields.config,
				firestore: tt.fields.firestore,
			}
			got, err := c.Add(tt.args.ctx, tt.args.collection, tt.args.document)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Client.Add() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			if got.String() != tt.want {
				t.Errorf("Client.Add().String() = %s, want %v", got.String(), tt.want)
			}
		})
	}

	teardown()
}

func TestClient_Delete(t *testing.T) {
	setup()

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
			name: "deletes a doc without error",
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
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:    tt.fields.config,
				firestore: tt.fields.firestore,
			}
			err := c.Delete(tt.args.ctx, tt.args.collection, tt.args.docID)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
		})
	}

	teardown()
}

func setup() {
	helper = test.NewHelper()
	helper.CreateData()
}

func teardown() {
	// Check if helper is initialized
	if helper == nil {
		return
	}

	// Delete all documents and close helper
	if err := helper.DeleteAll(); err != nil {
		panic(err)
	}
	err := helper.Close()
	if err != nil {
		panic(err)
	}

	// Set helper to nil to prevent re-teardown
	helper = nil
}
