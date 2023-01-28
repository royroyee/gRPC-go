package post

import postpb "github.com/royroyee/gRPC-go/ex3/protos/v1/post"

type PostData struct {
	UserID string
	Posts  []*postpb.PostMessage
}

var UserPosts = []*PostData{
	{
		UserID: "1",
		Posts: []*postpb.PostMessage{
			{
				PostId: "1",
				Author: "",
				Title:  "test1",
				Body:   "test1",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "2",
				Author: "",
				Title:  "test2",
				Body:   "test2",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "3",
				Author: "",
				Title:  "test3",
				Body:   "test3",
				Tags:   []string{"Golang", "context"},
			},
			{
				PostId: "4",
				Author: "",
				Title:  "test4",
				Body:   "test4",
				Tags:   []string{"test4", "test", "test", "test"},
			},
		},
	},
	{
		UserID: "3",
		Posts: []*postpb.PostMessage{
			{
				PostId: "5",
				Author: "",
				Title:  "test-userID:3",
				Body:   "Test-userID:3",
				Tags:   []string{"TEST", "TEST"},
			},
		},
	},
	{
		UserID: "4",
		Posts: []*postpb.PostMessage{
			{
				PostId: "6",
				Author: "",
				Title:  "TEST6",
				Body:   "TEST",
				Tags:   []string{"TEST", "TEST"},
			},
			{
				PostId: "7",
				Author: "",
				Title:  "test7",
				Body:   "test7.body",
				Tags:   []string{"test7", "test7 test7", "test7-test7"},
			},
		},
	},
	{
		UserID: "5",
		Posts: []*postpb.PostMessage{
			{
				PostId: "8",
				Author: "",
				Title:  "test8",
				Body:   "test8",
				Tags:   []string{"test", "test", "test", "test"},
			},
		},
	},
}
