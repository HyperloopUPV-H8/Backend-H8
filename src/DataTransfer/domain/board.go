package domain

import excelParser "github.com/HyperloopUPV-H8/Backend-H8/Shared/ExcelParser/domain/board"

type Board struct {
	Name    string
	Packets map[uint16]Packet
}

func NewBoard(rawBoard excelParser.Board) Board {
	return Board{
		Name:    rawBoard.Name,
		Packets: NewPackets(rawBoard.GetPackets()),
	}
}