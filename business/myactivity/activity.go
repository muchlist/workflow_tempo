package myactivity

import "context"

func Activity(ctx context.Context, name string) (string, error) {
	return "Hello " + name + "!", nil
}
