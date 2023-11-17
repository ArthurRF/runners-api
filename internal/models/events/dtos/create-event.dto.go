package eventDtos

type CreateEventDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
