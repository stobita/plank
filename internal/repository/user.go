package repository

import "context"

func (r *repository) GetUserSession(sid string) error {
	ctx := context.Background()
	_, err := r.redisClient.Get(ctx, sid).Result()
	if err != nil {
		return err
	}
	return nil
}
