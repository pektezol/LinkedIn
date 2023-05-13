<template>
  <header>
    <Navigation/>
    <div class=" hero relative z-10 flex  flex-wrap  justify-center  md:justify-between ml-2 md:ml-28 ">
      <div class=" hero-right md:-ml-4">
        <div class="hero-right-heading mt-10 md:mt-16">
          <h2 class="hero-right-heading-1 text-3xl md:text-5xl mb-0 md:mb-4 font-sans  md:font-normal text-blue-600 tracking-wide ">Welcome to your </h2>
          <h2 class="hero-right-heading-2 text-3xl md:text-5xl mb-2 md:mb-20 font-sans  md:font-normal text-blue-600 tracking-wide">Professional Community</h2>
        </div>
        <div class="hero-right-cta flex flex-col cursor-pointer justify-between ">
          <div class="hero-right-cta flex justify-between items-center mb-4 border w-full md:w-96 h-16 rounded-lg hover:shadow-lg hover:border-gray-50">
            <p class="hero-right-cta-paragraph text-black text-xl  py-4 pl-4">Search for a job</p>
            <icon aria-hidden="true" class="chevron-right-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="35px" height="30px" class="artdeco-icon lazy-loaded" focusable="false">
              <path d="M14 12L8 3.06 9.55 2l6.19 9.15a1.5 1.5 0 010 1.69L9.55 22 8 20.94z" fill="rgba(0, 0, 0, 0.6)"></path>
              </svg>
            </icon>
          </div>
          <div class="hero-right-cta flex justify-between items-center mb-4 border w-full md:w-96 h-16 rounded-lg hover:shadow-lg hover:border-gray-50">
            <p class="hero-right-cta-paragraph text-black  text-xl  py-4 pl-4">Find a person you know</p>
            <icon aria-hidden="true" class="chevron-right-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="35px" height="30px" class="artdeco-icon lazy-loaded" focusable="false">
              <path d="M14 12L8 3.06 9.55 2l6.19 9.15a1.5 1.5 0 010 1.69L9.55 22 8 20.94z" fill="rgba(0, 0, 0, 0.6)"></path>
              </svg>
            </icon>
          </div>
          <div class="hero-right-cta flex justify-between items-center -mb-4 md:mb-4 border w-full md:w-96 h-16 rounded-lg hover:shadow-lg hover:border-gray-50">
            <p class="hero-right-cta-paragraph text-black  text-xl  py-4 pl-4">Learn a Skill</p>
            <icon aria-hidden="true" class="chevron-right-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="35px" height="30px" class="artdeco-icon lazy-loaded" focusable="false">
              <path d="M14 12L8 3.06 9.55 2l6.19 9.15a1.5 1.5 0 010 1.69L9.55 22 8 20.94z" fill="rgba(0, 0, 0, 0.6)"></path>
              </svg>
            </icon>
          </div>
        </div>
      </div>
      <div class="hero-image">
        <img src="https://static-exp1.licdn.com/sc/h/dxf91zhqd2z6b0bwg85ktm5s4" class="z-10 relative top-12 w-auto " alt="">
      </div>
    </div>
  </header>
</template>

<script>
import Navigation from '@/components/Navigation.vue'
export default {
  name: 'header',
  components: {
    Navigation
  }
}
</script>

<style lang="scss" scoped>
.hero-image{
  width: 57.3%;
}
</style>
