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
            <post :id="post.id" :date="post.date" :user_id="post.user.id" :profile_picture="post.user.profile_picture" :first_name="post.user.first_name" :last_name="post.user.last_name" :headline="post.user.headline" :likes="post.likes" :content_text="post.content.text" :content_image="post.content.image_base64" :comments="post.comments" :liked_status="post.liked_status"></post> 
          </div> 
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
      post_data: null
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