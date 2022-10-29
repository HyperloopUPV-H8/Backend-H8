package dto

type StructureDTO struct {
	packetName   string
	measurements []string
}

func newStructureDTO(column []string) StructureDTO {
	return StructureDTO{
		packetName:   column[0],
		measurements: column[1:],
	}
}