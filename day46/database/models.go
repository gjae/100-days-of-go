package database

type Movie struct {
	Name string `bson:"name"`
	Year string `bson:"year"`
	Directors []string `bson:"directors"`
	Writers []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`
}