<template>
  <div class="container">
    <div class="card mt-4">
      <div class="card-body">
        <div>
          <div class="header-photo">
            <img alt="Header Photo"
              src="https://64.media.tumblr.com/682b7be9273636dffb1d8fbe3220628b/tumblr_pdz53tIneb1sx8ybdo10_1280.png" />
          </div>
          <div class="profile-photo ml-5">
            <img class="img-fluid rounded-circle mb-1" src="../assets/images/profilPhoto.jpeg" alt="Profile Image" />
          </div>
          <b-card>
            <b-card-body>
              <div class="row">
                <div class="col">
                  <h5 class="card-title">{{ profile_data.first_name }} {{ profile_data.last_name }}</h5>
                  <h6>{{ profile_data.headline }}</h6>
                  <h6 class="text-primary">Connection Count: {{ connection_count }}</h6>
                </div>
                <div class="col">
                  <ul class="list-unstyled">
                    <li><i class="fa fa-envelope-o mr-2"></i>Email: {{ profile_data.email }}</li>
                    <li><i class="fa fa-map-marker mr-2"></i>Location: {{ profile_data.location }}</li>
                    <li><i class="fa fa-map-marker mr-2"></i>Date of Birth: {{ profile_data.date_of_birth.split("T")[0] }}
                    </li>
                    <li><i class="fa fa-map-marker mr-2"></i>{{ profile_data.location }}</li>
                  </ul>
                </div>
              </div>
              <div class="button-group">
                <b-button variant="primary" @click="post_connection(profile_data.user_name)">Connect</b-button>
                <b-button variant="primary">Message</b-button>
              </div>
            </b-card-body>
          </b-card>
        </div>
      </div>
    </div>
    <div>
      <div class="card mt-4">
        <div class="card-body-about">
          <!-- About Section -->
          <h4 class="card-title">About</h4>
          <p>{{ profile_data.bio }}</p>
        </div>
      </div>

    </div>
    <div class="card mt-4">
      <div class="card-body">
        <!-- Experience Section -->
        <div class="row">
          <div class="col-11">
            <h4 class="card-title ">Experiences</h4>
          </div>
          <div class="col ml-4">
             
          </div>
        </div>
        <ul class="list-unstyled">
          <li v-for="experience in experiences" :key="experience.id">
            <div class="border-bottom">
              <div class="row">
                <div class="col-1">
                  <img
                    src="https://media.licdn.com/dms/image/C560BAQFWRMIzQg82hg/company-logo_100_100/0/1678953885463?e=1694044800&v=beta&t=-i7zCz5l88oVKU0LuL7EMpj6IYd8-ovAkdJ_BYKAvDo"
                    style="height: 60px;" class="rounded" alt="...">
                </div>
                <div class="col-10">
                  <h4>{{ experience.title }}</h4>
                  <h5>{{ experience.company.name }} </h5>
                  <h6 class="font-weight-lighter">{{ experience.start_date.split("T")[0] }} - {{
                    experience.end_date.split("T")[0] }} {{ experience.location }}</h6>
                  <h6 class="font-weight-lighter"> {{ experience.location }}</h6>
                  <h6> {{ experience.description }}</h6>
                </div>
                <div class="col-1">
                  <b-button variant="link text-danger" @click="delete_education(education.id)"
                    v-if="edit_experiences">Delete</b-button>
                </div>
              </div>
            </div>
          </li> 
        </ul>
         
      </div>
    </div>
    <div class="card mt-4">
      <div class="card-body">
        <!-- Education Section -->
        <div class="row">
          <div class="col-11">
            <h4 class="card-title ">Education</h4>
          </div>
          <div class="col ml-4">
             
          </div>
        </div>

        <ul class="list-unstyled">
          <li v-for="education in educations" :key="education.id">
            <div class="border-bottom py-4">
              <div class="">
                <div class="row">
                  <div class="col-1">
                    <img
                      src="https://media.licdn.com/dms/image/D4D0BAQFQNpPNEX8Yng/company-logo_100_100/0/1684474197824?e=1694044800&v=beta&t=RYRrgx34t-oJTnSUz_C4aeMSNoVpiB3zj_XZs6Bd1xM"
                      style="height: 60px;" class="rounded" alt="...">
                  </div>
                  <div class="col-10">
                    <h4>{{ education.school_name }}</h4>
                    <h5>{{ education.degree }}, {{ education.field_of_study }} </h5>
                    <h6 class="font-weight-lighter">{{ education.start_date.split("T")[0] }} - {{
                      education.end_date.split("T")[0] }} {{ education.location }}</h6>
                    <h6 class="font-weight-lighter"> {{ education.location }}</h6>
                    <h6> {{ education.description }}</h6>
                  </div>
                  <div class="col-1">
                    <b-button variant="link text-danger" @click="delete_education(education.id)"
                      v-if="edit_educations">Delete</b-button>
                  </div>
                </div>
              </div>
            </div>
          </li>
        </ul> 
      </div>
    </div>
    <div class="card mt-4">
      <div class="card-body">
        <!-- Skills Section -->
        <div class="row">
          <div class="col-11">
            <h4 class="card-title ">Skills</h4>
          </div>
          
        </div>
        <ul class="list-unstyled">
          <li v-for="skill in skills" :key="skill.id">
            <div class="border-bottom mb-3">
              <div class="row">
                <div class="col-10">
                  <h6>{{ skill.name }}</h6>
                </div> 
              </div>
            </div>
          </li>
        </ul>
      </div>
      <Message></Message>
    </div>
  </div>
</template>

<script>
import '@/assets/main.css'
import axios from "axios";
import Message from '../components/Layout/Message.vue';
export default {
  name: 'LinkedInProfile',
  components: {Message},
  data() {
    return {
      profile_data: null,
      connection_count: null,

      profileImage: 'profile-image.jpg',

      profile_edit: {
        headline: "",
        bio: "",
        location: ""
      },

      //educations
      edit_educations: false,
      edit_educations_form: {
        school_name: "",
        degree: "",
        field_of_study: "",
        description: "",
        start_date: null,
        end_date: null
      },
      educations: [],
      //skills
      edit_skills: false,
      create_skill: {
        name: ""
      },
      skills: [],
      //experiences
      edit_experiences: false,
      create_experience_form: {
        company_id: "",
        title: "",
        location: "",
        description: "",
        start_date: null,
        end_date: null
      },
      experiences: [],
    };
  },
  async mounted() {
    this.profile_datas()
  },
  methods: {
    profile_datas() {
      const user_name =  this.$route.path.split("/")[2]
      axios.get(`https://software.ardapektezol.com/api/users/${user_name}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log(res.data.data);
        //user infos
        this.profile_data = res.data.data.user
        this.connection_count = res.data.data.connection_count 
        //educations
        if (res.data.data.education != null) {
          this.educations = []
          for (let index = 0; index < res.data.data.educations.length; index++) {
            this.educations.push({
              id: `${res.data.data.educations[index].id}`,
              description: `${res.data.data.educations[index].description}`,
              school_name: `${res.data.data.educations[index].school_name}`,
              field_of_study: `${res.data.data.educations[index].field_of_study}`,
              degree: `${res.data.data.educations[index].degree}`,
              start_date: `${res.data.data.educations[index].start_date}`,
              end_date: `${res.data.data.educations[index].end_date}`
            })
          }
        }
        if (res.data.data.skills != null) {
          //skills
          this.skills = []
          for (let index = 0; index < res.data.data.skills.length; index++) {
            this.skills.push({
              id: `${res.data.data.skills[index].id}`,
              name: `${res.data.data.skills[index].title}`
            })
          }
        }
        if (res.data.data.experiences != null) {
          //experience
        this.experiences = []
        for (let index = 0; index < res.data.data.experiences.length; index++) {
          this.experiences.push({
            id: `${res.data.data.experiences[index].id}`,
            company: {
              id: `${res.data.data.experiences[index].company.id}`,
              logo: `${res.data.data.experiences[index].company.logo}`,
              name: `${res.data.data.experiences[index].company.name}`
            },
            description: `${res.data.data.experiences[index].description}`,
            title: `${res.data.data.experiences[index].title}`,
            location: `${res.data.data.experiences[index].location}`,
            start_date: `${res.data.data.experiences[index].start_date}`,
            end_date: `${res.data.data.experiences[index].end_date}`
          })
        }
        }
        this.profile_edit.bio = this.profile_data.bio
        this.profile_edit.headline = this.profile_data.headline
        this.profile_edit.location = this.profile_data.location
      })
    },   
    post_connection(user_name){
      alert(this.$cookies.get("user_id"))
      axios.post(`https://software.ardapektezol.com/api/connections/${user_name}`, {}, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(() => {
        this.profile_datas()
      })
    }
  },
};
</script>

<style scoped> .header-photo {
   position: relative;
   width: 100%;
   height: 350px;
   overflow: hidden;
 }

 .profile-photo {
   position: absolute;
   top: 200px;
   width: 200px;
   height: 200px;
   overflow: hidden;
   z-index: 1;
 }

 .profile-photo img {
   object-fit: cover;
   width: 100%;
   height: 100%;
 }

 .card-body-about {
   position: relative;

   padding: 20px;

 }

 .button-group {
   display: flex;
   justify-content: left;
   margin-top: 15px;
 }

 .button-group button {
   margin: 0 5px;
 }</style>
