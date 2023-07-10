import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
// 路由分为三种,常量路由(比如首页),异步路由(),任意路由(比如404)
export const constantRoutes = [
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      //侧边栏显示的文字和图标及面包屑当中显示
      meta: {
        title: '首页',
        icon: 'dashboard'
      }
    }]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
      //用户管理 
  {
    path: '/user',
    component: Layout,
    name: 'User',
    //默认二重
    redirect: '/user/list',
    meta: {
      title: '用户管理',
      icon: 'el-icon-user-solid'
    },
    children: [
      {
        path: 'list',
        component: () =>
          import('@/views/user/list'),
        name: 'User1',
        meta: {
          title: '用户查看'
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
      redirect: '/dish/categoryList',
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
          },
          // children:[
          //   {
          //   path: 'product',
          //   component: ()=>
          //   import('@/views/dish/product.vue'),
          //   name:'Product',
          //   }
          // ]
        },
      ]
    },
    //菜品菜品菜品
    {
      path: '/dish/categoryList/product/',
      // component: Layout,
      component: () => import("@/views/dish/product.vue"),
      name: 'Product',
  },
    //订单管理 
    {
      path: '/order',
      component: Layout,
      name: 'Order',
      //默认二重
      redirect: '/order/list',
      meta: {
        title: '订单管理',
        icon: 'el-icon-s-order'
      },
      children: [
        {
          path: 'list',
          name: "Order1",
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
      redirect: '/deliver/list',
      meta: {
        title: '配送员管理',
        icon: 'el-icon-user'
      },
      children: [
        {
          path: 'list',
          component: () =>
            import('@/views/deliver/list'),
          name: 'Deliver1',
          meta: {
            title: '配送员'
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
      redirect: '/carousel/upload',
      meta: {
        title: '轮播图管理',
        icon: 'el-icon-picture-outline'
      },
      children: [
        {
          path: 'upload',
          component: () =>
            import('@/views/carousel/index'),
          name: 'Carousel1',
          meta: {
            title: '轮播图上传'
          }
        },
      ]
      },
]

//异步路由
export const asyncRoutes = [
  
]

//任意路由
export const anyRoutes = [
  // 404 page must be placed at the end !!!(404页面必须被放置在末尾)
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
