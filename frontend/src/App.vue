<script> 
import Navbar from "./components/Layout/Navbar.vue" 

export default {
  components: {
    'navbar': Navbar
  }
}

</script>

<template> 
  <div class="background1">
    <navbar></navbar>
  <router-view>

  </router-view>

  </div>
</template>

<style>

</style>
