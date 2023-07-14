/*
 * @Description: 配置文件
 */
module.exports = {
  publicPath: './',
  devServer: {
    open: true,
    proxy: {
      '/api': {
        target: ' http://2eb7d14.r3.cpolar.top',
        changeOrigin: true, //允许跨域
        pathRewrite: {
          '^/api': '/api/v1'
        }
      }
    }
  }
}