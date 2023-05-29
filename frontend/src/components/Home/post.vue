<template>
  <div class="card mt-2 mb-3">
    <div class="card-header">

      <img src="../../assets/images/profilPhoto.jpeg" class="postphoto">
      <div class="d-inline-flex flex-column ml-1 align-middle">
        <span class="posttext">
          <a href="../../assets/images/profilPhoto.jpeg" target="_blank" style="color: #212529">
            {{ first_name }} {{ last_name }}
          </a>
        </span>
        <span class="profile-desc"> {{ headline }}</span>
        <span class="profile-desc"> {{ date.split("T")[0] }}
          <i class="fas fa-globe-americas fa-sm"></i>
        </span>
      </div>

      <span class="float-right">
        <i class="fas fa-ellipsis-h fa-md"></i>
      </span>
    </div>
    <div>
      <p class="card-text m-3 card-desc">
        {{ content_text }}
      </p>
      <div>
        <img src="https://miro.medium.com/max/1080/1*gdDWFSvHDt57DS2zsh_0Bg.png" class="card-img-top mb-1" alt="..." />
      </div>
      <div class="ml-3 mb-2">


        <span class="profile-desc">
          <img src="https://static-exp1.licdn.com/sc/h/7fx9nkd7mx8avdpqm5hqcbi97" alt="Heart Icon">
          {{ likes }}
        </span>
      </div>
      <div class="card-footer mt-1  text-center">
        <div class="row">
          <div class="col-2"> <span>
              <b-button class="ref" variant="outline-danger" @click="like_post()">
                <i class="far fa-thumbs-up fa-md" style="font-size: 1.2rem">
                  <span class="ml-2 mediatext">
                    Like
                  </span>
                </i>
              </b-button>
            </span> </div>
          <div class="col-10">
            <!-- Using components -->
            <b-input-group class=" ">
              <b-form-input></b-form-input>
              <b-input-group-append>
                <b-button variant="outline-success">Yorum Yap</b-button>
              </b-input-group-append>
            </b-input-group>
          </div>
        </div>
      </div>
      <div>
        <div v-for="(comment) in comments" :key="comment.id">
          <b-card>
            <b-row>
              <b-col md="5">
                <h4>{{ comment.user.first_name }} {{ comment.user.last_name }}</h4>
              </b-col>
              <b-col md="3" offset-md="4">
                <p>{{ comment.date.split("T")[0] }}</p>
              </b-col>
            </b-row>
            <b-card-text class="ml-2">
              {{ comment.comment }}
            </b-card-text>
          </b-card>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import axios from "axios";
export default {
  name: "post",
  props: ['id', 'first_name', 'last_name', 'headline', 'content_text', 'content_image', 'likes', 'comments', 'date'],
  data() {
    return {
      comment: false
    }
  },
  created() {

  },
  methods: {
    like_post() {
      axios.post(`https://software.ardapektezol.com/api/like/${this.id}`,{} ,{
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log("psot like", res.data.data.posts);
        this.post_data = res.data.data.posts
      })
    },
    comment_post() {
      axios.post(`https://software.ardapektezol.com/api/like/${this.id}`,{} ,{
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log("psot like", res.data.data.posts);
        this.post_data = res.data.data.posts
      })
    },
    delete_comment() {

    }
  }
}
</script>
  