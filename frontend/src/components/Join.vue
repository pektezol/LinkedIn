<template>
  <section class="join-cta">
    <div class="md:z-10 transform md:ml-5 md:translate-x-16  md:translate-y-16 md:relative ">
     <h2 class="join-cta-heading text-6xl text-blue-600 leading-tight font-medium w-10/12">Join your colleagues, classmates, and friends on LinkedIn</h2>
     <router-link to="https://www.linkedin.com/talent/post-a-job?trk=homepage-basic_talent-finder-cta"><p class="join-cta-btn bg-blue-600 mb-10 md:mt-10 md:mr-96 rounded-3xl w-36 cursor-pointer py-4 px-2 text-center text-white font-medium text-xl hover:underline  ">Get Started</p></router-link>
    </div>
  </section>
</template>

<script>
export default {

}
</script>

<style lang="scss" scoped>
 @import '../assets/scss/index.scss';
 .join-cta{
   background: url('https://static-exp1.licdn.com/sc/h/2vt8plqbv2l2pi6kxm89bqs59') repeat-x bottom/auto 100%;
   width: 100%;
   height: 840px;
   overflow: hidden;
 }
</style>
