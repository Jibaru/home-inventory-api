package services

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"strconv"
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
	userRepository repositories.UserRepository
	eventBus       services.EventBus
	mailSender     services.MailSender
}

func NewBoxService(
	boxRepository repositories.BoxRepository,
	itemRepository repositories.ItemRepository,
	roomRepository repositories.RoomRepository,
	userRepository repositories.UserRepository,
	eventBus services.EventBus,
	mailSender services.MailSender,
) *BoxService {
	return &BoxService{
		boxRepository,
		itemRepository,
		roomRepository,
		userRepository,
		eventBus,
		mailSender,
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

	err = s.eventBus.Publish(services.BoxItemAddedEvent{
		Quantity:   quantity,
		BoxID:      boxID,
		Item:       *item,
		HappenedAt: happenedAt,
	})
	if err != nil {
		return nil, err
	}

	return boxItem, nil
}

func (s *BoxService) CreateAddBoxTransaction(
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

	err = s.eventBus.Publish(services.BoxItemRemovedEvent{
		Quantity:   quantity,
		BoxID:      boxID,
		Item:       *item,
		HappenedAt: happenedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxService) CreateRemoveBoxTransaction(
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

func (s *BoxService) TransferItem(
	fromBoxID string,
	toBoxID string,
	itemID string,
) error {
	fromBoxItem, err := s.boxRepository.GetBoxItem(fromBoxID, itemID)
	if err != nil {
		return err
	}

	quantity := fromBoxItem.Quantity

	err = s.RemoveItemFromBox(
		quantity,
		fromBoxID,
		itemID,
	)
	if err != nil {
		return err
	}

	_, err = s.AddItemIntoBox(
		quantity,
		toBoxID,
		itemID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxService) DeleteWithTransactionsAndItemQuantities(boxID string) error {
	err := s.boxRepository.DeleteBoxTransactionsByBoxID(boxID)
	if err != nil {
		return err
	}

	err = s.boxRepository.DeleteBoxItemsByBoxID(boxID)
	if err != nil {
		return err
	}

	err = s.boxRepository.Delete(boxID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxService) Update(
	boxID string,
	name string,
	description *string,
) (*entities.Box, error) {
	box, err := s.boxRepository.GetByID(boxID)
	if err != nil {
		return nil, err
	}

	err = box.Update(name, description)
	if err != nil {
		return nil, err
	}

	err = s.boxRepository.Update(box)
	if err != nil {
		return nil, err
	}

	return box, nil
}

func (s *BoxService) TransferToRoom(
	boxID string,
	roomID string,
) error {
	box, err := s.boxRepository.GetByID(boxID)
	if err != nil {
		return err
	}

	err = box.ChangeRoomID(roomID)
	if err != nil {
		return err
	}

	err = s.boxRepository.Update(box)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxService) GetBoxTransactions(
	boxID string,
	pageFilter PageFilter,
) ([]*entities.BoxTransaction, error) {
	queryFilter := s.makeGetBoxTransactionsQueryFilter(boxID)

	boxTransactions, err := s.boxRepository.GetBoxTransactionsByQueryFilters(
		*queryFilter,
		&repositories.PageFilter{
			Offset: (pageFilter.Page - 1) * pageFilter.Size,
			Limit:  pageFilter.Size,
		},
	)
	if err != nil {
		return nil, err
	}

	return boxTransactions, nil
}

func (s *BoxService) CountBoxTransactions(
	boxID string,
) (int64, error) {
	queryFilter := s.makeGetBoxTransactionsQueryFilter(boxID)

	count, err := s.boxRepository.CountBoxTransactionsByQueryFilters(*queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *BoxService) makeGetBoxTransactionsQueryFilter(
	boxID string,
) *repositories.QueryFilter {
	return &repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "box_id",
						Operator: repositories.EqualComparisonOperator,
						Value:    boxID,
					},
				},
			},
		},
	}
}

func (s *BoxService) NotifyBoxItemAdded(
	quantity float64,
	boxID string,
	item entities.Item,
	happenedAt time.Time,
) error {
	user, err := s.userRepository.GetUserByBoxID(boxID)
	if err != nil {
		return err
	}

	quantityStr := strconv.FormatFloat(quantity, 'f', -1, 64)

	body := quantityStr + " " + item.Name + " added into box " + boxID + " at " + happenedAt.String()

	err = s.mailSender.SendMail(
		user.Email,
		"Item added into box "+boxID,
		body,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxService) NotifyBoxItemRemoved(
	quantity float64,
	boxID string,
	item entities.Item,
	happenedAt time.Time,
) error {
	user, err := s.userRepository.GetUserByBoxID(boxID)
	if err != nil {
		return err
	}

	quantityStr := strconv.FormatFloat(quantity, 'f', -1, 64)

	body := quantityStr + " " + item.Name + " removed from box " + boxID + " at " + happenedAt.String()

	err = s.mailSender.SendMail(
		user.Email,
		"Item removed from box "+boxID,
		body,
	)
	if err != nil {
		return err
	}

	return nil
}
