package handler

type Input struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

type Output struct {
	Input
	ID uint `json:"id"`
}

func CreateAlbum() {

}
