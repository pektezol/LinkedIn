import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import firebase from 'firebase'

import './assets/scss/index.scss'

var firebaseConfig = {
  apiKey: 'AIzaSyB750Hl0y_S1PTe6ZArvfuC1plyStZApWs',
  authDomain: 'linkedin-clone-83fae.firebaseapp.com',
  projectId: 'linkedin-clone-83fae',
  storageBucket: 'linkedin-clone-83fae.appspot.com',
  messagingSenderId: '674131164733',
  appId: '1:674131164733:web:e8e02fc0b40fb9bd92b454'
}
// Initialize Firebase
firebase.initializeApp(firebaseConfig)
const db = firebase.firestore()
export default db

firebase.auth().onAuthStateChanged(user => {
  store.dispatch('fetchUser', user)
})

Vue.config.productionTip = false

let app = ''
firebase.auth().onAuthStateChanged(user => {
  if (!app) {
    app = new Vue({
      router,
      store,
      render: h => h(App)
    }).$mount('#app')
  }
})
