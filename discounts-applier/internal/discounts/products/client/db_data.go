package client

type DBData struct {
	Database   string
	Collection string
}

func getDBData() DBData {
	return DBData{
		Database:   "myTheresa",
		Collection: "products",
	}
}
