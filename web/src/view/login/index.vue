<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
         <el-select v-model="value" placeholder="请选择">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        </div>
        <el-form
          ref="loginForm"
          :model="loginForm"
          :rules="rules"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入用户名">
              <i slot="suffix" class="el-input__icon el-icon-user" />
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <i
                slot="suffix"
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
              />
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                width="100%"
                height="100%"
                alt="请输入验证码"
                @click="loginVerify()"
              >
            </div>
          </el-form-item>
          <div />
          <el-form-item>
            <el-button
              type="primary"
              style="width: 28%"
              @click="checkInit"
            >老 师 登 录</el-button>
            <el-button
              type="primary"
              style="width: 28%;margin-left:8%"
              @click="submitForm"
            >学 生 登 录</el-button>
            <el-button
              type="primary"
              style="width: 28%;margin-left:8%"
              @click="addUser"
            >学 生 注 册</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panle_right" />
      <div class="login_panle_foot">
        <div class="links">
          <a href="http://doc.henrongyi.top/"><img src="@/assets/docs.png" class="link-icon"></a>
          <a href="https://www.yuque.com/flipped-aurora/"><img src="@/assets/yuque.png" class="link-icon"></a>
          <a href="https://github.com/flipped-aurora/gin-vue-admin"><img src="@/assets/github.png" class="link-icon"></a>
          <a href="https://space.bilibili.com/322210472"><img src="@/assets/video.png" class="link-icon"></a>
        </div>
        <div class="copyright">Copyright &copy; {{ curYear }} 💖 flipped-aurora</div>
      </div>
       <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" append-to-body title="新增用户">
      <el-form ref="userForm" :rules="rules2" :model="userInfo">
        <el-form-item label="用户名" label-width="80px" prop="username">
          <el-input v-model="userInfo.username" />
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item label="别名" label-width="80px" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item label="身份证" label-width="80px" prop="idcard">
          <el-input v-model="userInfo.idcard" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
      </div>
    </el-dialog>
    </div>
  </div>
</template>

<script>
import {
  getUserList,
  setUserAuthority,
  register,
  deleteUser
} from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }
    return {
      curYear: 0,
      lock: 'lock',
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: '4',
        idcard: ''
      },
      loginForm: {
        username: 'admin',
        password: '123456',
        captcha: '',
        captchaId: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
      },
      rules2: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 5, message: '最低5位字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        idcard: [
          { required: true, message: '请选择用户身份证号', trigger: 'blur' },
          { min: 18, message: '最低18位字符', trigger: 'blur' }
        ]
      },
      logVerify: '',
      picPath: '',
       options: [{
          value: '选项1',
          label: '学生登录'
        }, {
          value: '选项2',
          label: '校园管理'
        }, {
          value: '选项3',
          label: '市管理员'
        }, {
          value: '选项4',
          label: '省管理员'
        }, {
          value: '选项5',
          label: '总管理员'
        }],
        value: ''
    }
  },
  created() {
    this.loginVerify()
    this.curYear = new Date().getFullYear()
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async checkInit() {
      const res = await checkDB()
      if (res.code === 0) {
        if (res.data?.needInit) {
          this.$store.commit('user/NeedInit')
          this.$router.push({ name: 'Init' })
        } else {
          this.$message({
            type: 'info',
            message: '已配置数据库信息，无法初始化'
          })
        }
      }
    },
    async login() {
      return await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async(v) => {
        if (v) {
          const flag = await this.login()
          if (!flag) {
            this.loginVerify()
          }
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          this.loginVerify()
          return false
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo)
          if (res.code === 0) {
            this.$message({ type: 'success', message: '创建成功' })
          }
          this.closeAddUserDialog()
        }
      })
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields()
      this.addUserDialog = false
    },
    addUser() {
      this.addUserDialog = true
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath
        this.loginForm.captchaId = ele.data.captchaId
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>
