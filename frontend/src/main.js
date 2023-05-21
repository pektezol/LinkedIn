import Vue from "vue";
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue' 
import VueCookies from 'vue-cookies'


import "axios"

import App from "./App.vue";
import router from "./router"; 
import pinia from "./store" 

import "./assets/main.css";
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)
Vue.use(pinia) 
Vue.use(VueCookies, { expires: '30d'})

new Vue({ 
  router,
  pinia, 
  render: (h) => h(App),
}).$mount("#app");
