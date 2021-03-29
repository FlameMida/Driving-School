package request

type InitDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
}
