package services

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang/glog"
	"spotify/client"
	"spotify/dto"
	"spotify/helper/common"
	"time"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type MediaService interface {
	Upload(in dto.UploadIn) (*uploader.UploadResult, common.SubReturnCode)
}

func NewMediaService(cldClient client.CloudinaryAPI) MediaService {
	return &mediaServiceImpl{
		cld: cldClient,
	}
}

type mediaServiceImpl struct {
	cld client.CloudinaryAPI
}

func (m mediaServiceImpl) Upload(in dto.UploadIn) (*uploader.UploadResult, common.SubReturnCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	resp, err := m.cld.Upload(ctx, in)
	cancel()
	if err != nil {
		glog.Errorln("Upload fail err: ", err)
		return nil, common.SystemError
	}
	return resp, common.OK
}
