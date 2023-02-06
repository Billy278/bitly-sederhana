package web

type ResponseBitly struct {
	Id        int    `validate:"required" json:"id"`
	LongLink  string `validate:"required" json:"longlink"`
	ShortLink string `validate:"required" json:"shortlink"`
}
