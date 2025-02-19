package tools

import (
	"os"
	"strconv"

	"github.com/dapidsurya/cakap-aws-internship/ms-user/model/dto"
)

func GetProductByUserId(userId int) ([]dto.ProductDto, error) {
	var products []dto.ProductDto
	baseUrl := os.Getenv("MS_PRODUCT")
	url := baseUrl + "/product/user/" + strconv.Itoa(userId)

	err := get(url, &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}
