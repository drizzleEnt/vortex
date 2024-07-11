package deployer

func (s *deployer) GetPodList() ([]string, error) {
	names := make([]string, 0, len(s.pods))

	for k := range s.pods {
		names = append(names, k)
	}

	return names, nil
}
