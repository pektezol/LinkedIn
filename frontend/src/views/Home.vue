<template>
  <div class="home">
    <Header/>
    <main>
      <Suggested/>
      <Job/>
      <!-- <Slider/> -->
      <Product/>
      <Marketing/>
      <Join/>
    </main>
    <Footer/>
    <Copyright/>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Suggested from '@/components/Suggested.vue'
import Job from '@/components/Job.vue'
// import Slider from '@/components/Slider.vue'
import Product from '@/components/Product.vue'
import Marketing from '@/components/Marketing'
import Join from '@/components/Join'
import Footer from '@/components/Footer'
import Copyright from '@/components/Copyright'

export default {
  name: 'Home',
  components: {
    Header,
    Suggested,
    Job,
    // Slider,
    Product,
    Marketing,
    Join,
    Footer,
    Copyright
  }
}
</script>

<style lang="scss" scoped>
  @media screen and (max-width: 600px) {
    .home{
      content-visibility: hidden;
      background-color: blue!important;
    }
  }
</style>
