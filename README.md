run migration
goose -dir migrations mysql "root:root@tcp(localhost:3307)/devbook?charset=utf8&parseTime=True&loc=Local" up