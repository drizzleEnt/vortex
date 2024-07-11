package deployer

func (s *deployer) DeletePod(name string) error {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.pods, name)

	return nil
}
