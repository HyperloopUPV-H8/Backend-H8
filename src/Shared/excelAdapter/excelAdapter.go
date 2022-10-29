package excelAdapter

import (
	"github.com/HyperloopUPV-H8/Backend-H8/Shared/excelAdapter/dto"
	"github.com/HyperloopUPV-H8/Backend-H8/Shared/excelRetriever/domain"
)

func GetBoardDTOs(structure domain.Document) map[string]dto.BoardDTO {
	boardDTOs := make(map[string]dto.BoardDTO, len(structure.Sheets))
	for name, sheet := range structure.Sheets {
		boardDTOs[name] = dto.NewBoardDTO(sheet)
	}
	return boardDTOs
}