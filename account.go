package cloudns

import (
	"context"
	"net"
)

type AccountService struct {
	api *Client
}

func (svc *AccountService) Login(ctx context.Context) (result BaseResult, err error) {
	err = svc.api.request(ctx, "POST", "/dns/login.json", nil, nil, &result)
	return
}

func (svc *AccountService) GetCurrentIP(ctx context.Context) (net.IP, error) {
	var result struct {
		IP net.IP `json:"ip"`
	}

	err := svc.api.request(ctx, "POST", "/ip/get-my-ip.json", nil, nil, &result)
	return result.IP, err
}

func (svc *AccountService) GetBalance(ctx context.Context) (float64, error) {
	var result struct {
		Funds float64 `json:"funds,string"`
	}

	err := svc.api.request(ctx, "POST", "/account/get-balance.json", nil, nil, &result)
	return result.Funds, err
}
