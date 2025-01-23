run migration
goose -dir migrations mysql "root:root@/devbook?charset=utf8&parseTime=True&loc=Local" up