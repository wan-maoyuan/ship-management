package database

import (
	"fmt"
	"ship-management/pkg/model"
)

func InsertPMSEquipmentFirstCategory(category *model.PMSEquipmentFirstCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentFirstCategory{}).
		Create(category)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func DeletePMSEquipmentFirstCategory(category *model.PMSEquipmentFirstCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentFirstCategory{}).
		Where("id = ?", category.ID).
		Delete(&model.PMSEquipmentFirstCategory{})

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func UpdatePMSEquipmentFirstCategory(category *model.PMSEquipmentFirstCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentFirstCategory{}).
		Save(category)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func QueryAllPMSEquipmentFirstCategory() ([]*model.PMSEquipmentFirstCategory, error) {
	var list []*model.PMSEquipmentFirstCategory

	db := msDB.
		Model(&model.PMSEquipmentFirstCategory{}).
		Find(&list)

	if err := db.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertPMSEquipmentSecondCategory(category *model.PMSEquipmentSecondCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentSecondCategory{}).
		Create(category)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSEquipmentSecondCategoryByCode(firstCode string) ([]*model.PMSEquipmentSecondCategory, error) {
	var list []*model.PMSEquipmentSecondCategory

	db := msDB.
		Model(&model.PMSEquipmentSecondCategory{}).
		Where("first_code = ?", firstCode).
		Find(&list)

	if err := db.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertPMSEquipmentThirdCategory(category *model.PMSEquipmentThirdCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentThirdCategory{}).
		Create(category)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSEquipmentThirdCategoryByCode(firstCode, secondCode string) ([]*model.PMSEquipmentThirdCategory, error) {
	var list []*model.PMSEquipmentThirdCategory

	db := msDB.
		Model(&model.PMSEquipmentThirdCategory{}).
		Where("first_code = ? and second_code = ?", firstCode, secondCode).
		Find(&list)

	if err := db.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertPMSEquipmentFourthCategory(category *model.PMSEquipmentFourthCategory) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	db := msDB.
		Model(&model.PMSEquipmentFourthCategory{}).
		Create(category)

	if err := db.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSEquipmentFourthCategoryByCode(firstCode, secondCode, thirdCode string) ([]*model.PMSEquipmentFourthCategory, error) {
	var list []*model.PMSEquipmentFourthCategory

	db := msDB.
		Model(&model.PMSEquipmentFourthCategory{}).
		Where("first_code = ? and second_code = ? and third_code = ?", firstCode, secondCode, thirdCode).
		Find(&list)

	if err := db.Error; err != nil {
		return nil, err
	}

	return list, nil
}
