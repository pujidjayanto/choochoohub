package bootstrap

import "context"

type CleanupFunc func(ctx context.Context) error

func ChainCleanup(funcs ...CleanupFunc) CleanupFunc {
	return func(ctx context.Context) error {
		for _, f := range funcs {
			if f == nil {
				continue
			}
			if err := f(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}
