package model

import pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"

type Illustrated struct {
	Text     string `json:"text,omitempty"`
	ImageURL string `json:"imageURL,omitempty"`
}

func (il *Illustrated) illustratedToPb() *pb.Illustrated {
	return &pb.Illustrated{
		Text:     il.Text,
		ImageURL: il.ImageURL,
	}
}

func FullIllustratedToPb(data []*Illustrated) *pb.CreateResponse {
	resp := make([]*pb.Illustrated, 0, len(data))
	for _, v := range data {
		resp = append(resp, v.illustratedToPb())
	}

	return &pb.CreateResponse{
		Illustrated: resp,
	}
}
