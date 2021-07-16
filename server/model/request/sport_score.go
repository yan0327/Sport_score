package request

import "gin-vue-admin/model"

type Sport_scoreSearch struct{
    model.Sport_score
    PageInfo
}