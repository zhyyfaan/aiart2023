// const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// })
module.exports = {
  devServer:{
    host:'localhost',
    port:8080,
    proxy:{
      '/api':{
        target:'http://124.221.224.196:9172/algorithm/run',
        changeOrigin:true,
        pathRewrite:{
          '/api':''
        }
      }
    }
  },
  
}