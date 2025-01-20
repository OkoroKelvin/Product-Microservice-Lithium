package handlers

import (
	"context"
	"fmt"
	"subscription_service/internal/dtos"
	"subscription_service/internal/models"
	"subscription_service/internal/services"
	"subscription_service/internal/utils"
	pb "subscription_service/proto"
)

// GRPCHandler handles all gRPC related operations

type GRPCHandler struct {
	subscriptionService services.SubscriptionService
	pb.UnimplementedSubscriptionServiceServer
}

// CreateSubscription implements GRPCHandler.
func (g *GRPCHandler) CreateSubscription(ctx context.Context, req *pb.CreateSubscriptionRequest) (*pb.Subscription, error) {
	var (
		subscriptionModel    models.Subscription
		subscriptionResponse *pb.Subscription
	)

	err := utils.UnmarshalToType(req.Subscription, &subscriptionModel)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription request: %w", err)
	}

	subscription, err := g.subscriptionService.CreateSubscriptionPlan(ctx, &subscriptionModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	err = utils.UnmarshalToType(subscription, &subscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription response: %w", err)
	}

	return subscriptionResponse, nil
}

// DeleteSubscription implements GRPCHandler.
func (g *GRPCHandler) DeleteSubscription(ctx context.Context, req *pb.DeleteSubscriptionRequest) (*pb.DeleteSubscriptionResponse, error) {
	response := &pb.DeleteSubscriptionResponse{
		Success: true,
	}

	if _, err := g.subscriptionService.GetSubscriptionPlan(ctx, req.Id); err != nil {
		return nil, fmt.Errorf("subscription not found: %w", err)
	}

	err := g.subscriptionService.DeleteSubscriptionPlan(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete subscription: %w", err)
	}

	return response, nil
}

// GetSubscription implements GRPCHandler.
func (g *GRPCHandler) GetSubscription(ctx context.Context, req *pb.GetSubscriptionRequest) (*pb.Subscription, error) {
	var (
		subscriptionResponse *pb.Subscription
	)

	subscription, err := g.subscriptionService.GetSubscriptionPlan(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	err = utils.UnmarshalToType(subscription, &subscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription response: %w", err)
	}

	return subscriptionResponse, nil
}

// ListSubscriptions implements GRPCHandler.
func (g *GRPCHandler) ListSubscriptions(ctx context.Context, req *pb.ListSubscriptionsRequest) (*pb.ListSubscriptionsResponse, error) {
	var (
		subscriptionResponse []*pb.Subscription
	)

	subscriptions, err := g.subscriptionService.ListSubscriptionPlans(ctx, &dtos.Pagination{
		Page:     int(req.PaginationQuery.Page),
		PageSize: int(req.PaginationQuery.PageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list subscriptions: %w", err)
	}
	err = utils.UnmarshalToType(subscriptions, &subscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription response: %w", err)
	}

	return &pb.ListSubscriptionsResponse{
		Subscriptions: subscriptionResponse,
	}, nil
}

// UpdateSubscription implements GRPCHandler.
func (g *GRPCHandler) UpdateSubscription(ctx context.Context, req *pb.UpdateSubscriptionRequest) (*pb.Subscription, error) {
	var (
		subscriptionModel    models.Subscription
		subscriptionResponse *pb.Subscription
	)

	err := utils.UnmarshalToType(req.Subscription, &subscriptionModel)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription request: %w", err)
	}

	if _, err := g.subscriptionService.GetSubscriptionPlan(ctx, req.Subscription.Id); err != nil {
		return nil, fmt.Errorf("subscription not found: %w", err)
	}

	subscription, err := g.subscriptionService.UpdateSubscriptionPlan(ctx, req.Subscription.Id, &subscriptionModel)
	if err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}

	err = utils.UnmarshalToType(subscription, &subscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal subscription response: %w", err)
	}

	return subscriptionResponse, nil
}

func NewGRPCHandler(subscriptionService services.SubscriptionService) *GRPCHandler {
	return &GRPCHandler{subscriptionService: subscriptionService}
}
