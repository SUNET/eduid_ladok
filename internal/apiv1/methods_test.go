package apiv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestESI(t *testing.T) {
	c := &Client{}

	tests := []struct {
		name       string
		externtUID string
		want       string
	}{
		{
			name:       "standard UID",
			externtUID: "abc-123-def",
			want:       "urn:schac:personalUniqueCode:int:esi:ladok.se:externtstudentuid-abc-123-def",
		},
		{
			name:       "empty UID",
			externtUID: "",
			want:       "urn:schac:personalUniqueCode:int:esi:ladok.se:externtstudentuid-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.ESI(t.Context(), tt.externtUID)
			assert.Equal(t, tt.want, got)
		})
	}
}
