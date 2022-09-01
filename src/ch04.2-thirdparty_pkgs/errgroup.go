package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Bug struct {
	ctx                     context.Context
	name, sound             string
	lifespan, soundInterval time.Duration
}

func NewBug(ctx context.Context, name, sound string, life, soundInt time.Duration) *Bug {
	return &Bug{
		ctx:           ctx,
		name:          name,
		sound:         sound,
		lifespan:      life,
		soundInterval: soundInt,
	}
}

func (b *Bug) Sing() error {
	sing := time.NewTicker(b.soundInterval)
	defer sing.Stop()
	dead := time.NewTimer(b.lifespan)
	defer dead.Stop()

	for {
		select {
		case <-b.ctx.Done():
			return b.ctx.Err()
		case <-sing.C:
			log.Infof("%s says %s", b.name, b.sound)
		case t := <-dead.C:
			log.Infof("%s died at %v", b.name, t)
			// return nil // 자기의 삶 동안 계속 노래
			return fmt.Errorf("%s died", b.name) // 누구라도 노래가 끝나면 합창 중단
		}
	}
}

// ---

func errgroupDemo() {
	eg, ctx := errgroup.WithContext(context.Background())

	criket := NewBug(ctx, "귀뚜라미", "귀뚤", 10*time.Second, 1*time.Second)
	mosquito := NewBug(ctx, "모기", "에엥", 7*time.Second, 5*time.Second)
	bee := NewBug(ctx, "벌", "우웅", 5*time.Second, 2*time.Second)

	// 벌레 합창단 노래 시작
	eg.Go(criket.Sing)
	eg.Go(bee.Sing)
	eg.Go(mosquito.Sing)

	// 누구라도 노래가 끝나면 그 즉시 중단
	if err := eg.Wait(); err != nil {
		log.Error(err)
	}
}
