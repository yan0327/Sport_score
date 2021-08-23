package request

import "gin-vue-admin/model"

type SportProvinceSearch struct{
    model.SportProvince
    PageInfo
}