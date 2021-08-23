import service from '@/utils/request'

// @Tags SportStudent
// @Summary 创建SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "创建SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudent/createSportStudent [post]
export const createSportStudent = (data) => {
  return service({
    url: '/sportStudent/createSportStudent',
    method: 'post',
    data
  })
}

// @Tags SportStudent
// @Summary 删除SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "删除SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudent/deleteSportStudent [delete]
export const deleteSportStudent = (data) => {
  return service({
    url: '/sportStudent/deleteSportStudent',
    method: 'delete',
    data
  })
}

// @Tags SportStudent
// @Summary 删除SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudent/deleteSportStudent [delete]
export const deleteSportStudentByIds = (data) => {
  return service({
    url: '/sportStudent/deleteSportStudentByIds',
    method: 'delete',
    data
  })
}

// @Tags SportStudent
// @Summary 更新SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "更新SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportStudent/updateSportStudent [put]
export const updateSportStudent = (data) => {
  return service({
    url: '/sportStudent/updateSportStudent',
    method: 'put',
    data
  })
}

// @Tags SportStudent
// @Summary 用id查询SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "用id查询SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportStudent/findSportStudent [get]
export const findSportStudent = (params) => {
  return service({
    url: '/sportStudent/findSportStudent',
    method: 'get',
    params
  })
}

// @Tags SportStudent
// @Summary 分页获取SportStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SportStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudent/getSportStudentList [get]
export const getSportStudentList = (params) => {
  return service({
    url: '/sportStudent/getSportStudentList',
    method: 'get',
    params
  })
}