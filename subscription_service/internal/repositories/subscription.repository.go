package repositories

import (
	"context"
	"subscription_service/internal/dtos"
	"subscription_service/internal/models"

	"gorm.io/gorm"
)

// Repository handles database operations
type SubscriptionRepository interface {
	CreateSubscriptionPlan(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error)
	GetSubscriptionPlan(ctx context.Context, id string) (*models.Subscription, error)
	UpdateSubscriptionPlan(ctx context.Context, id string, subscription *models.Subscription) (*models.Subscription, error)
	DeleteSubscriptionPlan(ctx context.Context, id string) error
	ListSubscriptionPlans(ctx context.Context, pagination *dtos.Pagination) ([]*models.Subscription, error)
}

type subscriptionRepository struct {
	db *gorm.DB
}

// CreateSubscription implements SubscriptionRepository.
func (p *subscriptionRepository) CreateSubscriptionPlan(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error) {
	err := p.db.Create(&subscription).Error
	return subscription, err
}

// DeleteSubscription implements SubscriptionRepository.
func (p *subscriptionRepository) DeleteSubscriptionPlan(ctx context.Context, id string) error {
	return p.db.Where("id = ?", id).Delete(&models.Subscription{}).Error
}

// GetSubscription implements SubscriptionRepository.
func (p *subscriptionRepository) GetSubscriptionPlan(ctx context.Context, id string) (*models.Subscription, error) {
	var subscription models.Subscription
	err := p.db.Where("id = ?", id).First(&subscription).Error
	return &subscription, err
}

// ListSubscriptions implements SubscriptionRepository.
func (p *subscriptionRepository) ListSubscriptionPlans(ctx context.Context, pagination *dtos.Pagination) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription

	offset := (pagination.Page - 1) * pagination.PageSize
	err := p.db.Limit(pagination.PageSize).Offset(offset).Order("created_at DESC").Find(&subscriptions).Error
	return subscriptions, err
}

// UpdateSubscription implements SubscriptionRepository.
func (p *subscriptionRepository) UpdateSubscriptionPlan(ctx context.Context, id string, subscription *models.Subscription) (*models.Subscription, error) {
	err := p.db.Model(&models.Subscription{}).Where("id = ?", id).Updates(subscription).Error

	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{
		db: db,
	}
}
