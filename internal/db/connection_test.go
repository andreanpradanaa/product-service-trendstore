package db

import (
	"testing"

	"github.com/andreanpradanaa/product-service-trendstore/internal/config"
)

func TestOpenDBConnection(t *testing.T) {

	dbconf := config.GetConfig().Database

	tests := []struct {
		name    string
		dbconf  config.Database
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Sukses",
			dbconf:  dbconf,
			wantErr: false,
		},
		{
			name: "Gagal",
			dbconf: config.Database{
				Host: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := OpenDBConnection(tt.dbconf)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDBCOnnection() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && db == nil {
				t.Error("OpenDBConnection() = nil, padahal diharapkan koneksi berhasil")
			}
		})
	}
}
