package service

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type VideoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &VideoService{}
}

func (service *VideoService) Save(video entity.Video) entity.Video {
	service.videos = append(servuce, videos, video)
	return video
}

func (service *VideoService) FindAll() []entity.Video {
	return service.videos
}
