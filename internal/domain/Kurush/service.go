package Kurush

type Service struct {
	adapter IAdapter
	storage IStorage
}

func NewService(adapter IAdapter, storage IStorage) *Service {
	return &Service{
		adapter: adapter,
		storage: storage,
	}
}
