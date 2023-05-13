<template>
  <section class="cta flex flex-col md:flex-row justify-between items-center px-10 md:px-10  py-10 md:py-3 bg-red-50">
    <div class="md:flex-col cta-header text-red-600 md:text-5xl leading-8 tracking-normal w-6/12 mb-10 md:mt-10 md:pl-24">
      <h2> Post a job and find </h2>
      <h2> people you know</h2>
    </div>
    <router-link to="https://www.linkedin.com/talent/post-a-job?trk=homepage-basic_talent-finder-cta"><p class="cta-button mb-10 md:mt-10 md:mr-96 border rounded-3xl w-36 cursor-pointer py-4 px-2 text-center font-medium text-2xl hover:underline hover:bg-gray-300  ">Post a job</p></router-link>
  </section>
</template>

<script>
export default {

}
</script>

<style lang="scss" scoped>
 @import '../assets/scss/index.scss';
 .cta{
   background: $light-gray;
   &-header{
     color: $light-brown;
   }
   &-button{
     border-color:$light-black;
     margin-right: 35rem;
     width: 11rem;
   }
 }
</style>
