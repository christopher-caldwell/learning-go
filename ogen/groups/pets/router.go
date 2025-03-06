package pets

import (
	"context"

	petstore "oapigen/petstore"
)

type PetsService struct {
	pets map[int64]petstore.Pet
	id   int64
}

var _ petstore.PetHandler = (*PetsService)(nil)

func (p *PetsService) AddPet(ctx context.Context, req *petstore.Pet) (*petstore.Pet, error) {
	p.pets[p.id] = *req
	p.id++
	return req, nil
}

func (p *PetsService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	pet, ok := p.pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &petstore.GetPetByIdNotFound{}, nil
	}
	return &pet, nil
}

var Service = PetsService{
	pets: map[int64]petstore.Pet{},
	id:   0,
}
