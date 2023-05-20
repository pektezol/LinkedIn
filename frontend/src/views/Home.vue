<template>
  <div>
    <div class="container">

      <div class="row">
        <!-- left - start -->
        <div class="col-3">
          <leftSide></leftSide>
        </div>

        <!-- middle - start -->
        <div class="col-6">
          
          <mainUpper></mainUpper>

          <hr>
          <post></post>

        </div>

   
        <!-- Right - Start -->
        <div class="col-3">
          <rightSide></rightSide>
        </div>
      </div>
    </div>
    <message></message>

  </div>
</template>

<script>
import leftSide from "../components/Home/leftSide.vue"
import mainUpper from "../components/Home/mainUpper.vue"
import post from "../components/Home/post.vue"
import rightSide from "../components/Home/rightSide.vue"
import message from "../components/Layout/Message.vue" 
import { useAuthStore } from '../store/auth';

import '@/assets/main.css'

export default {
  components: {
    leftSide,
    mainUpper,
    post,
    rightSide,
    message
  },
  mounted() {
    const authStore = useAuthStore()  
    //token cookie
  },
  data() {
    return {
      fullName: 'Irem Eker',
      headline: 'Software Engineer',
      location: 'Istanbul, Turkey',
      profilePicture: '/path/to/profile-picture.jpg',
      about: 'Passionate software engineer with expertise in web development and a focus on frontend technologies. Skilled in JavaScript, Vue.js, HTML, CSS, and Bootstrap. Strong problem-solving and communication skills.',
      experiences: [
        {
          id: 1,
          jobTitle: 'Frontend Developer',
          company: 'ABC Inc.',
          description: 'Developed and maintained frontend applications using Vue.js, HTML, CSS, and JavaScript. Collaborated with cross-functional teams to deliver high-quality software products.'
        },
        {
          id: 2,
          jobTitle: 'Web Developer',
          company: 'XYZ Corp.',
          description: 'Worked on building responsive web applications using modern web technologies. Implemented user-friendly interfaces and optimized performance.'
        }
      ],
      educations: [
        {
          id: 1,
          degree: 'Bachelor of Science in Computer Science',
          university: 'University of ABC',
          description: 'Studied computer science and gained a solid foundation in software development. Completed coursework in algorithms, data structures, and software engineering principles.'
        },
        {
          id: 2,
          degree: 'Master of Science in Software Engineering',
          university: 'University of XYZ',
          description: 'Specialized in software engineering principles and practices. Conducted research on software design patterns and agile methodologies.'
        }
      ],
    };
  }, 
};
</script>

<style scoped>
* {
  padding: 0px;
  margin: 0px;
  box-sizing: border-box;
}

html,
h1,
h2,
h3,
h4,
h5,
h6 {
  padding: 0px;
  margin: 0px;
  font: 100%;
}

.mobile-message {
  display: none;
}

@media (max-width: 600px) {
  #app {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  * {
    background-color: rgb(61, 114, 211);
  }

  .mobile-message {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-content: center;
    color: white;
    margin-top: 300px;
    font-size: 20px;
    text-align: center;
    margin-left: 20px;
    margin-right: 20px;
    font-family: 'Courier New', Courier, monospace;
    background: blueviolet;
    background-size: cover;
    max-height: 100%;
  }
}
</style>