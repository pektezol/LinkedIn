<template>
  <div class="flex flex-col md:flex-row justify-between items-center md:pl-32 md:py-32 product-cta">
    <div class="product-cta-left">
      <img class="md:w-72 mb-12" src="https://static-exp1.licdn.com/sc/h/b1fxwht7hdbeusleja7ciftsj" alt="Illustration of boy and girl on a couch">
      <h3 class="product-cta-heading w-5/6 text-5xl font-normal leading-snug">Connect with people who can help</h3>
      <p class="product-cta-button mt-10 w-60 text-center text-xl font-medium p-4 rounded-3xl border border-black hover:underline hover:bg-gray-300">Find people you know</p>
    </div>
    <div class="product-cta-right">
      <img class="md:w-72 mb-12" src="https://static-exp1.licdn.com/sc/h/dkfub4sc7jgzg3o31flfr91rv" alt="Illustration of boy and girl on a couch">
      <h3 class="product-cta-heading w-3/4 text-5xl font-normal leading-snug">Learn the skills that can help you grow</h3>
      <div class="flex flex-wrap">
        <div class=" sm:w-6/12 md:w-3/5 mt-10 md:h-2/5">
          <div class="align-middle w-full">
            <button class="flex justify-between text-black font-normal text-xl px-6 py-5 rounded border outline-none focus:outline-none mr-1 mb-1 bg-blueGray-700 font-normal px-6 py-2 rounded-lg outline-none focus:outline-none mr-1 mb-1  w-full" type="button" v-on:click="toggleDropdown()" ref="btnDropdownRef">
              Choose a topic to learn about
              <icon aria-hidden="true" class="caret-down-icon learning-cta__caret lazy-loaded"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" class="artdeco-icon lazy-loaded" focusable="false"><path d="M5 9h14l-6.2 6.7c-.2.2-.5.3-.8.3-.3 0-.6-.1-.8-.3L5 9z" fill="rgba(0, 0, 0, 0.6)"></path></svg></icon>
            </button>
            <div v-bind:class="{'hidden': !dropdownPopoverShow, 'block': dropdownPopoverShow}" class="bg-white text-base list-none text-left border rounded shadow-lg overflow-y-scroll h-60" style="min-width:24rem" ref="popoverDropdownRef">
              <a href="#pablo" class="text-sm hover:underline py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                <h5 class="text-base font-medium">Career Development</h5>
                <p class="text-sm">320+ courses</p>
              </a>
              <a href="#pablo" class=" py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                <h5 class="text-base font-medium">Leadership and Management</h5>
                <p class="text-sm">1,200+ courses</p>
              </a>
              <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                <h5 class="text-base font-medium">Mobile Development</h5>
                <p class="text-sm">370+ courses</p>
              </a>
              <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                <h5 class="text-base font-medium">Marketing</h5>
                <p class="text-sm">690+ courses</p>
              </a>
              <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                <h5 class="text-base font-medium">Cloud Computing</h5>
                <p class="text-sm">570+ courses</p>
              </a>
              <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                Something else here
              </a>
              <a href="#pablo" class="text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-blueGray-700">
                Something else here
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { createPopper } from '@popperjs/core'
export default {
  name: 'dropdown',
  data () {
    return {
      dropdownPopoverShow: false
    }
  },
  methods: {
    toggleDropdown: function () {
      if (this.dropdownPopoverShow) {
        this.dropdownPopoverShow = false
      } else {
        this.dropdownPopoverShow = true
        createPopper(this.$refs.btnDropdownRef, this.$refs.popoverDropdownRef, {
          placement: 'bottom'
        })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
 .product-cta{
   &-heading{
    font-size: 3rem;
    font-weight: 400;
    text-align: left;
   }
 }
</style>
