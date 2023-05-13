<template>
  <section class="suggested flex flex-col md:flex-row justify-between z-10 pl-6 py-20 md:py-48 bg-gray-50">
    <div class="suggested-text">
      <h3 class="text-4xl md:text-5xl md:w-3/4 mb-24 md:mb-20 md:-mt-20 md:pl-32  md:leading-snug">Find open jobs and internships</h3>
    </div>
    <div class="suggested-search flex flex-col justify-between md:w-4/6 pr-24">
       <p class="suggested-paragraph mb-2 -mt-20 font-medium leading-9 text-base  md:text-left">SUGGESTED SEARCHES</p>
        <ul class=" suggested-list flex flex-wrap md:flex-wrap items-center">
          <li class="suggested-items text-xl hover:underline cursor-pointer text-gray-600 text-base font-medium border rounded-3xl bg-gray-200 px-2 py-2 md:p-3.5 mr-4 mb-4 w-auto" v-for="(items, index) in jobs" :key="index"   v-show="index < limitBy">{{jobs[index]}}</li>
        </ul>
      <p @click="toggle(initial, length)" class="cursor-pointer">{{ limitBy===10? 'Show more': 'Show less'}}</p>
    </div>
  </section>
</template>

<script>
export default {
  data () {
    return {
      jobs: [
        'Engineering', 'Business Development', 'Finance', 'Administrative Assistant', 'Retail Associate',
        'Customer Service', 'Operations', 'Information Technology', 'Marketing', 'Human Resources',
        'Healthcare Service', 'Sales', 'Program and Project Management', 'Accounting', 'Arts and Design',
        'Community and Social Services', 'Consulting', 'Education', 'Entrepreneurship', 'Legal',
        'Media and Communications', 'Military and Protective Services', 'Product Management', 'Purchasing', 'Quality Assurance', 'Real Estate',
        'Research', 'Support', 'Administrative'
      ],
      initial: 10,
      limitBy: 10
    }
  },
  methods: {
    toggle (initial, length) {
      this.limitBy = (this.limitBy === this.initial) ? length : this.limitBy
    }
  }
}
</script>

<style>
.rounded-3xl{
  border-radius: 3.5rem;
}
.text-xl{
  line-height: 1.5rem;
}
p{
 color: rgba(0,0,0,0.6);
}
.md\:w-8\/12 {
    width: 51.666667%;
}
</style>
