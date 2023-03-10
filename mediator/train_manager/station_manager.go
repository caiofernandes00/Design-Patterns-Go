package train_manager

import m "patterns/mediator"

type StationManager struct {
	isPlatformFree bool
	trainQueue     []m.Train
}

func NewStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) CanArrive(t m.Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
