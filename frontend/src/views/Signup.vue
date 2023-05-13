<template>
 <div class="flex flex-col">
  <div class="signup flex flex-col bg:white md:bg-gray-100 pt-6 pb-10 w-11/12 md:w-full md:px-36 justify-center items-center">
    <section class="flex flex-col justify-center items-center">
      <router-link to="/"><img class="nav-logo w-20 md:w-32 mb-5" src="@/assets/images/login-logo.svg" alt="image of linkedln logo"></router-link>
      <h2 class="text-2xl md:text-3xl md:mb-8">Make the most of your Professional life</h2>
    </section>
    <section class="flex flex-col items-center justify-center bg-white md:border pt-6 rounded-lg w-auto md:w-2/5 h-full md:h-auto ">
      <form action="" class="w-full ml-8 md:ml-16">
        <p class="text-sm mb-2">Email or phone number</p>
        <input  class="border mb-2 border-gray-700 rounded w-full md:w-5/6 h-8"
        type="email"
        name="email"
        v-model="userInput.email"/>

        <p class="text-sm mb-2">Password (6 or more characters)</p>
        <input class="border mb-4 border-gray-700 rounded w-full md:w-5/6 h-8"
         type="password"
         name="password"
         v-model="userInput.password"/>
      </form>
      <p class="text-center text-xs px-10 mb-3">By clicking Agree & join, you agree to the linkedln <span class="text-blue-700 font-medium">User Agreement</span>,
        <span class="text-blue-700 font-medium">Privacy Policy</span>,
        and <span class="text-blue-700 font-medium">Cookie Policy</span>
      </p>
      <button @click="register(userInput)" class="border w-11/12 md:w-5/6 h-12 rounded-3xl bg-blue-700 text-white text-md font-medium mb-3"
        type="submit">
        Agree & Join
      </button>
      <p class="mb-3 visible-hidden md:visible "> or</p>
      <button @click="signInWithGoogle" class="border border-blue-600 w-11/12 md:w-5/6 h-12 flex items-center justify-center rounded-3xl text-blue-700 text-md font-medium mb-3 hover:bg-blue-100 hover:border-blue-900 hover:border-8" type="submit">
        <svg viewBox="0 0 24 24" width="24px" height="24px" x="0" y="0" preserveAspectRatio="xMinYMin meet" class="third-party-join__google-icon mr-2"><g class="color-icon">
          <path style="fill:#E94435" d="M12.1,5.8c1.6-0.1,3.1,0.5,4.3,1.6l2.6-2.7c-1.9-1.8-4.4-2.7-6.9-2.7c-3.8,0-7.2,2-9,5.3l3,2.4C7.1,7.2,9.5,5.7,12.1,5.8z"></path>
          <path style="fill:#F8BB15" d="M5.8,12c0-0.8,0.1-1.6,0.4-2.3l-3-2.4C2.4,8.7,2,10.4,2,12c0,1.6,0.4,3.3,1.1,4.7l3.1-2.4C5.9,13.6,5.8,12.8,5.8,12z"></path>
          <path style="fill:#34A751" d="M15.8,17.3c-1.2,0.6-2.5,1-3.8,0.9c-2.6,0-4.9-1.5-5.8-3.9l-3.1,2.4C4.9,20,8.3,22.1,12,22c2.5,0.1,4.9-0.8,6.8-2.3L15.8,17.3z"></path>
          <path style="fill:#547DBE" d="M22,12c0-0.7-0.1-1.3-0.2-2H12v4h6.1v0.2c-0.3,1.3-1.1,2.4-2.2,3.1l3,2.4C21,17.7,22.1,14.9,22,12z"></path></g>
        </svg>
      Join with Google
      </button>
      <p class="mb-8 mt-3 text-gray-700" >Already on linkedln? <router-link :to="{name: 'Login', params:{userInfo}}" class="text-blue-700 font-medium hover:underline cursor-pointer" >Sign in</router-link></p>
    </section>
  </div>
  <Copyright/>
 </div>
</template>

<script>
import Copyright from '@/components/Copyright.vue'
import { mapActions } from 'vuex'
// import firebase from 'firebase'
// import db from '@/main'
export default {
  components: {
    Copyright
  },
  data () {
    return {
      userInput: {
        name: '',
        displayName: '',
        email: '',
        password: '',
        photoURL: null
      },
      userInfo: []
    }
  },
  methods: {
    // register (userInput) {
    // this.$store.dispatch('register', userInput)
    // },
    ...mapActions([
      'register'
    ]),
    // signInWithGoogle () {
    //   this.$store.dispatch('signInWithGoogle')
    // }
    ...mapActions([
      'signInWithGoogle'
    ])
  }
}
</script>

<style lang="scss" scoped>
 @import '../assets/scss/index.scss';

</style>
