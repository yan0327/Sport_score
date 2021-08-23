package request

import "gin-vue-admin/model"

type SportSchoolSearch struct{
    model.SportSchool
    PageInfo
}