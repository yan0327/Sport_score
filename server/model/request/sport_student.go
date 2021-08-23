package request

import "gin-vue-admin/model"

type SportStudentSearch struct{
    model.SportStudent
    PageInfo
}