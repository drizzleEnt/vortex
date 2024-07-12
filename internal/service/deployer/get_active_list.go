package deployer

import "context"

func (d *deployer) GetActiveList(ctx context.Context) ([]int64, []int64, error) {
	resActive, resInactive, err := d.repo.GetActiveList(ctx)
	if err != nil {
		return nil, nil, err
	}
	return resActive, resInactive, nil
}
