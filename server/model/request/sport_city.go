package request

import "gin-vue-admin/model"

type SportCitySearch struct{
    model.SportCity
    PageInfo
}