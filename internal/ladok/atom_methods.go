package ladok

import (
	"context"
	"fmt"
	"time"

	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/model"

	"github.com/go-redis/redis/v8"
	"github.com/SUNET/goladok3"
	"github.com/SUNET/goladok3/ladoktypes"
)

func (s *AtomService) run(ctx context.Context) {
	superFeed, resp, err := s.ladok.Feed.Recent(ctx)
	if err != nil {
		s.logger.Warn("recent_feed_error", "error", err, "response", helpers.FormatResponse(resp))
	}
	if superFeed == nil {
		s.logger.Warn("Feed return nil")
		return
	}
	currentID := superFeed.ID

	if currentID == 0 {
		s.logger.Info("No feedID, try again later")
		return
	}

	ids, err := s.unprocessedIDs(ctx, currentID)
	if err != nil {
		s.logger.Warn("unprocessed_ids_error", "error", err, "currentID", currentID)
		return
	}
	if ids == nil {
		s.logger.Info(fmt.Sprintf("No feed entry to process, resting for %d seconds", s.Service.config.Ladok.Atom.Periodicity))
		return
	}

	for _, id := range ids {
		superFeed, resp, err := s.ladok.Feed.Historical(ctx, &goladok3.HistoricalReq{ID: id})
		if err != nil {
			s.logger.Warn("historical_feed_error", "error", err, "feedID", id, "response", helpers.FormatResponse(resp))
			return
		}
		s.process(ctx, superFeed)
	}
	return
}

func (s *AtomService) process(ctx context.Context, superFeed *ladoktypes.SuperFeed) {
	for _, superEvent := range superFeed.SuperEvents {
		if superEvent.EventTypeName == ladoktypes.LokalStudentEventName {
			s.logger.Info("Adding message to queue", "id", fmt.Sprintf("%d:%s", superFeed.ID, superEvent.EntryID))
			channelEvent := model.LadokToAggregateMSG{
				Event: superEvent,
				TS:    time.Now(),
			}
			s.Channel <- &channelEvent
		}
	}
	if err := s.addToCache(ctx, superFeed.ID); err != nil {
		s.logger.Warn("add_to_cache_error", "error", err, "feedID", superFeed.ID)
	}
}

func (s *AtomService) latestID(ctx context.Context) (int, error) {
	id, err := s.db.HGet(ctx, s.Service.schoolName, "latest").Int()
	if err == redis.Nil {
		return 0, nil
	}
	return id, nil
}

func (s *AtomService) unprocessedIDs(ctx context.Context, currentID int) ([]int, error) {
	latestID, err := s.latestID(ctx)
	if err != nil {
		return nil, err
	}

	delta := currentID - latestID
	ids := []int{}

	switch {
	case delta == 0:
		return nil, nil
	case delta == 1:
		return []int{currentID}, nil
	case delta >= 25:
		for id := currentID - 25 + 1; id <= currentID; id++ {
			ids = append(ids, id)
		}
		return ids, nil
	default:
		for id := latestID + 1; id <= currentID; id++ {
			ids = append(ids, id)
		}
		return ids, nil
	}
}

func (s *AtomService) addToCache(ctx context.Context, id int) error {
	if err := s.db.HSet(ctx, s.Service.schoolName, "latest", id).Err(); err != nil {
		return err
	}
	s.logger.Info("Adding feed id to cache", "id", id)

	return nil
}
