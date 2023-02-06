package web

type CreateRequest struct {
	LongLink string `validate:"required" json:"longlink"`
}
