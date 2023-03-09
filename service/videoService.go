package service

import (


	"example.com/go-walkthrough/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct{
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}


func (service *videoService) Save(video entity.Video) entity.Video{
	service.videos = append(service.videos, video)
	println(service.videos)
	return video
}

func (service *videoService) FindAll() []entity.Video{
	return service.videos
}
