// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false

import axios from 'axios';

Vue.prototype.$http = axios;
Vue.prototype.$http.defaults.headers.common['Authorization'] = ""

/* eslint-disable no-new */
new Vue({
  
  el: '#app', 
  router,
  components: { App },
  template: '<App/>',
  data(){
    return {
      token: "",
      backIP: "3.133.150.55",
      //backIP: "127.0.0.1",
    }
  }
})
