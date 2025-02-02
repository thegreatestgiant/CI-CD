// Code generated by sqlc. DO NOT EDIT.

// versions:

//   sqlc v1.18.0

// source: users.sql

package database

import (
	"context"
	"testing"
)

func TestQueries_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateUserParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.CreateUser(tt.args.ctx, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
