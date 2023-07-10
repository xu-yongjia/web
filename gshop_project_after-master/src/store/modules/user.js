// import { login, logout, getInfo } from '@/api/user'
import {login,logout,getInfo} from "@/api/acl/user"
// 通过cookie存储用户token
import { getToken, setToken, removeToken } from '@/utils/auth'
import router, { resetRouter } from '@/router'
//3种路由全选
import {constantRoutes,asyncRoutes,anyRoutes} from "@/router"
//深克隆依赖
import cloneDeep from "lodash/cloneDeep"
import axios from "axios"
import { Message } from 'element-ui'
import Layout from '@/layout'
const getDefaultState = () => {
  return {
    //第一次从token当中读取
    token: '',
    //食堂编号
    canteen_id: 1,
    //头像
    avatar: '',
    //当前登录用户可具有的路由全选
    currentAsyncRoutes:[
      //用户管理 
      {
        path: '/user',
        component: Layout,
        name: 'User',
        //默认二重
        redirect: '/list',
        meta: {
          title: '用户管理',
          icon: 'el-icon-user-solid'
        },
        children: [
          {
            path: 'list',
            component: () =>
              import('@/views/user/list'),
            name: 'User',
            meta: {
              title: '用户管理'
            }
          },
        ]
      },
      //菜品种类管理 
    {
      path: '/dish',
      component: Layout,
      name: 'Dish',
      //默认二重
      redirect: '/categoryList',
      meta: {
        title: '菜品管理',
        icon: 'el-icon-dish'
      },
      children: [
        {
          path: 'categoryList',
          component: () =>
            import('@/views/dish/category.vue'),
          name: 'Category',
          meta: {
            title: '种类'
          }
        },
      ]
    },
    //订单管理 
    {
        path: '/order',
        component: Layout,
        name: 'Order',
        //默认二重
        redirect: '/list',
        meta: {
          title: '订单管理',
          icon: 'el-icon-s-order'
        },
        children: [
          {
            path: 'list',
            name: "Order",
            component: () => import("@/views/order/list"),
            meta: {
              title: "订单"
            }
          },
        ]
    },
    //配送员管理 
    {
      path: '/deliver',
      component: Layout,
      name: 'Deliver',
      //默认二重
      redirect: '/list',
      meta: {
        title: '配送员',
        icon: 'el-icon-user'
      },
      children: [
        {
          path: 'list',
          component: () =>
            import('@/views/deliver/list'),
          name: 'Deliver',
          meta: {
            title: '配送员管理'
          }
        },
      ]
    },
    //轮播图管理 
    {
      path: '/carousel',
      component: Layout,
      name: 'Carousel',
      //默认二重
      redirect: '/upload',
      meta: {
        title: '轮播图管理',
        icon: 'el-icon-picture-outline'
      },
      children: [
        {
          path: 'upload',
          component: () =>
            import('@/views/carousel/index'),
          name: 'Carousel',
          meta: {
            title: '轮播图管理'
          }
        },
      ]
      },
    ],
  }
}

const state = getDefaultState()

const mutations = {
  setToken(state, token,canteen_id) {
    state.token = token;
    state.canteen_id = canteen_id
  },
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },

}

/**
 * 
 * @param {array} all 所有的异步路由信息
 * @param {array} selfRoutes 服务器获取到的当前账户所具有的异步路由信息
 */

const actions = {
  initializeApp({ commit }) {
    // 从 localStorage 中获取 token
    const token = localStorage.getItem('token');
    const canteen_id = localStorage.getItem('canteen_id');
    // 将 token 存储到状态中
    commit('setToken', token,canteen_id);
  },
  logout(){
    
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

