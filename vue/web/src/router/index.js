import Vue from 'vue'
import Router from 'vue-router'
import getwinrate from '@/components/getwinrate'
import handmanager from '@/components/handmanager'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
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
