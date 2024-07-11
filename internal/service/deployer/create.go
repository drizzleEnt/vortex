package deployer

func (s *deployer) CreatePod(name string) error {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.pods[name]
	if !ok {
		s.pods[name] = struct{}{}
	}

	return nil
}
