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
          <div v-for="(post) in post_data" :key="post.id">
            <post :id="post.id" :date="post.date" :first_name="post.user.first_name" :last_name="post.user.last_name" :headline="post.user.headline" :likes="post.likes" :content_text="post.content.text" :content_image="post.content.image_base64" :comments="post.comments" ></post> 
          </div>
          <post id="" first_name="" last_name="" ></post> 
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
import axios from "axios";
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
    axios.get(`https://software.ardapektezol.com/api/posts`, {
      headers: {
        Authorization: this.$cookies.get("token")
      }
    }).then(res => {
      console.log("home",res.data.data.posts);
       this.post_data = res.data.data.posts
    })
  },
  data() {
    return {
      post_data: null,
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