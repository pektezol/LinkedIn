<template>
  <nav class="feed_navigation flex md:flex justify-between py items-center border bg-white">
    <div class="feed_navigation-left md:flex justify-between items-center  pl-28">
     <router-link to="/"> <img  class="feed_navigation-left-logo" src="@/assets/images/linkedin.png" alt="linkedin logo" srcset=""></router-link>
      <input class="feed_navigation-left--searchbar border ml-3 w-72 rounded-md h-8 bg-gray-100 border-gray-100" type="text" placeholder="Search">
    </div>
    <div class="feed_navigation-right flex text-xs items-center cursor-pointer justify-center text-center ">
      <div class="feed_navigation-right--items  mr-8">
        <img class="transform translate-x-1" src="@/assets/images/nav-home.svg" alt="" srcset="">
        <p>Home</p>
      </div>
      <div class="feed_navigation-right--items mr-8">
        <img class="transform translate-x-6" src="@/assets/images/nav-network.svg" alt="" srcset="">
        <p>My Networks</p>
      </div>
      <div class="feed_navigation-right--items mr-8">
        <img class="transform translate-x-px" src="@/assets/images/nav-jobs.svg" alt="" srcset="">
        <p>Jobs</p>
      </div>
      <div class="feed_navigation-right--items mr-8">
        <img class="transform translate-x-6" src="@/assets/images/nav-messaging.svg" alt="" srcset="">
        <p>Messaging</p>
      </div>
      <div class="feed_navigation-right--items mr-8">
        <img class="transform translate-x-6" src="@/assets/images/nav-notifications.svg" alt="" srcset="">
        <p> Notifications</p>
      </div>
      <div class="feed_navigation-right--items mr-8">
        <div class="flex flex-wrap">
          <div class="w-full text-center">
            <img :src="user.data.photoURL" alt="" srcset=""  ref="btnRef" class="transform translate-x-px rounded-full w-7 h-7 ease-linear transition-all duration-150">
           <p class="flex">
              Me
              <icon @click="togglePopover()" class="language-selector__label-chevron mx-1 lazy-loaded" aria-hidden="true">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" preserveAspectRatio="xMinYMin meet" focusable="false" class="lazy-loaded">
                  <path d="M8.8 10.66L14 5.12a.07.07 0 00-.07-.12H2.07a.07.07 0 00-.07.12l5.2 5.54a1.1 1.1 0 001.6 0z" fill="currentColor"></path>
                </svg>
              </icon>
            </p>
            <div ref="popoverRef" v-bind:class="{'hidden': !popoverShow, 'block': popoverShow}" class="bg-white  border border-gray-100 block z-50 font-normal leading-normal text-sm max-w-xs text-left no-underline break-words rounded-lg">
              <div >
                <div class="text-gray-700 p-6 hover:underline" @click="signOut()">
                  Sign Out
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <span class=" h-12 border border-gray-100 mr-4"></span>
      <div class="feed_navigation-right--items mr-4">
        <img class="transform translate-x-3 " src="@/assets/images/nav-work.svg" alt="" srcset="">
        <p class="flex">
          Work
          <icon class="language-selector__label-chevron mx-1 lazy-loaded" aria-hidden="true">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" preserveAspectRatio="xMinYMin meet" focusable="false" class="lazy-loaded">
              <path d="M8.8 10.66L14 5.12a.07.07 0 00-.07-.12H2.07a.07.07 0 00-.07.12l5.2 5.54a1.1 1.1 0 001.6 0z" fill="currentColor"></path>
            </svg>
          </icon>
        </p>
      </div>
      <p class="mr-24 text-xs text-center text-yellow-800 w-24">Try premium free for 1 month</p>
    </div>
  </nav>
</template>

<script>
import { createPopper } from '@popperjs/core'
export default {
  name: 'bottom-pink-popover',
  data () {
    return {
      popoverShow: false,
      user: {
        data: {
          photoURL: 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200',
          displayName: 'VolkanO6',
          email: 'volkanoztoklu@gmail.com'
        }
      }
    }
  },
  methods: {
    togglePopover: function () {
      if (this.popoverShow) {
        this.popoverShow = false
      } else {
        this.popoverShow = true
        createPopper(this.$refs.btnRef, this.$refs.popoverRef, {
          placement: 'bottom'
        })
      }
    },
    signOut () {}
  }
}
</script>
