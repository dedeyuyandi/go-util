package query

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestWhere(t *testing.T) {
	ids := uuid.New()
	type args struct {
		request map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "uint64",
			args: args{
				request: map[string]interface{}{
					"uint64": uint64(1),
				},
			},
			want: "WHERE uint64 = 1",
		},
		{
			name: "float*",
			args: args{
				request: map[string]interface{}{
					"float64": float64(2.0001),
				},
			},
			want: "WHERE float64 = 2.0001",
		},
		{
			name: "bool",
			args: args{
				request: map[string]interface{}{
					"bool": true,
				},
			},
			want: "WHERE bool = true",
		},
		{
			name: "string",
			args: args{
				request: map[string]interface{}{
					"string": "this is string",
				},
			},
			want: "WHERE string = 'this is string'",
		},
		{
			name: "uuid",
			args: args{
				request: map[string]interface{}{
					"uuid": ids,
				},
			},
			want: fmt.Sprintf("WHERE uuid = %v", ids),
		},
		{
			name: "default",
			args: args{
				request: map[string]interface{}{
					"byte": []byte("this is byte"),
				},
			},
			want: "WHERE byte = 'this is byte'",
		},
		{
			name: "return null",
			args: args{
				request: map[string]interface{}{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Where(tt.args.request); got != tt.want {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}
