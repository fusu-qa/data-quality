package task

import "github.com/jasonlvhit/gocron"

func Run(check func()) {
	s := gocron.NewScheduler()
	//err := s.Every(1).Days().At("10:30").Do(check)
	err := s.Every(10).Second().Do(check)
	if err != nil {
		return
	}
	sc := s.Start()
	<-sc
}
