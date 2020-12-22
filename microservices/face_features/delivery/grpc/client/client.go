package client

import (
	"context"
	proto "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/protobuf"
	models "park_2020/2020_2_tmp_name/models"

	"google.golang.org/grpc"
)

type FaceClient struct {
	client proto.FaceGRPCHandlerClient
}

func NewFaceClient(conn *grpc.ClientConn) *FaceClient {
	c := proto.NewFaceGRPCHandlerClient(conn)
	return &FaceClient{
		client: c,
	}
}

func (fc *FaceClient) HaveFace(ctx context.Context, photo *models.Photo) (bool, error) {
	if photo == nil {
		return false, nil
	}
	protoPhoto := transformIntoGRPCPhoto(photo)

	face, err := fc.client.HaveFace(ctx, protoPhoto)
	if err != nil {
		return false, err
	}
	return face.Have, nil
}

func (fc *FaceClient) AddMask(ctx context.Context, photo *models.Photo) (models.Photo, error) {
	if photo == nil {
		return models.Photo{}, nil
	}
	protoPhoto := transformIntoGRPCPhoto(photo)

	face, err := fc.client.AddMask(ctx, protoPhoto)
	if err != nil {
		return models.Photo{}, err
	}
	return transformIntoPhoto(face), nil
}

func transformIntoGRPCPhoto(photo *models.Photo) *proto.Photo {
	return &proto.Photo{
		Path: photo.Path,
		Mask: photo.Mask,
	}
}

func transformIntoPhoto(photo *proto.Photo) models.Photo {
	return models.Photo{
		Path: photo.Path,
		Mask: photo.Mask,
	}
}
