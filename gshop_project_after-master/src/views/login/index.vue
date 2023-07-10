<template>
  <div class="login-container">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">登 录</h3>
      </div>

      <el-form-item prop="canteen_id">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          ref="canteen_id"
          v-model="loginForm.canteen_id"
          placeholder="Canteen_id"
          name="canteen_id"
          type="text"
          tabindex="1"
          auto-complete="on"
          class="custom-placeholder"
        />
      </el-form-item>

      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          placeholder="Password"
          name="password"
          tabindex="2"
          auto-complete="on"
          @keyup.enter.native="handleLogin"
          class="custom-placeholder"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
        </span>
      </el-form-item>

      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleLogin">Login</el-button>

    </el-form>
  </div>
</template>

<script>
import { validUsername } from '@/utils/validate'
import Axios from 'axios'
import { Message } from 'element-ui';
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        canteen_id: null,
        password: ''
      },
      loginRules: {
        canteen_id: [{ required: true, trigger: 'blur', }],
        password: [{ required: true, trigger: 'blur', }]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    //显示密码
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    //登录逻辑
    handleLogin() {
      const url = '/api1/v2/admin/login'
      const tmp = parseInt(this.loginForm.canteen_id)
      this.loginForm.canteen_id = tmp
      Axios.post(url,{canteen_id : this.loginForm.canteen_id,password : this.loginForm.password})
      .then(response => {
        // 处理响应数据
        if(response.data.status === 200){
          this.loading=false
          localStorage.setItem('canteen_id',this.loginForm.canteen_id)
          localStorage.setItem('token',response.data.data.token)
          
          this.$store.state.token = response.data.data.token
          
          this.$store.state.canteen_id = this.loginForm.canteen_id
          Message.success('登录成功！')
          // this.$router.replace({ path: this.redirect || '/' })
          this.$router.replace({ name: 'Dashboard' })
        }
        else if(response.data.status === 201){
            Message.error(response.data.msg)
        }
        else if(response.data.status === 404){
            Message.error(response.data.msg)
            this.$store.dispatch(logout)
          }
      })
      .catch(error => {
        // 处理请求错误
        Message.error('登录失败');
      })
    },
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg:#283443;
$light_gray:#fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  //添加背景图
  background-image: url(../../assets/b5.jpg);
  // background-image: url('@/assets/bg.jpg');
  background-size: cover;
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      --webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: black;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
    // color: black;
  }
}
</style>

<style lang="scss" scoped>
$bg:#2d3a4b;
$dark_gray:#889aa4;
$light_gray:#eee;

.login-container {
  min-height: 100%;
  width: 100%;
  // background-color: $bg;
  background-color: black;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    // color: black;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: black;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: #1b1818cc;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    // color: $dark_gray;
    color: black;
    cursor: pointer;
    user-select: none;
  }
}
.custom-placeholder ::placeholder {
  color: rgba(0, 0, 0, 0.311); /* 使用十六进制颜色值 */
}
</style>