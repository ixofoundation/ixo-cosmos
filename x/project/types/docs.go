package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported")

//type ProjectDoc struct {
//	TxHash     string          `json:"txHash" yaml:"txHash"`
//	ProjectDid didexported.Did         `json:"projectDid" yaml:"projectDid"`
//	SenderDid  didexported.Did         `json:"senderDid" yaml:"senderDid"`
//	PubKey     string          `json:"pubKey" yaml:"pubKey"`
//	Status     ProjectStatus   `json:"status" yaml:"status"`
//	Data       json.RawMessage `json:"data" yaml:"data"`
//}

func NewProjectDoc(txHash string, projectDid, senderDid didexported.Did,
	pubKey string, status ProjectStatus, data json.RawMessage) ProjectDoc {
	return ProjectDoc{
		TxHash:     txHash,
		ProjectDid: projectDid,
		SenderDid:  senderDid,
		PubKey:     pubKey,
		Status:     string(status),
		Data:       data,
	}
}

func (pd ProjectDoc) GetProjectData() (dataMap ProjectDataMap) {
	err := json.Unmarshal(pd.Data, &dataMap)
	if err != nil {
		panic(err)
	}
	return dataMap
}

func (pd ProjectDoc) GetProjectFeesMap() (feesMap ProjectFeesMap) {
	feesMapRaw := pd.GetProjectData()["fees"]
	err := json.Unmarshal(feesMapRaw, &feesMap)
	if err != nil {
		panic(err)
	}
	return feesMap
}

//type UpdateProjectStatusDoc struct {
//	Status          ProjectStatus `json:"status" yaml:"status"`
//	EthFundingTxnID string        `json:"ethFundingTxnID" yaml:"ethFundingTxnID"`
//}

func NewUpdateProjectStatusDoc(status ProjectStatus, ethFundingTxnID string) UpdateProjectStatusDoc {
	return UpdateProjectStatusDoc{
		Status:          string(status),
		EthFundingTxnId: ethFundingTxnID,
	}
}

//type CreateAgentDoc struct {
//	AgentDid didexported.Did `json:"did" yaml:"did"`
//	Role     string  `json:"role" yaml:"role"`
//}

func NewCreateAgentDoc(agentDid didexported.Did, role string) CreateAgentDoc {
	return CreateAgentDoc{
		AgentDid: agentDid,
		Role:     role,
	}
}

//type UpdateAgentDoc struct {
//	Did    didexported.Did     `json:"did" yaml:"did"`
//	Status AgentStatus `json:"status" yaml:"status"`
//	Role   string      `json:"role" yaml:"role"`
//}

func NewUpdateAgentDoc(did didexported.Did, status AgentStatus, role string) UpdateAgentDoc {
	return UpdateAgentDoc{
		Did:    did,
		Status: status,
		Role:   role,
	}
}

//type CreateClaimDoc struct {
//	ClaimID         string `json:"claimID" yaml:"claimID"`
//	ClaimTemplateID string `json:"claimTemplateID" yaml:"claimTemplateID"`
//}

func NewCreateClaimDoc(claimId string, claimTemplateID string) CreateClaimDoc {
	return CreateClaimDoc{
		ClaimId:         claimId,
		ClaimTemplateId: claimTemplateID,
	}
}

//type CreateEvaluationDoc struct {
//	ClaimID string      `json:"claimID" yaml:"claimID"`
//	Status  ClaimStatus `json:"status" yaml:"status"`
//}

func NewCreateEvaluationDoc(claimId string, status ClaimStatus) CreateEvaluationDoc {
	return CreateEvaluationDoc{
		ClaimId: claimId,
		Status:  string(status),
	}
}

//type WithdrawalInfoDoc struct {
//	ProjectDid   didexported.Did  `json:"projectDid" yaml:"projectDid"`
//	RecipientDid didexported.Did  `json:"recipientDid" yaml:"recipientDid"`
//	Amount       sdk.Coin `json:"amount" yaml:"amount"`
//}

func NewWithdrawalInfoDoc(projectDid, recipientDid didexported.Did, amount sdk.Coin) WithdrawalInfoDoc {
	return WithdrawalInfoDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
	}
}

//type WithdrawFundsDoc struct {
//	ProjectDid   didexported.Did `json:"projectDid" yaml:"projectDid"`
//	RecipientDid didexported.Did `json:"recipientDid" yaml:"recipientDid"`
//	Amount       sdk.Int `json:"amount" yaml:"amount"`
//	IsRefund     bool    `json:"isRefund" yaml:"isRefund"`
//}

func NewWithdrawFundsDoc(projectDid, recipientDid didexported.Did, amount sdk.Int, isRefund bool) WithdrawFundsDoc {
	return WithdrawFundsDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
		IsRefund:     isRefund,
	}
}
