import Vue from 'vue'
import Router from 'vue-router'
import getwinrate from '@/components/getwinrate'
import handmanager from '@/components/handmanager'
import login from '@/components/login'
Vue.use(Router)

const router  =  new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: login
    },
    {
      path: '/handmanager',
      name: 'handmanager',
      component: handmanager
    },
    {
      path: '/getwinrate',
      name: 'getwinrate',
      component: getwinrate
    }
  ]
})

/*
axios.interceptors.response.use(
  response => {
    // 未登录或会话已过期
    if ('401' === response.data.code) {
      // 重定向到登录页
      router.replace({
        path: '/auth/login',
        query: {redirect: router.currentRoute.fullPath}
      })
    }
    return response;
  },
  error => {
    if (500 === error.response.status) {
      // 服务端异常  
    }
    return Promise.reject(error) // 返回接口返回的错误信息
  }
);
*/

export default router