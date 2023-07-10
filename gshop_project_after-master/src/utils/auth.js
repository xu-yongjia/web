import Cookies from 'js-cookie'
import store from '../store/modules/user.js';
const TokenKey = 'vue_admin_template_token'

export function getToken() {
  // 第一次没有获取到token,就会返回undefined
  return store.state.token
}
