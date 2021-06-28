package server

import (
	"context"
	"fmt"

	pb "github.com/weaveworks/weave-gitops/pkg/api/applications"
	"github.com/weaveworks/weave-gitops/pkg/kube"
)

type server struct {
	pb.UnimplementedApplicationsServer

	kube kube.Kube
}

func NewApplicationsServer(kubeSvc kube.Kube) pb.ApplicationsServer {
	return &server{
		kube: kubeSvc,
	}
}

func (s *server) ListApplications(ctx context.Context, msg *pb.ListApplicationsRequest) (*pb.ListApplicationsResponse, error) {

	fakeApps := []string{"cool-app", "slick-app", "neat-app"}

	list := []*pb.Application{}

	for _, a := range fakeApps {
		list = append(list, &pb.Application{Name: a})
	}
	return &pb.ListApplicationsResponse{
		Applications: list,
	}, nil
}

func (s *server) GetApplication(ctx context.Context, msg *pb.GetApplicationRequest) (*pb.GetApplicationResponse, error) {
	app, err := s.kube.GetApplication(msg.ApplicationName)

	if err != nil {
		return nil, fmt.Errorf("could not get application: %s", err)
	}

	return &pb.GetApplicationResponse{Application: &pb.Application{
		Name: app.Name,
		Url:  app.Spec.URL,
		Path: app.Spec.Path,
	}}, nil
}