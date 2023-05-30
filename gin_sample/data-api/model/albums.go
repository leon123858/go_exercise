package model

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`                        // 唯一編號
	Title  string  `json:"title" binding:"required"`  // 標題
	Artist string  `json:"artist" binding:"required"` // 這裡的註解會進到 swagger
	Price  float64 `json:"price" binding:"required"`
}

// albums slice to seed record album data.
var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
