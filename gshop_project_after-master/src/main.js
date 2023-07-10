import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

// import locale from 'element-ui/lib/locale/lang/en' // lang i18n
import locale from 'element-ui/lib/locale/lang/zh-CN'
Vue.use(ElementUI, { locale })

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control

import * as $API from "@/api/index.js";// API接口文档

import hasBtn from "@/utils/permission";//用户按钮权限判断

//三级分类
import CategorySelector from "@/components/CategorySelector"
//自定义提示按钮
import MyButton from "@/components/MyButton"
import Axios from 'axios'
//注册全局组件
Vue.component(CategorySelector.name,CategorySelector);
Vue.component(MyButton.name,MyButton);

/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
if (process.env.NODE_ENV === 'production') {
  const { mockXHR } = require('../mock')
  mockXHR()
}

// set ElementUI lang to EN
Vue.use(ElementUI, { locale })
// 如果想要中文版 element-ui，按如下方式声明
// Vue.use(ElementUI)

Vue.config.productionTip = false

Vue.prototype.$API = $API;//API接口
Vue.prototype.$hasBtn = hasBtn;//判断是否有权限
Axios.interceptors.request.use(
  config => {
    let token = window.localStorage.getItem('token')
    if (token) {
      //将token放到请求头发送给服务器,将tokenkey放在请求头中
      config.headers.Authorization = token
      //也可以这种写法
      // config.headers['accessToken'] = Token;
    }
    return config
  },
  error => {
    // Do something with request error
    return Promise.reject(error)
  }
)
new Vue({
  el: '#app',
  router,
  store,
  created() {
    // 初始化应用程序
    // this.$store.dispatch('initializeApp')
    this.$store.state.token = localStorage.getItem('token')
    this.$store.state.canteen_id = localStorage.getItem('canteen_id')
    
  },
  render: h => h(App)
})
