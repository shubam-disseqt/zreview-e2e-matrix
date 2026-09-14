package store

type Item struct {
	ID   string
	Name string
}

type Store struct {
	items map[string]Item
}

func New() *Store {
	return &Store{items: make(map[string]Item)}
}

func (s *Store) Get(id string) (Item, bool) {
	it, ok := s.items[id]
	return it, ok
}

func (s *Store) Put(it Item) {
	s.items[it.ID] = it
}
