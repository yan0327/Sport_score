<template>
  <div class="inituser_page">
    <div class="inituser_page_panle">
      <div v-if="hello < 2" id="hello" :class="[hello < 1 ? 'slide-in-fwd-top' : 'slide-out-right']" class="hello  ">
        <div>
          <div class="hello_title">学生信息初始化</div>
          <p class="in-two a-fadeinT">初始化须知</p>
          <p class="inituser_p">1.请按照格式要求填写真实信息</p>
          <p class="inituser_p">2.若信息不一致可能会影响到成绩的录入</p>
          <p class="inituser_p">3.请仔细填写</p>
          <p class="inituser_btn">
            <el-button type="primary" @click="showNext">
              我已确认
            </el-button>
          </p>
        </div>
      </div>
      <div v-if="hello > 0 " :class="[(hello > 0 && !out)? 'slide-in-left' : '' , out ? 'slide-out-right' : '']" class=" form">
        
        <el-form :inline="true" :model="form" class="demo-form-inline">
        <el-form-item label="姓名">
            <el-input v-model="form.name" placeholder="请输入姓名" />
          </el-form-item>
          <el-form-item label="性别">
            <el-input v-model="form.sex" placeholder="请输入性别" />
          </el-form-item>
          <el-form-item label="民族">
            <el-input v-model="form.ethnicity" placeholder="请输入民族" />
          </el-form-item>
          <el-form-item label="学校">
            <el-input v-model="form.school" placeholder="请输入学校" />
          </el-form-item>
          <el-form-item label="班级">
            <el-input v-model="form.class" placeholder="请输入班级" />
          </el-form-item>
          <el-form-item label="学号">
            <el-input v-model="form.testid" placeholder="请输入学号" />
          </el-form-item>
           <el-form-item label="政治面貌">
            <el-input v-model="form.politicalstatus" placeholder="请输入政治面貌" />
          </el-form-item>
          <el-form-item label="身份证">
            <el-input v-model="form.idcard" placeholder="请输入身份证" />
          </el-form-item>
          
          <el-form-item label="联系方式">
            <el-input v-model="form.Phonenum1" placeholder="请输入联系方式" />
          </el-form-item>
          <el-form-item label="监护人联系方式">
            <el-input v-model="form.Phonenum2" placeholder="请输入监护人联系方式" />
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">立即初始化</el-button>
            </div>
          </el-form-item>
      </el-form>


        <el-form ref="form" :model="form" label-width="100px">
          
        </el-form>
      </div>
    </div>
  </div>
</template>
<script>
import { initDB } from '@/api/initdb'
import { createInformation } from '@/api/user'
export default {
  name: 'Inituser',
  data() {
    return {
      hello: 0,
      out: false,
      form: {
        class: '',
        ethnicity: '',
        idcard: '',
        name: '',
        phonenum1:'',
        phonenum2:'',
        politicalstatus: '',
        school: '',
        sex:'',
        testid:''
      }
    }
  },
  created() {
    // setInterval(() => {
    //   if (this.hello < 3) {
    //     this.hello = this.hello + 1
    //   }
    // }, 2000)
  },
  methods: {
    showNext() {
      this.hello = this.hello + 1
      console.log(this.hello)
    },
    goDoc() {
      window.open('https://www.gin-vue-admin.com/docs/first_master#3-init')
    },
    async onSubmit() {
      const loading = this.$loading({
        lock: true,
        text: '正在初始化用户信息，请稍候',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      try {
        const res = await createInformation(this.form)
        if (res.code === 0) {
          this.out = true
          this.$message({
            type: 'success',
            message: res.msg
          })
          this.$router.push({ name: 'Login' })
        }
        loading.close()
      } catch (err) {
        this.$message({
            type: 'fail',
            message: res.msg
          })
        loading.close()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.inituser_page{
  margin: 0;
  padding: 0;
  background-image: url("~@/assets/login_background.svg");
  background-size: cover;
  width: 100%;
  height: 100%;
  position: relative;
  .inituser_page_panle{
    position: absolute;
    top: 3vh;
    left: 2vw;
    width: 96vw;
    height: 94vh;
    background-color: rgba(255,255,255,.8);
    backdrop-filter: blur(5px);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    .hello{
      position: absolute;
      z-index: 2;
      text-align: center;
      width: 100%;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      .hello_title{
        font-size: 32px;
        line-height: 98px;
      }
      .in-two{
        font-size: 22px;
      }
      .inituser_p{
        margin-top: 20px;
        color: #777777;
      }
      .inituser_btn{
        margin-top: 20px;
      }
    }
    .form{
      position: absolute;
      z-index: 3;
      margin-top: 60px;
      width: 700px;
      height: auto;
      padding: 20px;
      border-radius: 6px;
    }
  }
}

.slide-in-fwd-top {
  -webkit-animation: slide-in-fwd-top 0.4s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
  animation: slide-in-fwd-top 0.4s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}
.slide-out-right {
  -webkit-animation: slide-out-right 0.5s cubic-bezier(0.550, 0.085, 0.680, 0.530) both;
  animation: slide-out-right 0.5s cubic-bezier(0.550, 0.085, 0.680, 0.530) both;
}
.slide-in-left {
  -webkit-animation: slide-in-left 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
  animation: slide-in-left 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}
@-webkit-keyframes slide-in-fwd-top {
  0% {
    -webkit-transform: translateZ(-1400px) translateY(-800px);
    transform: translateZ(-1400px) translateY(-800px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateZ(0) translateY(0);
    transform: translateZ(0) translateY(0);
    opacity: 1;
  }
}
@keyframes slide-in-fwd-top {
  0% {
    -webkit-transform: translateZ(-1400px) translateY(-800px);
    transform: translateZ(-1400px) translateY(-800px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateZ(0) translateY(0);
    transform: translateZ(0) translateY(0);
    opacity: 1;
  }
}
@-webkit-keyframes slide-out-right {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateX(1000px);
    transform: translateX(1000px);
    opacity: 0;
  }
}
@keyframes slide-out-right {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateX(1000px);
    transform: translateX(1000px);
    opacity: 0;
  }
}
@-webkit-keyframes slide-in-left {
  0% {
    -webkit-transform: translateX(-1000px);
    transform: translateX(-1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
}
@keyframes slide-in-left {
  0% {
    -webkit-transform: translateX(-1000px);
    transform: translateX(-1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
}
@media (max-width: 750px) {
  .form{
    width: 94vw !important;
    padding: 0;
  }
}

</style>
