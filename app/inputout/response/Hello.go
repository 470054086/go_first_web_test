package response

import "first_web/app/model"

type Hello struct {
	User *model.User `json:"user"`
} 