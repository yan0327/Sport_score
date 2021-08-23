<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="大类:">
    <el-input v-model="formData.class" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="方法名称:">
    <el-input v-model="formData.name" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="分数:">
    <el-input v-model="formData.score" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="学校ID:">
    <el-input v-model.number="formData.sportSchoolID" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="学生ID:">
    <el-input v-model.number="formData.sportStudentID" clearable placeholder="请输入"/>
    </el-form-item>
    
      <el-form-item label="测试数据:">
    <el-input v-model="formData.testdata" clearable placeholder="请输入" />
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
  createSportStudentScore,
  updateSportStudentScore,
  findSportStudentScore
} from '@/api/sportStudentScore' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SportStudentScore',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            class: '',
            name: '',
            score: '',
            sportSchoolID: 0,
            sportStudentID: 0,
            testdata: '',
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSportStudentScore({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resportStudentScore
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
          res = await createSportStudentScore(this.formData)
          break
        case 'update':
          res = await updateSportStudentScore(this.formData)
          break
        default:
          res = await createSportStudentScore(this.formData)
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