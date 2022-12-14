package models

import (
	"log"
	"strconv"

	excelAdapterModels "github.com/HyperloopUPV-H8/Backend-H8/excel_adapter/models"
)

type OrderData map[string]OrderDescription

func (orderData *OrderData) AddPacket(board string, ip string, desc excelAdapterModels.Description, values []excelAdapterModels.Value) {
	if desc.Type != "order" {
		return
	}

	id, err := strconv.ParseUint(desc.ID, 10, 16)
	if err != nil {
		log.Fatalf("order transfer: AddPacket: %s\n", err)
	}

	fields := make(map[string]string, len(values))
	for _, value := range values {
		fields[value.Name] = value.Type
	}

	(*orderData)[desc.Name] = OrderDescription{
		ID:     uint16(id),
		Name:   desc.Name,
		Fields: fields,
	}
}

type OrderDescription struct {
	ID     uint16            `json:"id"`
	Name   string            `json:"name"`
	Fields map[string]string `json:"fieldDescriptions"`
}
