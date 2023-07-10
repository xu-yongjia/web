<!--
 * @Description: 登录组件
 -->
<template>
  <div id="myLogin">
    <el-dialog title="登录" width="300px" center :visible.sync="isLogin">
      <el-form :model="LoginUser" :rules="rules" status-icon ref="ruleForm" class="demo-ruleForm">
        <el-form-item prop="name">
          <el-input prefix-icon="el-icon-user-solid" placeholder="请输入账号" v-model="LoginUser.name"></el-input>
        </el-form-item>
        <el-form-item prop="pass">
          <el-input
            prefix-icon="el-icon-view"
            type="password"
            placeholder="请输入密码"
            v-model="LoginUser.pass"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="medium" type="primary" @click="Login" style="width:100%;">登录</el-button>
        </el-form-item>
      </el-form>
      <el-link
            type="primary"
            @click="Changelogin"            
            style="float:right;margin-bottom:10px;"            
          >没有账号？请先注册></el-link>
    </el-dialog>
  </div>
</template> 

<script> 
import { mapActions } from "vuex";

export default {
  name: "MyLogin",
  data() {
    // 用户名的校验方法
    let validateName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("请输入用户名"));
      }
      // 用户名长度在5-16之间,允许字母数字下划线
      const userNameRule = /^[a-zA-Z0-9_]{5,16}$/;
      if (userNameRule.test(value)) {
        this.$refs.ruleForm.validateField("checkPass");
        return callback();
      } else {
        return callback(new Error("长度5-16之间,允许字母数字下划线"));
      }
    };
    // 密码的校验方法
    let validatePass = (rule, value, callback) => {
      if (value === "") {
        return callback(new Error("请输入密码"));
      }
      // 密码以字母开头,长度在6-18之间,允许字母数字和下划线
      const passwordRule = /^[a-zA-Z0-9_]{5,16}$/;
      if (passwordRule.test(value)) {
        this.$refs.ruleForm.validateField("checkPass");
        return callback();
      } else {
        return callback(
          new Error("长度6-18之间,允许字母数字和下划线")
        );
      }
    };
    return {
      LoginUser: {
        name: "",
        pass: ""
      },
      // 用户信息校验规则,validator(校验方法),trigger(触发方式),blur为在组件 Input 失去焦点时触发
      rules: {
        name: [{ validator: validateName, trigger: "blur" }],
        pass: [{ validator: validatePass, trigger: "blur" }]
      }
    };
  },
  computed: {
    // 获取vuex中的showLogin，控制登录组件是否显示
    isLogin: {
      get() {
        return this.$store.getters.getShowLogin;
      },
      set(val) {
        this.$refs["ruleForm"].resetFields();
        this.setShowLogin(val);
      }
    }
  },
  methods: {
    ...mapActions(["setUser", "setShowLogin"]),
    Changelogin(){
      this.setShowLogin(false);
      this.$router.push('/register')
    },
    Login() {
      // 通过element自定义表单校验规则，校验用户输入的用户信息
      this.$refs["ruleForm"].validate(valid => {
        //如果通过校验开始登录
        if (valid) {
          this.$axios
            .post("/api/users/login", {
              userName: this.LoginUser.name,
              password: this.LoginUser.pass
            })
            .then(res => {
              if (res.status === 404) {
                this.notifyError('验证失败', res.msg)
              } else if (res.status === 200) {
                // 登录信息存到本地
               let user = JSON.stringify(res.data.data.user)
              //  let user = res.data.data.user
                localStorage.setItem('user', user)
                localStorage.setItem('token', res.data.data.token)

                // 登录信息存到vuex
                this.setUser(res.data.data.user)
                // 弹出通知框提示登录成功信息
                this.notifySucceed('登录成功')
                this.isLogin = false
                this.$router.push({
                  name: 'Home'
                })
              } else {
                this.notifyError('登录失败', res.msg)
              }
            })
            .catch(error => {
              this.notifyError('登录失败', error)
            })
        } else {
          return false;
        }
      });
    }
  }
};
</script>
<style>
  /* 设置“登录”的样式 */
  .el-dialog .dialog-title {
  font-size: 20px; /* 设置字体大小 */
  color: red; /* 设置字体颜色 */
  }  

  .el-button--primary {
  background-color: #005826;
  color: #fff;
  }
  .el-button--primary:hover {
  background-color: #5d8c49;
  border-color: #5d8c49;
  }

  .el-input {
  border-radius: 4px;
  }
</style>