<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="班级:">
    <el-input v-model="formData.class" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="民族:">
    <el-input v-model="formData.ethnicity" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="身份证:">
    <el-input v-model="formData.idcard" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="姓名:">
    <el-input v-model="formData.name" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="联系方式:">
    <el-input v-model="formData.phonenum1" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="监护人联系方式:">
    <el-input v-model="formData.phonenum2" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="政治面貌:">
    <el-input v-model="formData.politicalstatus" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="学校:">
    <el-input v-model="formData.school" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="性别:">
    <el-input v-model="formData.sex" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="学校ID:">
    <el-input v-model.number="formData.sportSchoolID" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="学号:">
    <el-input v-model="formData.testid" clearable placeholder="请输入" />
    </el-form-item>
    <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createSportStudent,
  updateSportStudent,
  findSportStudent
} from '@/api/sportStudent' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SportStudent',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            class: '',
            ethnicity: '',
            idcard: '',
            name: '',
            phonenum1: '',
            phonenum2: '',
            politicalstatus: '',
            school: '',
            sex: '',
            sportSchoolID: 0,
            testid: '',
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSportStudent({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resportStudent
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSportStudent(this.formData)
          break
        case 'update':
          res = await updateSportStudent(this.formData)
          break
        default:
          res = await createSportStudent(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>