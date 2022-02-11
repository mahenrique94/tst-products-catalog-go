package dtos

type ProductDTOIn struct {
	Name        string  `json:"name" binding:"max=120,min=4,required"`
	Description string  `json:"description" binding:"max=255,min=10,required"`
	Price       float32 `json:"price" binding:"min=0.0,required"`
	Quantity    uint16  `json:"quantity" binding:"min=0,required"`
}
