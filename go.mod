module github.com/javier162380/eatfy

require (
	eatfy/app v0.0.0
	eatfy/models v0.0.0
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/gorilla/mux v1.7.1
	github.com/jinzhu/gorm v1.9.4
	github.com/jinzhu/now v1.0.0 // indirect
	github.com/lib/pq v1.0.0
)

replace eatfy/app v0.0.0 => ./app

replace eatfy/models v0.0.0 => ./models
