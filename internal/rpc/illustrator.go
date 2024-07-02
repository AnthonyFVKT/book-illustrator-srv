package rpc

import (
	"context"
	"fmt"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/model"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/service"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
	"sync"
)

type IllustratorService interface {
	Create(ctx context.Context, req *pb.CreateRequest) (pb.CreateResponse, error)
}

type Illustrator struct {
	illustrator   *service.Illustrator
	textProcessor *service.Processor

	pb.UnsafeIllustratorServiceServer
}

func NewIllustrator(illustrator *service.Illustrator, textProcessor *service.Processor) *Illustrator {
	return &Illustrator{
		illustrator:   illustrator,
		textProcessor: textProcessor,
	}
}

func (il *Illustrator) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	illustrated := make([]*model.Illustrated, 0)
	requestLimiter := make(chan struct{}, 3)
	mu := &sync.Mutex{}
	for i, v := range il.textProcessor.SplitText(req.Text) {
		requestLimiter <- struct{}{}
		go func(mu *sync.Mutex, i int, v string) {
			defer func() { <-requestLimiter }()

			imgURL, err := il.illustrator.MakeImageByText(ctx, v, true)
			if err != nil {
				fmt.Printf("Generate image %d failed: %s \n", i, err.Error())
				return
			}

			mu.Lock()
			illustrated = append(illustrated, &model.Illustrated{
				Text:     v,
				ImageURL: imgURL,
			})
			mu.Unlock()

			fmt.Printf("Task %d done...\n", i)
		}(mu, i, v)
	}
	for i := 0; i < cap(requestLimiter); i++ {
		requestLimiter <- struct{}{}
	}
	close(requestLimiter)

	// TODO: remove
	for _, v := range illustrated {
		fmt.Println("text: ", v.Text)
		fmt.Println("img: ", v.ImageURL)
		fmt.Println()
	}

	return model.FullIllustratedToPb(illustrated), nil
}
