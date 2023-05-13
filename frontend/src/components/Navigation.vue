<template>
  <nav class="nav flex justify-between items-center mt-4 ml-4">
    <img class="nav-logo md:ml-28 w-20 md:w-32" src="@/assets/images/login-logo.svg" alt="image of linkedln logo">
    <ul class="nav-link md:mr-20 text-base md:text-base font-medium flex justify-between items-center">
     <router-link to="/Signup" ><li class="nav-link-item-1 mr-4 md:mr-12 text-gray-500 hover:bg-gray-200 cursor-pointer hover:border hover:rounded-sm  hover:text-black">Join now</li></router-link>
     <router-link to="/Login" ><li class="nav-link-item-2 mr-4 md:mr-12 p-1 md:p-2 border w-20 md:w-28 border-blue-700 text-blue-700 text-center rounded-3xl cursor-pointer hover:bg-blue-200">Sign in</li></router-link>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'Navigation'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
 .nav{
   &-link{
     &-item{
       &-1{
         &:hover{
           border-radius: 3px;
           transform:none;
         }
       }
     }
   }
 }
</style>
