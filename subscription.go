package braze

import (
	"context"
	"net/http"
)

type SubscriptionEndPoint interface {
	StatusSetV2(ctx context.Context, r *SubscriptionStatusSetV2Request) (*Response, error)
}

const (
	subscriptionStatusSetV2Path = "/v2/subscription/status/set"
)

type SubscriptionStatusSetV2Request struct {
	SubscriptionGroups []SubscriptionGroup `json:"subscription_groups"`
}

type SubscriptionGroup struct {
	SubscriptionGroupId string            `json:"subscription_group_id"`
	SubscriptionState   SubscriptionState `json:"subscription_state"`
	ExternalIDs         *[]string         `json:"external_ids,omitempty"`
	Emails              *[]string         `json:"emails,omitempty"`
	Phones              *[]string         `json:"phones"`
}

type SubscriptionState string

var (
	SubscriptionStateSubscribed   SubscriptionState = "subscribed"
	SubscriptionStateUnsubscribed SubscriptionState = "unsubscribed"
)

type SubscriptionService struct {
	client *Client
}

func (s *SubscriptionService) StatusSetV2(ctx context.Context, r *SubscriptionStatusSetV2Request) (*Response, error) {
	req, err := s.client.http.newRequest(http.MethodPost, subscriptionStatusSetV2Path, r)
	if err != nil {
		return nil, err
	}

	var res Response
	if err := s.client.http.do(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
