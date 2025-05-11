package grpcclient

import (
	"context"
	postpb "postaggregator/proto"

	"google.golang.org/grpc"
)

type GRPCClient struct {
	Conn   *grpc.ClientConn
	Client postpb.PostServiceClient
}

func NewGRPCClient(serverAddr string) (*GRPCClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := postpb.NewPostServiceClient(conn)
	return &GRPCClient{
		Conn:   conn,
		Client: client,
	}, nil
}

func (c *GRPCClient) Close() {
	c.Conn.Close()
}

func (c *GRPCClient) ListPostsByUser(userID int32) ([]*postpb.Post, error) {
	req := &postpb.ListPostsRequest{UserId: userID}
	res, err := c.Client.ListPostsByUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Posts, nil
}

// lists the following of the user
func (c *GRPCClient) ListFollowing(userID int32) ([]int32, error) {
	req := &postpb.ListFollowingRequest{UserId: userID}
	res, err := c.Client.ListFollowing(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.FollowingIds, nil
}

// list user feed from fetching from its followers
func (c *GRPCClient) GetUserFeed(userID int32) ([]*postpb.Post, error) {
	req := &postpb.ListPostsRequest{UserId: userID}
	res, err := c.Client.GetUserFeed(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Posts, nil
}
