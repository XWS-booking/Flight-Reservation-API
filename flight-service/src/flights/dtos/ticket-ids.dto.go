package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type TicketIdsDto struct {
	Ids []string `json:"ids"`
}

func NewTicketIdsDto(objectIds []primitive.ObjectID) *TicketIdsDto {
	ids := make([]string, 0)

	for _, v := range objectIds {
		ids = append(ids, v.Hex())
	}
	return &TicketIdsDto{
		Ids: ids,
	}

}
