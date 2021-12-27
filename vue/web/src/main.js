// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import BootstrapVue from 'bootstrap-vue';

import axios from 'axios';

import VCharts from 'v-charts'
Vue.use(VCharts)

Vue.config.productionTip = false

Vue.use(BootstrapVue);
Vue.prototype.$http = axios;
Vue.prototype.$http.defaults.headers.common['Authorization'] = "";

import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

/* eslint-disable no-new */
new Vue({
  
  el: '#app', 
  router,
  components: { App },
  template: '<App/>',
  data(){
    return {
      token: "",
      user: "",
      //backIP: "ec2-3-128-204-27.us-east-2.compute.amazonaws.com:8000",
      backIP: "localhost:8000",
    }
  }
})
