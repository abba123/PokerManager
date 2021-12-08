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
      component: login,
    },
    {
      path: '/handmanager',
      name: 'handmanager',
      component: handmanager,
      beforeEnter (to, from, next) {
        if (Vue.prototype.$http.defaults.headers.common['Authorization'] == ""){
          next({name: 'login'})
        }
        next()
      }
    },
    {
      path: '/getwinrate',
      name: 'getwinrate',
      component: getwinrate,
      beforeEnter (to, from, next) {
        if (Vue.prototype.$http.defaults.headers.common['Authorization'] == ""){
          next({name: 'login'})
        }
        next()
      }
    }
  ]
})

export default router