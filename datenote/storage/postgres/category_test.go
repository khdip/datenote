package postgres

import (
	"context"
	"datenote/datenote/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateCategory(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Category
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCCESS",
			in: storage.Category{
				Title: "Example Category",
			},
			want: 1,
		},
		{
			name: "FAILED_DUPLICATE_CATEGORY",
			in: storage.Category{
				Title: "Example Category",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCategory(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		want    storage.Category
		wantErr bool
	}{
		{
			name: "GET_CATEGORY_SUCCESS",
			in:   1,
			want: storage.Category{
				ID:    1,
				Title: "Example Category",
			},
		},
		{
			name:    "GET_CATEGORY_INVALID",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCategory(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestGetAllCategories(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []storage.Category
		wantErr bool
	}{
		{
			name: "GET_ALL_CATEGORIES_SUCCESS",
			want: []storage.Category{
				{
					ID:    1,
					Title: "Example Category",
				},
				{
					ID:    2,
					Title: "Example Category 2",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := s.GetAllCategories(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID < tt.want[j].ID
			})

			sort.Slice(gotList, func(i, j int) bool {
				return gotList[i].ID < gotList[j].ID
			})

			for i, got := range gotList {

				if !cmp.Equal(got, tt.want[i]) {
					t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want[i]))
				}

			}

		})
	}
}

func TestUpdateCategory(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Category
		wantErr bool
	}{
		{
			name: "UPDATE_CATEGORY_SUCCESS",
			in: storage.Category{
				ID:    1,
				Title: "Example Category",
			},
			wantErr: false,
		},
		{
			name:    "UPDATE_CATEGORY_FAILED",
			in:      storage.Category{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.UpdateCategory(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name:    "DELETE_CATEGORY_SUCCESS",
			in:      1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteCategory(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
