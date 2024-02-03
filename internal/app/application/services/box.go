package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"log"
	"time"
)

var (
	ErrBoxServiceRoomDoesNotExists                            = errors.New("room does not exists")
	ErrBoxServiceQuantityShouldBeLessOrEqualToBoxItemQuantity = errors.New("quantity should be less than or equal to box item quantity")
)

type BoxService struct {
	boxRepository  repositories.BoxRepository
	itemRepository repositories.ItemRepository
	roomRepository repositories.RoomRepository
}

func NewBoxService(
	boxRepository repositories.BoxRepository,
	itemRepository repositories.ItemRepository,
	roomRepository repositories.RoomRepository,
) *BoxService {
	return &BoxService{
		boxRepository,
		itemRepository,
		roomRepository,
	}
}

func (s *BoxService) Create(name string, description *string, roomID string) (*entities.Box, error) {
	exists, err := s.roomRepository.ExistsByID(roomID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrBoxServiceRoomDoesNotExists
	}

	box, err := entities.NewBox(name, description, roomID)
	if err != nil {
		return nil, err
	}

	err = s.boxRepository.Create(box)
	if err != nil {
		return nil, err
	}

	return box, nil
}

func (s *BoxService) AddItemIntoBox(
	quantity float64,
	boxID string,
	itemID string,
) (*entities.BoxItem, error) {
	item, err := s.itemRepository.GetByID(itemID)
	if item == nil {
		return nil, err
	}

	boxItem, err := s.boxRepository.GetBoxItem(boxID, item.ID)
	if err != nil && !errors.Is(err, repositories.ErrBoxRepositoryBoxItemNotFound) {
		return nil, err
	}

	if err != nil && errors.Is(err, repositories.ErrBoxRepositoryBoxItemNotFound) {
		boxItem, err = entities.NewBoxItem(
			quantity,
			boxID,
			*item,
		)
		if err != nil {
			return nil, err
		}
		err = s.boxRepository.CreateBoxItem(boxItem)
		if err != nil {
			return nil, err
		}
	} else {
		boxItem.Quantity += quantity

		err = s.boxRepository.UpdateBoxItem(boxItem)
		if err != nil {
			return nil, err
		}
	}

	happenedAt := time.Now()

	go func() {
		_, err := s.createAddBoxTransaction(
			quantity,
			boxID,
			*item,
			happenedAt,
		)
		if err != nil {
			log.Println(err)
		}
	}()

	return boxItem, nil
}

func (s *BoxService) createAddBoxTransaction(
	quantity float64,
	boxID string,
	item entities.Item,
	happenedAt time.Time,
) (*entities.BoxTransaction, error) {
	boxTransaction, err := entities.NewAddBoxTransaction(
		quantity,
		boxID,
		item,
		happenedAt,
	)
	if err != nil {
		return nil, err
	}

	err = s.boxRepository.CreateBoxTransaction(boxTransaction)
	if err != nil {
		return nil, err
	}

	return boxTransaction, nil
}

func (s *BoxService) RemoveItemFromBox(
	quantity float64,
	boxID string,
	itemID string,
) error {
	item, err := s.itemRepository.GetByID(itemID)
	if item == nil {
		return err
	}

	boxItem, err := s.boxRepository.GetBoxItem(boxID, item.ID)
	if err != nil {
		return err
	}

	if quantity > boxItem.Quantity {
		return ErrBoxServiceQuantityShouldBeLessOrEqualToBoxItemQuantity
	}

	if quantity == boxItem.Quantity {
		err = s.boxRepository.DeleteBoxItem(boxID, item.ID)
		if err != nil {
			return err
		}
	} else {
		boxItem.Quantity -= quantity

		err = s.boxRepository.UpdateBoxItem(boxItem)
		if err != nil {
			return err
		}
	}

	happenedAt := time.Now()

	go func() {
		_, err := s.createRemoveBoxTransaction(
			quantity,
			boxID,
			*item,
			happenedAt,
		)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (s *BoxService) createRemoveBoxTransaction(
	quantity float64,
	boxID string,
	item entities.Item,
	happenedAt time.Time,
) (*entities.BoxTransaction, error) {
	boxTransaction, err := entities.NewRemoveBoxTransaction(
		quantity,
		boxID,
		item,
		happenedAt,
	)
	if err != nil {
		return nil, err
	}

	err = s.boxRepository.CreateBoxTransaction(boxTransaction)
	if err != nil {
		return nil, err
	}

	return boxTransaction, nil
}

func (s *BoxService) GetAll(
	roomID string,
	userID string,
	search string,
	pageFilter PageFilter,
) ([]*entities.Box, error) {
	queryFilter := s.makeGetAllQueryFilter(search, roomID, userID)

	boxes, err := s.boxRepository.GetByQueryFilters(*queryFilter, &repositories.PageFilter{
		Offset: (pageFilter.Page - 1) * pageFilter.Size,
		Limit:  pageFilter.Size,
	})
	if err != nil {
		return nil, err
	}

	return boxes, nil
}

func (s *BoxService) CountAll(
	userID string,
	search string,
	roomID string,
) (int64, error) {
	queryFilter := s.makeGetAllQueryFilter(search, roomID, userID)

	count, err := s.boxRepository.CountByQueryFilters(*queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *BoxService) makeGetAllQueryFilter(
	search string,
	roomID string,
	userID string,
) *repositories.QueryFilter {
	queryFilter := &repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "rooms.user_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    userID,
					},
				},
			},
		},
	}

	if roomID != "" {
		roomIDConditionGroup := repositories.ConditionGroup{
			Operator: repositories.OrLogicalOperator,
			Conditions: []repositories.Condition{
				{
					Field:    "boxes.room_id",
					Operator: repositories.EqualComparisonOperator,
					Value:    roomID,
				},
			},
		}
		queryFilter.ConditionGroups = append(
			queryFilter.ConditionGroups,
			roomIDConditionGroup,
		)
	}

	if search != "" {
		searchConditionGroup := repositories.ConditionGroup{
			Operator: repositories.OrLogicalOperator,
			Conditions: []repositories.Condition{
				{
					Field:    "boxes.name",
					Operator: repositories.LikeComparisonOperator,
					Value:    "%" + search + "%",
				},
				{
					Field:    "boxes.description",
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
