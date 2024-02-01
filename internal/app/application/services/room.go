package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
)

type RoomService struct {
	roomRepository repositories.RoomRepository
}

func NewRoomService(
	roomRepository repositories.RoomRepository,
) *RoomService {
	return &RoomService{
		roomRepository,
	}
}

func (s *RoomService) Create(name string, description *string, userID string) (*entities.Room, error) {
	room, err := entities.NewRoom(name, description, userID)
	if err != nil {
		return nil, err
	}

	err = s.roomRepository.Create(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) GetAll(
	search string,
	userID string,
	pageFilter PageFilter,
) ([]*entities.Room, error) {
	queryFilter := s.makeGetAllQueryFilter(search, userID)

	rooms, err := s.roomRepository.GetByQueryFilters(*queryFilter, &repositories.PageFilter{
		Offset: (pageFilter.Page - 1) * pageFilter.Size,
		Limit:  pageFilter.Size,
	})
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *RoomService) CountAll(
	search string,
	userID string,
) (int64, error) {
	queryFilter := s.makeGetAllQueryFilter(search, userID)

	count, err := s.roomRepository.CountByQueryFilters(*queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *RoomService) makeGetAllQueryFilter(
	search string,
	userID string,
) *repositories.QueryFilter {
	queryFilter := &repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    entities.RoomUserIDField,
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	if search != "" {
		searchConditionGroup := repositories.ConditionGroup{
			Operator: repositories.OrLogicalOperator,
			Conditions: []repositories.Condition{
				{
					Field:    entities.RoomNameField,
					Operator: repositories.LikeComparisonOperator,
					Value:    "%" + search + "%",
				},
				{
					Field:    entities.RoomDescriptionField,
					Operator: repositories.LikeComparisonOperator,
					Value:    "%" + search + "%",
				},
			},
		}
		queryFilter.ConditionGroups = append(
			queryFilter.ConditionGroups,
			searchConditionGroup,
		)
	}

	return queryFilter
}
