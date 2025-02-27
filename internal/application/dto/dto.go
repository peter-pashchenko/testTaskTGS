package dto

import (
	pb "github.com/peter-pashchenko/imagesService/internal/generated/grpc/images"
	"github.com/peter-pashchenko/imagesService/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertToPBList(images []*models.Image) *pb.ListImages {
	res := make([]*pb.ImageList, 0, len(images))

	for _, image := range images {
		toAdd := &pb.ImageList{
			Name:    image.Name,
			Created: timestamppb.New(*image.CreatedAt),
		}
		if image.UpdatedAt != nil {
			toAdd.Updated = timestamppb.New(*image.UpdatedAt)
		}

		res = append(res, toAdd)
	}

	return &pb.ListImages{Images: res}
}
