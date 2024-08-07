package singletonService

type SingletonIdService struct {
	idService *IdService
}

func (s *SingletonIdService) NewSingletonIdService() *IdService {
	if s.idService == nil {
		s.idService = newIdService()
		return s.idService
	}

	return s.idService
}

type IdService struct {
	counter int
}

func newIdService() *IdService {
	return &IdService{
		counter: 0,
	}
}

func (s *IdService) Next() int {
	s.counter++
	return s.counter
}
