package iomodel

type CreatePolicyReq struct {
	Data DataReqCreatePolicy `json:"data"`
}

type DataReqCreatePolicy struct {
	ObjectID   int          `json:"object_id"`
	ObjectType int          `json:"object_type"`
	ListPolicy []ListPolicy `json:"list_policy"`
}

type ListPolicy struct {
	GroupPolicyID int    `json:"group_policy_id"`
	Name          string `json:"name"`
}
