package utils

import "context"

func CreateContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.TODO())
	return
}
