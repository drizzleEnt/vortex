package deployer

import "context"

func (d *deployer) GetActiveList(ctx context.Context) ([]int64, error) {
	res, err := d.repo.GetActiveList(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
