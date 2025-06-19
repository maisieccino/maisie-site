package music_test

import (
	"testing"

	"github.com/maisieccino/maisie-site/internal/pkg/music"
	"github.com/stretchr/testify/assert"
)

func Test_PlaylistSlug(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "blank test",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "example",
			expected: "example",
		},
		{
			name:     "two words",
			input:    "example playlist",
			expected: "example-playlist",
		},
		{
			name:     "two words, capitals",
			input:    "example Playlist",
			expected: "example-playlist",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			playlist := music.Playlist{Name: tc.input}
			actual := playlist.Slug()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
