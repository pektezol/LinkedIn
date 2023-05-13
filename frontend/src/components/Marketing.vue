<template>
  <section class="marketing flex items-center md:py-5 md:mb-10">
    <div  class="marketing-media rounded md:px-10 md:z-10 md:left-10  md:relative  lazy-loaded">
      <iframe class="rounded-lg" width="600" height="400" title="Let’s Step Forward | LinkedIn"
        src="https://www.youtube-nocookie.com/embed/EEikRQ58NwM?rel=0">
      </iframe>
    </div>
    <div class="marketing-copy  md:py-72">
      <h2 class="marketing-copy-heading md:text-4xl md:mb-4 md:ml-10">Let's step forward</h2>
      <h3 class="marketing-copy-paragraph md:text-4xl text-black md:font-extralight md:text-left md:w-10/12 md:ml-10 md:leading-normal">With our communities by our side, there's no telling where our next small steps could lead</h3>
    </div>
  </section>
</template>

<script>
export default {

}
</script>

<style lang="scss" scoped>
@import '../assets/scss/index.scss';
.marketing{
  &::after{
    content: '';
    width: 69.1vw;
    height: 720px;
    left: 30%;
    position: absolute;
    z-index: -1;
    background-color: #F4FAFF;
  }
  &-copy{
    // background-color: $light-blue;
    &-heading{
      color:$light-brown;
    }
  }
}
</style>
