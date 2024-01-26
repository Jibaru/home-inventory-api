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
