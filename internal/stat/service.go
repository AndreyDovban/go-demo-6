package stat

import (
	"go-demo-6/pkg/di"
	"go-demo-6/pkg/event"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data ", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
