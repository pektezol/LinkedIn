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
            <img class="img-fluid rounded-circle mb-1" :src="profile_data.profile_picture" alt="Profile Image" />
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
                <div>
                  <div>
                    <b-button id="show-btn" @click="$bvModal.show('bv-modal-example')"><b-icon class="mr-2"
                        icon="gear-fill" aria-hidden="true"></b-icon>Edit Profile</b-button>
                    <b-modal id="bv-modal-example" hide-footer>
                      <form @submit.prevent="saveChangesProfileEdit()">
                        <div class="form-group">
                          <label for="position">Headline</label>
                          <input type="text" class="form-control" v-model="profile_edit.headline">
                        </div>
                        <div class="form-group">
                          <label for="position">Bio</label>
                          <input type="text" class="form-control" v-model="profile_edit.bio">
                        </div>
                        <div class="form-group">
                          <label for="position">Location</label>
                          <input type="text" class="form-control" v-model="profile_edit.location">
                        </div>
                        <div class="row">
                          <div class="row ml-4 ">
                            <b-button type="submit" @click="$bvModal.hide('bv-modal-example')"
                              variant="success">Submit</b-button>
                          </div>
                        </div>
                      </form>
                    </b-modal>

                    <b-button id="show-btn" @click="$bvModal.show('bv-modal-example-company')">Create Company</b-button>
                    <b-modal id="bv-modal-example-company" hide-footer>
                      <form @submit.prevent="create_company()">
                        <div class="form-group">
                          <label for="position">Name</label>
                          <input type="text" class="form-control" v-model="company_form.name">
                        </div>
                        <div class="form-group">
                          <label for="position">Industry</label>
                          <input type="text" class="form-control" v-model="company_form.industry">
                        </div>
                        <div class="form-group">
                          <label for="position">Location</label>
                          <input type="text" class="form-control" v-model="company_form.location">
                        </div>
                        <div class="form-group">
                          <label for="position">Description</label>
                          <b-form-textarea id="textarea" v-model="company_form.description"
                            placeholder="Enter something..." rows="3" max-rows="6"></b-form-textarea>
                        </div>
                        <div class="form-group">
                          <label for="position">Logo File</label> 
                            <b-form-file id="file-default" v-model="company_form.logo"></b-form-file> 
                          <div class="mt-3">Selected file: {{ company_form.logo ? company_form.logo.name : '' }}</div>
                        </div>
                        <div class="row">
                          <div class="row ml-4 ">
                            <b-button type="submit" @click="$bvModal.hide('bv-modal-example-company')"
                              variant="success">Submit</b-button>
                          </div>
                        </div>
                      </form>
                    </b-modal>
                  </div>
                </div>
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
            <b-button size="sm" class="mb-2" @click="open_experience_settings()">
              <b-icon icon="gear-fill" aria-hidden="true"></b-icon>
            </b-button>
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
        <div class="row">
          <div class="border-bottom">
            <div class="row px-4 pb-2">
              <b-form @submit.prevent="create_experience()" v-if="edit_experiences">
                <div class="row" style="max-width: 100%;">
                  <div class="col">
                    <b-form-input v-model="create_experience_form.company_id" placeholder="Company Name"
                      required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-input v-model="create_experience_form.title" placeholder="Title" required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-input v-model="create_experience_form.location" placeholder="Location"
                      required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-datepicker v-model="create_experience_form.start_date" class="mb-2"></b-form-datepicker>
                  </div>
                  <div class="col">
                    <b-form-datepicker v-model="create_experience_form.end_date" class="mb-2"></b-form-datepicker>
                  </div>
                </div>
                <div class="row">
                  <div class="col-10">
                    <b-form-textarea id="textarea" v-model="create_experience_form.description"
                      placeholder="Description..." rows="3" max-rows="6"></b-form-textarea>
                  </div>
                  <div class="col">
                    <div class="row ml-4 mt-5">
                      <b-button type="submit" variant="success">Submit</b-button>
                    </div>
                  </div>
                </div>
              </b-form>
            </div>
          </div>
        </div>
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
            <b-button size="sm" class="mb-2"
              @click="edit_educations == false ? edit_educations = true : edit_educations = false">
              <b-icon icon="gear-fill" aria-hidden="true"></b-icon>
            </b-button>
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
        <div class="row">
          <div class="border-bottom">
            <div class="row px-4 pb-2">
              <b-form @submit.prevent="create_education()" v-if="edit_educations">
                <div class="row" style="max-width: 100%;">
                  <div class="col">
                    <b-form-input v-model="edit_educations_form.school_name" placeholder="School Name"
                      required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-input v-model="edit_educations_form.degree" placeholder="Degree" required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-input v-model="edit_educations_form.field_of_study" placeholder="Field Of Study"
                      required></b-form-input>
                  </div>
                  <div class="col">
                    <b-form-datepicker v-model="edit_educations_form.start_date" class="mb-2"></b-form-datepicker>
                  </div>
                  <div class="col">
                    <b-form-datepicker v-model="edit_educations_form.end_date" class="mb-2"></b-form-datepicker>
                  </div>
                </div>
                <div class="row">
                  <div class="col-10">
                    <b-form-textarea id="textarea" v-model="edit_educations_form.description" placeholder="Description..."
                      rows="3" max-rows="6"></b-form-textarea>
                  </div>
                  <div class="col">
                    <div class="row ml-4 mt-5">
                      <b-button type="submit" variant="success">Submit</b-button>
                    </div>
                  </div>
                </div>
              </b-form>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="card mt-4">
      <div class="card-body">
        <!-- Skills Section -->
        <div class="row">
          <div class="col-11">
            <h4 class="card-title ">Skills</h4>
          </div>
          <div class="col ml-4">
            <b-button size="sm" class="mb-2" @click="edit_skills == false ? edit_skills = true : edit_skills = false">
              <b-icon icon="gear-fill" aria-hidden="true"></b-icon>
            </b-button>
          </div>
        </div>
        <ul class="list-unstyled">
          <li v-for="skill in skills" :key="skill.id">
            <div class="border-bottom mb-3">
              <div class="row">
                <div class="col-10">
                  <h6>{{ skill.name }}</h6>
                </div>
                <div class="col-1">
                  <b-button variant="link text-danger" @click="delete_skill(skill.id)"
                    v-if="edit_skills">Delete</b-button>
                </div>
              </div>
            </div>
          </li>
        </ul>
        <div class="row">
          <div class="border-bottom">
            <div class="row px-4 pb-2">
              <b-form @submit.prevent="create_skill_func()" v-if="edit_skills">
                <div class="row" style="max-width: 100%;">
                  <div class="col">
                    <b-form-input v-model="create_skill.name" placeholder="Skill Name" required></b-form-input>
                  </div>
                  <div class="col">
                    <b-button type="submit" variant="success">Submit</b-button>
                  </div>
                </div>
              </b-form>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Message></Message>
  </div>
</template>

<script>
import '@/assets/main.css'
import axios from "axios";
import Message from '../components/Layout/Message.vue';
export default {
  name: 'LinkedInProfile',
  components: { Message },
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
      //Create Company
      company_form: {
        name: "",
        industry: "",
        location: "",
        description: "",
        logo: null
      }
    };
  },
  async mounted() {
    this.profile_datas()
  },
  methods: {
    profile_datas() {
      axios.get(`https://software.ardapektezol.com/api/profile`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {

        //user infos
        this.profile_data = res.data.data.user
        this.connection_count = res.data.data.connection_count
        this.$cookies.set('user_id', this.profile_data.id)
        //educations

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
        //skills
        this.skills = []
        for (let index = 0; index < res.data.data.skills.length; index++) {
          this.skills.push({
            id: `${res.data.data.skills[index].id}`,
            name: `${res.data.data.skills[index].title}`
          })
        }
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
        this.profile_edit.bio = this.profile_data.bio
        this.profile_edit.headline = this.profile_data.headline
        this.profile_edit.location = this.profile_data.location
      })
    },
    saveChangesProfileEdit() {
      axios.put(`https://software.ardapektezol.com/api/profile`, {
        headline: this.profile_edit.headline,
        bio: this.profile_edit.bio,
        location: this.profile_edit.location
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(() => {
        this.profile_datas()
      })
    },
    resetForm() {
      // Reset form fields or any other necessary cleanup
      this.nameEdit = '';
      this.surnameEdit = '';
      this.emailEdit = '';
      this.position = '';
      this.education = '';
      this.locationEdit = '';
      this.editedAbout = '';
    },
    // EXPERIENCE
    open_experience_settings() {
      // şirket isteği at 
      if (this.edit_experiences == false) {
        this.edit_experiences = true
      } else {
        this.edit_experiences = false
      }
    },
    create_experience() {
      axios.post(`https://software.ardapektezol.com/api/experience`, {
        company_id: parseInt(this.create_experience_form.company_id),
        title: this.create_experience_form.title,
        description: this.create_experience_form.description,
        location: this.create_experience_form.location,
        start_date: this.create_experience_form.start_date.split("T")[0],
        end_date: this.create_experience_form.end_date.split("T")[0]
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(() => {
        this.profile_datas()
      })
    },
    // EDUCATION
    create_education() {
      axios.post(`https://software.ardapektezol.com/api/education`, {
        school_name: this.edit_educations_form.school_name,
        degree: this.edit_educations_form.degree,
        field_of_study: this.edit_educations_form.field_of_study,
        description: this.edit_educations_form.description,
        start_date: this.edit_educations_form.start_date,
        end_date: this.edit_educations_form.end_date
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.profile_datas()
      })
    },
    delete_education(ed_id) {
      axios.delete(`https://software.ardapektezol.com/api/education/${ed_id}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.profile_datas()
      })
    },
    //SKILLS
    create_skill_func() {
      axios.post(`https://software.ardapektezol.com/api/skill`, {
        name: this.create_skill.name
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.profile_datas()
        this.create_skill.name = ""
      })
    },
    delete_skill(ed_id) {
      axios.delete(`https://software.ardapektezol.com/api/skill/${ed_id}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.profile_datas()
      })
    },
    create_company() {   

      const file = this.company_form.logo
      const reader = new FileReader()
      let rawImg;
      reader.onloadend = () => {
        rawImg = reader.result; 
        axios.post(`https://software.ardapektezol.com/api/company`, {
        name: this.company_form.name,
        industry: this.company_form.industry,
        location: this.company_form.location,
        description: this.company_form.description,
        logo: rawImg
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log("company response", res.data.data);
      })
      }
      reader.readAsDataURL(file);   

      
    },
    encode_file(data) {
      const file = data
      const reader = new FileReader()
      let rawImg;
      reader.onloadend = () => {
        rawImg = reader.result;
        console.log(rawImg); 
      }
      reader.readAsDataURL(file);   
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
 }
</style>
