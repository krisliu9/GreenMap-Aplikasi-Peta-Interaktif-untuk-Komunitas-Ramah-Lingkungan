package usecase

import (
	"mini-project/repository"
)

type TierUseCase struct {
	TierRepo repository.TierRepository
}

func NewTierUseCase(repo repository.TierRepository) *TierUseCase {
	return &TierUseCase{
		TierRepo: repo,
	}
}

func (usecase *TierUseCase) CreateTier(tier repository.Tier) (repository.Tier, error) {
	tier, err := usecase.TierRepo.CreateTier(tier)
	if err != nil {
		return repository.Tier{}, err
	}
	return tier, nil
}

func (usecase *TierUseCase) UpdateTier(id uint, tier repository.Tier) (repository.Tier, error) {
	tier, err := usecase.TierRepo.UpdateTier(id, tier)
	if err != nil {
		return repository.Tier{}, err
	}
	return tier, nil
}

func (usecase *TierUseCase) DeleteTier(id uint) error {
	err := usecase.TierRepo.DeleteTier(id)
	if err != nil {
		return err
	}
	return nil
}
