package request

import "gin-vue-admin/model"

type SportSearch struct{
    model.Sport
    PageInfo
}