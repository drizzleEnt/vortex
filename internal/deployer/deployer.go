package deployer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/drizzleent/vortex/internal/service"
)

type Deployer struct {
	d       time.Duration
	service service.Deployer
}

func NewDeployer(srv service.Deployer, d time.Duration) *Deployer {
	return &Deployer{
		d:       d,
		service: srv,
	}
}

func (d *Deployer) Run() error {
	for {
		time.Sleep(d.d * time.Second)

		ctx := context.Background()
		resActive, resInactive, err := d.service.GetActiveList(ctx)

		if err != nil {
			return err
		}

		for _, v := range resInactive {
			err := d.service.DeletePod(strconv.FormatInt(v, 10))
			if err != nil {
				return err
			}
		}

		for _, v := range resActive {

			err := d.service.CreatePod(strconv.FormatInt(v, 10))
			if err != nil {
				return err
			}
		}

		list, err := d.service.GetPodList()

		if err != nil {
			return err
		}
		fmt.Println("active clients id", list)
	}
}
