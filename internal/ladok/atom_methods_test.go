package ladok

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomRun(t *testing.T) {
	tts := []struct {
		name        string
		keys        []int
		id          string
		latestID    int
		currentID   int
		queueLength int
	}{
		{
			name:        "OK",
			keys:        []int{99, 100},
			latestID:    98,
			currentID:   100,
			queueLength: 2,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())
			defer server.Close()
			redisMock.ExpectHGet("testSchoolName", "latest").SetVal(strconv.Itoa(tt.latestID))
			for _, key := range tt.keys {
				redisMock.ExpectHSet("testSchoolName", "latest", key).SetVal(int64(key))
			}

			service.Atom.run(t.Context())

			assert.Equal(t, tt.queueLength, len(service.Atom.Channel))

			if err := redisMock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestAtomRunNoNewEntries(t *testing.T) {
	tts := []struct {
		name     string
		latestID int
	}{
		{
			name:     "no new entries - same ID",
			latestID: 100,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())
			defer server.Close()
			redisMock.ExpectHGet("testSchoolName", "latest").SetVal(strconv.Itoa(tt.latestID))

			service.Atom.run(t.Context())

			assert.Equal(t, 0, len(service.Atom.Channel))
		})
	}
}

func TestAtomRunServerError(t *testing.T) {
	service, server, _, _ := mockService(t, 500, 0, 100, t.TempDir())
	defer server.Close()

	// Should handle error gracefully without panic
	service.Atom.run(t.Context())
	assert.Equal(t, 0, len(service.Atom.Channel))
}

func TestAddToCache(t *testing.T) {
	tts := []struct {
		name string
		want int
		have int
	}{
		{
			name: "OK",
			want: 1,
			have: 1,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())

			redisMock.ExpectHSet("testSchoolName", "latest", tt.have).SetVal(int64(tt.have))
			redisMock.ExpectHGet("testSchoolName", "latest").SetVal(strconv.Itoa(tt.have))
			err := service.Atom.addToCache(t.Context(), tt.have)
			assert.NoError(t, err)

			got, err := service.Atom.db.HGet(
				t.Context(),
				"testSchoolName",
				"latest",
			).Int()
			assert.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnprocessedIDs(t *testing.T) {
	tts := []struct {
		name      string
		currentID int
		latestID  int
		want      []int
	}{
		{
			name:      "delta 0, 1",
			currentID: 1,
			latestID:  0,
			want:      []int{1},
		},
		{
			name:      "delta 0",
			currentID: 100,
			latestID:  100,
			want:      nil,
		},
		{
			name:      "delta 1",
			currentID: 100,
			latestID:  99,
			want:      []int{100},
		},
		{
			name:      "delta 2",
			currentID: 100,
			latestID:  98,
			want:      []int{99, 100},
		},
		{
			name:      "delta 26",
			currentID: 26,
			latestID:  1,
			want:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())
			redisMock.ExpectHGet("testSchoolName", "latest").SetVal(strconv.Itoa(tt.latestID))

			ids, err := service.Atom.unprocessedIDs(t.Context(), tt.currentID)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Equal(t, tt.want, ids)

			if err := redisMock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestLatestID(t *testing.T) {
	tts := []struct {
		name string
		want int
	}{
		{
			name: "0",
			want: 0,
		},
		{
			name: "1",
			want: 1,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _, redisMock, _ := mockService(t, 200, 0, 100, t.TempDir())

			redisMock.ExpectHGet("testSchoolName", "latest").SetVal(strconv.Itoa(tt.want))

			got, err := service.Atom.latestID(t.Context())
			assert.NoError(t, err)

			assert.Equal(t, tt.want, got)

			if err := redisMock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}
