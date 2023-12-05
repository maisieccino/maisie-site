package music

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Playlist struct {
	Name     string
	ImageURL string
	URL      string
	Items    []string
}

func (p Playlist) Slug() string {
	return strings.ReplaceAll(strings.ToLower(p.Name), " ", "-")
}

type Track struct {
	Name     string
	Artist   string
	URL      string
	ImageURL string
}

type Playlists interface {
	List() ([]Playlist, error)
}

type StaticPlaylists struct {
	path  string
	items map[string]Playlist
}

func (p *StaticPlaylists) preload() error {
	if p.path == "" {
		return errors.New("path is empty")
	}

	var (
		file []byte
		err  error
	)
	if file, err = os.ReadFile(p.path); err != nil {
		return err
	}
	items := []Playlist{}
	if err = json.Unmarshal(file, &items); err != nil {
		return err
	}

	for _, i := range items {
		p.items[i.Slug()] = i
	}
	return nil
}

func (p *StaticPlaylists) List() ([]Playlist, error) {
	if len(p.items) == 0 {
		if err := p.preload(); err != nil {
			return nil, err
		}
	}
	items := []Playlist{}
	for _, i := range p.items {
		items = append(items, i)
	}
	return items, nil
}
