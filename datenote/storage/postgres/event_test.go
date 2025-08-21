package postgres

import (
	"context"
	"datenote/datenote/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateEvent(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Event
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_EVENT_SUCCESS",
			in: storage.Event{
				Name:     "John Doe",
				Date:     "29/09/1996",
				Info:     "Bla Bla Bla",
				Category: "Birthday",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateEvent(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEvent(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		want    storage.Event
		wantErr bool
	}{
		{
			name: "GET_EVENT_SUCCESS",
			in:   1,
			want: storage.Event{
				ID:       1,
				Name:     "John Doe",
				Date:     "29/09/1996",
				Info:     "Bla Bla Bla",
				Category: "Birtthday",
			},
		},
		{
			name:    "GET_EVENT_INVALID",
			in:      100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetEvent(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestGetAllEvents(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		want    []storage.Event
		wantErr bool
	}{
		{
			name: "GET_ALL_EVENTS_SUCCESS",
			want: []storage.Event{
				{
					ID:       1,
					Name:     "John Doe",
					Date:     "29/09/1996",
					Info:     "Bla Bla Bla",
					Category: "Birthday",
				},
				{
					ID:       2,
					Name:     "Melinda Doe",
					Date:     "05/12/1996",
					Info:     "Bla Bla Bla Bla Bla",
					Category: "Birthday",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := s.GetAllEvents(context.Background())
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

func TestUpdateEvent(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      storage.Event
		wantErr bool
	}{
		{
			name: "UPDATE_EVENT_SUCCESS",
			in: storage.Event{
				ID:       1,
				Name:     "John Doe",
				Date:     "29/09/1996",
				Info:     "Bla Bla Bla",
				Category: "Birthday",
			},
			wantErr: false,
		},
		{
			name:    "UPDATE_EVENT_FAILED",
			in:      storage.Event{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.UpdateEvent(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteEvent(t *testing.T) {
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name:    "DELETE_EVENT_SUCCESS",
			in:      1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteEvent(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSearchPost(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      string
		want    []storage.Event
		wantErr bool
	}{
		{
			name: "SEARCH_POST_SUCCESS",
			in:   "Title",
			want: []storage.Event{
				{
					ID:       1,
					Name:     "John Doe",
					Date:     "29.09/1996",
					Info:     "Bla Bla Bla",
					Category: "Birthday",
				},
				{
					ID:       2,
					Name:     "Melinda Doe",
					Date:     "05/12/1996",
					Info:     "Bla Bla Bla Bla Bla",
					Category: "Birthday",
				},
			},
		},
		{
			name: "SEARCH_EVENT_NOT_SUCCESS",
			in:   "Title",
			want: []storage.Event{
				{
					ID:       3,
					Name:     "John Doe",
					Date:     "29/05/1995",
					Info:     "Bla Bla Bla",
					Category: "Uncategorized",
				},
				{
					ID:       4,
					Name:     "Melinda Doe",
					Date:     "12/03/1995",
					Info:     "Bla Bla Bla Bla Bla",
					Category: "Uncategorized",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := s.SearchEvent(context.Background(), tt.in)
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
