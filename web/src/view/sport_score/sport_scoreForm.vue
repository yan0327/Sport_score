<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="学校:">
    <el-input v-model="formData.school" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="班级:">
    <el-input v-model="formData.class" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="考号:">
    <el-input v-model="formData.testid" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="性别:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.sex" clearable ></el-switch>
    </el-form-item>
    
      <el-form-item label="总分:">
    
      <el-input-number v-model="formData.total_score" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="过程评价分:">
    
      <el-input-number v-model="formData.process_evaluation" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="等级:">
    <el-input v-model="formData.grade" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="一类项目:">
    <el-input v-model="formData.item_one" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩1:">
    
      <el-input-number v-model="formData.grade_one" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数1:">
    
      <el-input-number v-model="formData.score_one" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="二类项目:">
    <el-input v-model="formData.item_two" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩2:">
    
      <el-input-number v-model="formData.grade_two" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数2:">
    
      <el-input-number v-model="formData.score_two" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="三类项目:">
    <el-input v-model="formData.item_three" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="成绩3:">
    
      <el-input-number v-model="formData.grade_three" :precision="2" clearable></el-input-number>
    </el-form-item>
    
      <el-form-item label="分数3:">
    
      <el-input-number v-model="formData.score_three" :precision="2" clearable></el-input-number>
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
  createSport_score,
  updateSport_score,
  findSport_score
} from '@/api/sport_score' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Sport_score',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            school: '',
            class: '',
            testid: '',
            sex: false,
            total_score: 0,
            process_evaluation: 0,
            grade: '',
            item_one: '',
            grade_one: 0,
            score_one: 0,
            item_two: '',
            grade_two: 0,
            score_two: 0,
            item_three: '',
            grade_three: 0,
            score_three: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSport_score({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resport_score
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
          res = await createSport_score(this.formData)
          break
        case 'update':
          res = await updateSport_score(this.formData)
          break
        default:
          res = await createSport_score(this.formData)
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