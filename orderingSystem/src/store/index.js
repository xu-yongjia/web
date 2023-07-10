/*
 * @Description: Vuex入口
 */
import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
import user from './modules/user'
import shoppingCart from './modules/shoppingCart'



export default new Vuex.Store({
  strict: true,
  modules: {
    user,
    shoppingCart
  }
})
