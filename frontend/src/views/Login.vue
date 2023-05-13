<template>
  <div class="flex flex-col justify-between items-center">
    <div class="signup flex flex-col bg:white md:pt-32 md:pb-10 md:mb-40 w-11/12 md:w-full md:px-36 justify-between items-center">
      <section class="flex flex-col justify-center items-center">
       <router-link to="/"> <img class="nav-logo w-20 md:w-20 mb-3" src="@/assets/images/login-logo.svg" alt="image of linkedln logo"></router-link>
        <h2 class="text-2xl md:text-lg font-medium md:mb-4">Welcome Back</h2>
        <p class="text-md md:text-lg font-normal md:mb-4">Don't miss your next opportunity.Sign in to stay updated on your professional world.</p>
      </section>
      <section class="flex flex-col items-center justify-center bg-white md:border pt-6  w-auto md:w-2/5 h-full md:h-auto ">
        <button @click="signInWithGoogle" class="border border-blue-600 w-11/12 md:w-5/6 h-12 flex items-center justify-center
         rounded-3xl text-blue-700 text-md font-medium mb-3 cursor-pointer focus:outline-none
         hover:bg-blue-100 hover:border-blue-900 hover:border-8" type="submit">
        <svg viewBox="0 0 24 24" width="24px" height="24px" x="0" y="0" preserveAspectRatio="xMinYMin meet" class="third-party-join__google-icon mr-2"><g class="color-icon">
          <path style="fill:#E94435" d="M12.1,5.8c1.6-0.1,3.1,0.5,4.3,1.6l2.6-2.7c-1.9-1.8-4.4-2.7-6.9-2.7c-3.8,0-7.2,2-9,5.3l3,2.4C7.1,7.2,9.5,5.7,12.1,5.8z"></path>
          <path style="fill:#F8BB15" d="M5.8,12c0-0.8,0.1-1.6,0.4-2.3l-3-2.4C2.4,8.7,2,10.4,2,12c0,1.6,0.4,3.3,1.1,4.7l3.1-2.4C5.9,13.6,5.8,12.8,5.8,12z"></path>
          <path style="fill:#34A751" d="M15.8,17.3c-1.2,0.6-2.5,1-3.8,0.9c-2.6,0-4.9-1.5-5.8-3.9l-3.1,2.4C4.9,20,8.3,22.1,12,22c2.5,0.1,4.9-0.8,6.8-2.3L15.8,17.3z"></path>
          <path style="fill:#547DBE" d="M22,12c0-0.7-0.1-1.3-0.2-2H12v4h6.1v0.2c-0.3,1.3-1.1,2.4-2.2,3.1l3,2.4C21,17.7,22.1,14.9,22,12z"></path></g>
        </svg>
        Sign in with Google
        </button>
         <p class="mb-8 mt-3 text-gray-700" >New to linkedln? <router-link to="/Signup" class="text-blue-700 font-medium hover:underline cursor-pointer" >Join now</router-link></p>
      </section>
    </div>
    <Copyright/>
  </div>
</template>

<script>
import Copyright from '@/components/Copyright.vue'
import { mapActions } from 'vuex'
// import firebase from 'firebase'

export default {
  components: {
    Copyright
  },
  methods: {
    ...mapActions([
      'signInWithGoogle'
    ])
  }
}
</script>

<style>
