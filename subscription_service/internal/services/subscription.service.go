package services

import (
	"context"
	"subscription_service/internal/dtos"
	"subscription_service/internal/models"
	"subscription_service/internal/repositories"
)

// SubscriptionService handles subscription related business logic
type SubscriptionService interface {
	// CreateSubscriptionPlan creates a new subscription plan
	CreateSubscriptionPlan(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error)

	// GetSubscriptionPlan gets a subscription plan by its ID
	GetSubscriptionPlan(ctx context.Context, id string) (*models.Subscription, error)

	// UpdateSubscriptionPlan updates a subscription plan
	UpdateSubscriptionPlan(ctx context.Context, id string, subscription *models.Subscription) (*models.Subscription, error)

	// DeleteSubscriptionPlan deletes a subscription plan
	DeleteSubscriptionPlan(ctx context.Context, id string) error

	// ListSubscriptionPlans lists all subscription plans
	ListSubscriptionPlans(ctx context.Context, pagination *dtos.Pagination) ([]*models.Subscription, error)
}

type subscriptionService struct {
	repo repositories.SubscriptionRepository
}

// CreateSubscriptionPlan implements SubscriptionService.
func (p *subscriptionService) CreateSubscriptionPlan(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error) {
	return p.repo.CreateSubscriptionPlan(ctx, subscription)
}

// DeleteSubscription implements SubscriptionService.
func (p *subscriptionService) DeleteSubscriptionPlan(ctx context.Context, id string) error {
	return p.repo.DeleteSubscriptionPlan(ctx, id)
}

// GetSubscription implements SubscriptionService.
func (p *subscriptionService) GetSubscriptionPlan(ctx context.Context, id string) (*models.Subscription, error) {
	return p.repo.GetSubscriptionPlan(ctx, id)
}

// ListSubscriptions implements SubscriptionService.
func (p *subscriptionService) ListSubscriptionPlans(ctx context.Context, pagination *dtos.Pagination) ([]*models.Subscription, error) {
	return p.repo.ListSubscriptionPlans(ctx, pagination)
}

// UpdateSubscription implements SubscriptionService.
func (p *subscriptionService) UpdateSubscriptionPlan(ctx context.Context, id string, subscription *models.Subscription) (*models.Subscription, error) {
	return p.repo.UpdateSubscriptionPlan(ctx, id, subscription)
}

func NewSubscriptionService(repo repositories.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{
		repo: repo,
	}
}
