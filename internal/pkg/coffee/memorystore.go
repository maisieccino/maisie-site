package coffee

import (
	"context"
	"errors"
)

// MemoryStore is an in-memory store of coffee map items. Ideal for unit
// testing or just messing around without a database.
type MemoryStore struct {
	items map[string]MapItem
}

func NewMemoryStore() Store {
	return &MemoryStore{
		items: make(map[string]MapItem),
	}
}

func (m *MemoryStore) Get(_ context.Context, id string) (MapItem, error) {
	if i, ok := m.items[id]; ok {
		return i, nil
	}
	return MapItem{}, errors.New("not found")
}

func (m *MemoryStore) Create(_ context.Context, i MapItem) error {
	if _, ok := m.items[i.ID]; ok {
		return errors.New("id already exists")
	}

	m.items[i.ID] = i
	return nil
}

func (m *MemoryStore) Update(_ context.Context, id string, item MapItem) error {
	if _, ok := m.items[id]; !ok {
		return errors.New("not found")
	}
	m.items[id] = item
	return nil
}

func (m *MemoryStore) List(_ context.Context) ([]MapItem, error) {
	results := []MapItem{}
	for _, item := range m.items {
		results = append(results, item)
	}
	return results, nil
}

func (m *MemoryStore) Delete(_ context.Context, id string) error {
	if _, ok := m.items[id]; !ok {
		return errors.New("not found")
	}
	delete(m.items, id)
	return nil
}

func (s *MemoryStore) SearchByArea(ctx context.Context, params SearchByAreaParams) ([]MapItem, error) {
	return nil, nil
}
