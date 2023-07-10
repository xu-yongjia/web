const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.token,
  avatar: state => state.avatar,
  canteen_id: state => state.canteen_id,
  //最终用户可拥有的路由
  currentAsyncRoutes: state => state.user.currentAsyncRoutes,
}
export default getters
