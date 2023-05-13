<template>
  <footer class="footer flex justify-evenly w-full md:py-7 md:px-32">
    <div class="footer-logo items-start w-20">
      <img src="@/assets/images/login-logo.svg" alt="image of linkedln logo">
    </div>
    <div class="footer-items--1">
      <p class ="font-medium text-black md:mb-4">General</p>
      <ul>
        <li class="md:mb-1 cursor-pointer hover:underline hover:text-blue-700 text-sm text-gray-600 font-medium " v-for="(items, index) in General" :key="index">{{General[index]}}</li>
      </ul>
    </div>
    <div class="footer-items--2">
      <p class ="font-medium text-black  md:mb-4">Browse linkedln</p>
      <ul>
        <li class="md:mb-1 cursor-pointer hover:underline hover:text-blue-700 text-sm text-gray-600 font-medium " v-for="(items, index) in Browse" :key="index">{{Browse[index]}}</li>
      </ul>
    </div>
    <div class="footer-items--3">
      <p class ="font-medium text-black  md:mb-4">Business Solutions</p>
      <ul>
        <li class="md:mb-1 cursor-pointer hover:underline hover:text-blue-700 text-sm text-gray-600 font-medium "  v-for="(items, index) in Solutions" :key="index">{{Solutions[index]}}</li>
      </ul>
    </div>
    <div class="footer-items--4">
      <p class ="font-medium text-black  md:mb-4">Directories</p>
      <ul>
        <li class="md:mb-1 cursor-pointer hover:underline hover:text-blue-700 text-sm text-gray-600 font-medium " v-for="(items, index) in Directories" :key="index">{{Directories[index]}}</li>
      </ul>
    </div>
  </footer>
</template>

<script>
export default {
  data () {
    return {
      General: ['Sign up', 'Help Center', 'About', 'Press', 'Blog', 'Careers', 'Developers'],
      Browse: ['Learning', 'Jobs', 'Salary', 'Mobile', 'Services'],
      Solutions: ['Talent', 'Marketing', 'Sales', 'Learning'],
      Directories: ['Members', 'Jobs', 'Companies', 'Salaries', 'Featured', 'Learning', 'Posts', 'Articles', 'School', 'News', 'News Letter', 'Services', 'Interview Prep']
    }
  }

}
</script>

<style lang="scss" scoped>
@import '../assets/scss/index.scss';
.footer{
  background: $light-gray;
}
</style>
