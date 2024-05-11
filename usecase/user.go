package usecase

import (
	"mini-project/auth"
	"mini-project/repository"
)

type UserUseCase struct {
	UserRepo repository.UserRepository
	TierRepo repository.TierRepository
}

func NewUserUseCase(repo repository.UserRepository, tierRepo repository.TierRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
		TierRepo: tierRepo,
	}
}

func (usecase *UserUseCase) Login(email, password string) (repository.User, string, error) {
	user, err := usecase.UserRepo.Login(email, password)
	if err != nil {
		return user, "", err
	}
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return user, token, err
	}
	return user, token, nil
}

func (usecase *UserUseCase) Register(name, email, password, role string) (string, error) {
	user, err := usecase.UserRepo.Register(name, email, password, role)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (usecase *UserUseCase) GetByID(userId uint) (repository.User, error) {
	user, err := usecase.UserRepo.GetByID(userId)
	if err != nil {
		return user, err
	}
	tier, err := usecase.TierRepo.GetTierByID(user.Tier_ID)
	if err != nil {
		return user, err
	}
	user.Tier_Name = tier.Tier_Name
	return user, nil
}

func (usecase *UserUseCase) UpdateTier(userId uint) error {
	user, err := usecase.UserRepo.GetByID(userId)
	if err != nil {
		return err
	}
	tiers, err := usecase.TierRepo.GetAllTier()
	if err != nil {
		return err
	}
	for _, tier := range tiers {
		if user.Current_Point >= tier.Minimal_Point {
			user, err = usecase.UserRepo.UpdateTier(userId, tier.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
