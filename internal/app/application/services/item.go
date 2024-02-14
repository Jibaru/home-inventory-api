package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/repositories"
	"log"
	"os"
)

type ItemService struct {
	itemRepository        repositories.ItemRepository
	itemKeywordRepository repositories.ItemKeywordRepository
	assetService          AssetServiceInterface
}

func NewItemService(
	itemRepository repositories.ItemRepository,
	itemKeywordRepository repositories.ItemKeywordRepository,
	assetService AssetServiceInterface,
) *ItemService {
	return &ItemService{
		itemRepository,
		itemKeywordRepository,
		assetService,
	}
}

func (s *ItemService) Create(
	sku string,
	name string,
	description *string,
	unit string,
	userID string,
	keywords []string,
	imageFile *os.File,
) (*entities.Item, error) {
	item, err := entities.NewItem(sku, name, description, unit, userID)
	if err != nil {
		return nil, err
	}

	var itemKeywords []*entities.ItemKeyword
	for _, keyword := range keywords {
		itemKeyword, err := entities.NewItemKeyword(item.ID, keyword)
		if err != nil {
			return nil, err
		}

		itemKeywords = append(itemKeywords, itemKeyword)
	}

	asset, err := s.assetService.CreateFromFile(imageFile, item)
	if err != nil {
		return nil, err
	}

	err = s.itemRepository.Create(item)
	if err != nil {
		go func() {
			err := s.assetService.Delete(asset)
			if err != nil {
				log.Println(err)
			}
		}()
		return nil, err
	}

	err = s.itemKeywordRepository.CreateMany(itemKeywords)
	if err != nil {
		go func() {
			err := s.assetService.Delete(asset)
			if err != nil {
				log.Println(err)
			}
		}()
		return nil, err
	}

	return item, nil
}

func (s *ItemService) GetAll(
	search string,
	userID string,
	pageFilter PageFilter,
) ([]struct {
	Item   *entities.Item
	Assets []*entities.Asset
}, error) {
	queryFilter := s.makeGetAllQueryFilter(search, userID)

	items, err := s.itemRepository.GetByQueryFilters(*queryFilter, &repositories.PageFilter{
		Offset: (pageFilter.Page - 1) * pageFilter.Size,
		Limit:  pageFilter.Size,
	})
	if err != nil {
		return nil, err
	}

	var entitySlice []entities.Entity
	for i := range items {
		entitySlice = append(entitySlice, items[i])
	}
	assets, err := s.assetService.GetByEntities(entitySlice)
	if err != nil {
		return nil, err
	}

	assetsByID := make(map[string][]*entities.Asset)
	for i := range assets {
		assetsByID[assets[i].EntityID] = append(assetsByID[assets[i].EntityID], assets[i])
	}

	output := make([]struct {
		Item   *entities.Item
		Assets []*entities.Asset
	}, 0)
	for i := range items {
		output = append(output, struct {
			Item   *entities.Item
			Assets []*entities.Asset
		}{
			Item:   items[i],
			Assets: assetsByID[items[i].EntityID()],
		})
	}

	return output, nil
}

func (s *ItemService) CountAll(
	search string,
	userID string,
) (int64, error) {
	queryFilter := s.makeGetAllQueryFilter(search, userID)

	count, err := s.itemRepository.CountByQueryFilters(*queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *ItemService) makeGetAllQueryFilter(
	search string,
	userID string,
) *repositories.QueryFilter {
	queryFilter := &repositories.QueryFilter{
		ConditionGroups: []repositories.ConditionGroup{
			{
				Operator: repositories.AndLogicalOperator,
				Conditions: []repositories.Condition{
					{
						Field:    "items.user_id",
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
					Field:    "items.name",
					Operator: repositories.LikeComparisonOperator,
					Value:    "%" + search + "%",
				},
				{
					Field:    "items.description",
					Operator: repositories.LikeComparisonOperator,
					Value:    "%" + search + "%",
				},
				{
					Field:    "item_keywords.value",
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

func (s *ItemService) Update(
	id string,
	name string,
	sku string,
	description *string,
	unit string,
	keywords []string,
	imageFile *os.File,
) (*entities.Item, error) {
	item, err := s.itemRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	err = item.Update(sku, name, description, unit)
	if err != nil {
		return nil, err
	}

	err = s.itemRepository.Update(item)
	if err != nil {
		return nil, err
	}

	err = s.itemKeywordRepository.DeleteByItemID(item.ID)
	if err != nil {
		return nil, err
	}

	if len(keywords) > 0 {
		var itemKeywords []*entities.ItemKeyword
		for _, keyword := range keywords {
			itemKeyword, err := entities.NewItemKeyword(item.ID, keyword)
			if err != nil {
				return nil, err
			}

			itemKeywords = append(itemKeywords, itemKeyword)
		}

		err = s.itemKeywordRepository.CreateMany(itemKeywords)
		if err != nil {
			return nil, err
		}

		item.Keywords = itemKeywords
	}

	if imageFile != nil {
		_, err = s.assetService.UpdateByEntity(item, imageFile)
		if err != nil {
			return nil, err
		}
	}

	return item, nil
}
