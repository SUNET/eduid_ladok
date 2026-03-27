package aggregate

import (
	"context"
	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/model"

	"github.com/SUNET/goladok3"
	"github.com/SUNET/goladok3/ladoktypes"
)

func (s *Service) run(ctx context.Context) {
	s.logger.Info("start run")
	for {
		select {
		case msg := <-s.ladok.Atom.Channel:
			s.logger.Info("received message", msg.Event.EntryID, msg.Event.EventTypeName)

			reply, resp, err := s.ladok.Rest.Ladok.Studentinformation.GetAktivPaLarosate(ctx, &goladok3.GetAktivPaLarosateReq{
				UID: msg.Event.StudentUID,
			})
			if err != nil {
				s.logger.Warn("get_aktiv_pa_larosate_error", "error", err, "response", helpers.FormatResponse(resp))
			}
			if reply != nil {
				for _, r := range reply.Studentkopplingar {
					if r.LarosateID == s.ladok.SchoolID {
						s.logger.Info("Student", r.StudentUID, "is active!")
					}
				}
			}
			//userReply, _, err := s.eduidiam.Users.Search(ctx, &goeduidiam.SearchUsersRequest{
			//	Data: goeduidiam.SearchRequest{
			//		Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"},
			//		Filter:     "",
			//		StartIndex: 0,
			//		Count:      0,
			//		Attributes: []string{},
			//	},
			//})
			//s.logger.Info("Aktiv pa larosate", reply)

			//s.

			//return
		//entry.AddTimestamp("aggregate arrived")
		//s.whatToDo(entry)
		//s.logger.Info("Process event", entry.Payload.EntryID, entry.Payload.EventType)
		case <-s.quitChannel:
			return
		}
		//	default:
		//	}
	}
}

func (s *Service) whatToDo(entry *model.LadokToAggregateMSG) {
	switch entry.Event.EventTypeName {
	case ladoktypes.LokalStudentEventName:
	default:
	}
}
