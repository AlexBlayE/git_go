package models

import "git_go/src/schemas"

var User = CreateMongoModel("User", schemas.User{}).setUniqueField("name")

var Repository = CreateMongoModel("Repository", schemas.Repository{})
