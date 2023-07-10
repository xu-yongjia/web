import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration
// 免登录白名单
const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
  // 加载进度条
  NProgress.start()
  document.title = '校园外卖系统'

  // 是否已有用户登录
  const hasToken = localStorage.getItem('token')
  if (hasToken!=null) {

    // 有token

    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      // 现在的路由是登录，那么就进入主页
      next()
      NProgress.done()
    } 
    else {
      // 现在的路由不是登录的路径，看食堂编号是否存在
      const hasGetCanteenId = localStorage.getItem('canteen_id')
      if (hasGetCanteenId!=null) {
        // 编号存在，继续路由
        next()
      } 
      else {
        Message.error(error || 'Has Error')
        next(`/login?redirect=${to.path}`)
        NProgress.done()
      }
    }
  } else {
    // 没有token
    if (whiteList.indexOf(to.path) !== -1) {
      // 在白名单里，继续导航
      next()
    } else {
      // 不在白名单里，重新回到登录界面
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
