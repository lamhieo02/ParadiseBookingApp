package policieshandler

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/policy/iomodel"
)

type PolicyUseCase interface {
	UpsertPolicy(ctx context.Context, dataReq *iomodel.CreatePolicyReq) error
	GetPolicyByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.Policy, error)
}

type policyHandler struct {
	policyUC PolicyUseCase
}

func NewPolicyHandler(policy PolicyUseCase) *policyHandler {
	return &policyHandler{policyUC: policy}
}
