package fcsnstoragesdk

import (
	"context"
	"time"

	pb "github.com/m110h/fcsn-proto/gen/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

/* Some description */
type Connection struct {
	Connection *grpc.ClientConn
}

/* Some description */
func (c *Connection) Connect(target string, isInsecure, isTraced bool) error {
	var opts []grpc.DialOption
	if isInsecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if isTraced {
		opts = append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	}
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return err
	}
	c.Connection = conn
	return nil
}

/* Some description */
func (c *Connection) Close() {
	c.Connection.Close()
}

/* Some description */
type Client struct {
	client pb.FamilyRootsStorageClient
}

/* Some description */
func NewClient(conn *Connection) *Client {
	client := Client{
		client: pb.NewFamilyRootsStorageClient(conn.Connection),
	}
	return &client
}

/* Some description */
func NewConnectionAndClient(target string, isInsecure, isTraced bool) (*Connection, *Client, error) {
	var conn Connection

	err := conn.Connect(target, isInsecure, isTraced)
	if err != nil {
		return nil, nil, err
	}

	cli := NewClient(&conn)

	return &conn, cli, nil
}

/* Some description */
func (c *Client) SetChart(ctx context.Context, uid, chart string) (time.Time, error) {
	req := &pb.SetChartRequest{
		Uid:   uid,
		Chart: chart,
	}
	r, err := c.client.SetChart(ctx, req)
	if err != nil {
		return time.Now(), err
	}

	return r.UpdatedAt.AsTime(), nil
}

/* Some description */
func (c *Client) GetChart(ctx context.Context, uid string) (string, error) {
	req := &pb.GetChartRequest{
		Uid: uid,
	}

	res, err := c.client.GetChart(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Chart, nil
}

/* Some description */
func (c *Client) Subscribe(ctx context.Context, source_uid, target_uid string) (time.Time, error) {
	req := &pb.SubscribeRequest{
		SourceUid: source_uid,
		TargetUid: target_uid,
	}

	r, err := c.client.Subscribe(ctx, req)
	if err != nil {
		return time.Now(), err
	}

	return r.CreatedAt.AsTime(), nil
}

/* Some description */
func (c *Client) Unsubscribe(ctx context.Context, source_uid, target_uid string) (time.Time, error) {
	req := &pb.UnsubscribeRequest{
		SourceUid: source_uid,
		TargetUid: target_uid,
	}

	r, err := c.client.Unsubscribe(ctx, req)
	if err != nil {
		return time.Now(), err
	}

	return r.DeletedAt.AsTime(), nil
}

/* Some description */
func (c *Client) IsSubscriber(ctx context.Context, target_uid, source_uid string) (bool, error) {
	req := &pb.IsSubscriberRequest{
		TargetUid: target_uid,
		SourceUid: source_uid,
	}

	r, err := c.client.IsSubscriber(ctx, req)
	if err != nil {
		return false, err
	}

	return r.IsSubscriber, nil
}

/* Some description */
func (c *Client) GetSubscribersCount(ctx context.Context, uid string) (int64, error) {
	req := &pb.GetSubscribersCountRequest{
		Uid: uid,
	}

	r, err := c.client.GetSubscribersCount(ctx, req)
	if err != nil {
		return 0, err
	}

	return r.Count, nil
}

/* Some description */
func (c *Client) GetSubscribers(ctx context.Context, uid string) ([]string, error) {
	req := &pb.GetSubscribersRequest{
		Uid: uid,
	}

	r, err := c.client.GetSubscribers(ctx, req)
	if err != nil {
		return []string{}, err
	}

	return r.Subscribers, nil
}

/* Some description */
func (c *Client) GetSubscribtionsCount(ctx context.Context, uid string) (int64, error) {
	req := &pb.GetSubscribtionsCountRequest{
		Uid: uid,
	}

	r, err := c.client.GetSubscribtionsCount(ctx, req)
	if err != nil {
		return 0, err
	}

	return r.Count, nil
}

/* Some description */
func (c *Client) GetSubscribtions(ctx context.Context, uid string) ([]string, error) {
	req := &pb.GetSubscribtionsRequest{
		Uid: uid,
	}

	r, err := c.client.GetSubscribtions(ctx, req)
	if err != nil {
		return []string{}, err
	}

	return r.Subscribtions, nil
}
