<template>
  <div>
    <!-- left -->
    <div class="card">
      <div class="card-header text-center"
        style="background-image: :url(); height: 50px;">
        <img
          :src="profile_data.profile_picture"
          class="rounded-circle" style="width: 70px; height: 70px; border: 2px solid white;" />
      </div>
      <div class="card-body mt-4" style="font-family: 'Poppins', sans-serif;">
        <p class="card-title header text-center">
          <a href="/profile" target="_blank" style="color: #212529;">{{ profile_data.first_name }}
             {{ profile_data.last_name }}</a>
        </p>
        <div class="text-center" style="margin-bottom: 10px; color: #5f5f5f;">
          {{ profile_data.headline }}
        </div>
         
         
      </div>
      
    </div>


  </div>
</template>
  
<script>
import axios from "axios";

export default {
  data() {
    return {
      profile_data: null,
      user: {
        data: {
          photoURL: 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200',
          displayName: 'VolkanO6',
          email: 'volkanoztoklu@gmail.com'
        }
      }
    }
  },
  async mounted() {
     await axios.get(`https://software.ardapektezol.com/api/profile`, {
      headers: {
        Authorization: this.$cookies.get("token")
      }
    }).then(res => {
      console.log("left", res.data.data.user);
      this.profile_data = res.data.data.user
    })
  }
}
</script>
  