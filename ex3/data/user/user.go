package user

// If it's real microservice,  get the data from the database.
// but it's a simple example, so I'll use static variable

import userpb "github.com/royroyee/gRPC-go/ex3/protos/v1/user"

var UserData = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "kiny",
		PhoneNumber: "01012345678",
		Age:         21,
	},

	{
		UserId:      "2",
		Name:        "roy",
		PhoneNumber: "01012345678",
		Age:         24,
	},

	{
		UserId:      "3",
		Name:        "mini",
		PhoneNumber: "01012484428",
		Age:         13,
	},

	{
		UserId:      "4",
		Name:        "jenny",
		PhoneNumber: "01012731533",
		Age:         27,
	},

	{
		UserId:      "5",
		Name:        "jamin",
		PhoneNumber: "01024674568",
		Age:         31,
	},
}
