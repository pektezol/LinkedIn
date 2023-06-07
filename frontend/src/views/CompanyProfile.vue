<template>
  <div class="container">
    <div class="card mt-4">
      <div class="card-body">
        <div>
          <div class="header-photo">
            <img alt="Header Photo" src="https://example.com/header-image.png" />
          </div>
          <div class="profile-photo ml-5">
            <img class="img-fluid rounded-circle mb-1" :src="`${company_info.logo}`" alt="Profile Image" />
          </div>
          <b-card>
            <b-card-body>
              <div class="row">
                <div class="col">
                  <h5 class="card-title">{{ company_info.name }}</h5>
                </div>
                <div class="col">
                  <ul class="list-unstyled">
                    <li><i class="fa fa-map-marker mr-2"></i>Location: {{ company_info.location }}</li>
                    <li><i class="fa fa-map-marker mr-2"></i>Industry: {{ company_info.industry }}</li>
                  </ul>
                </div>
              </div>
            </b-card-body>
            <b-button id="show-btn" @click="$bvModal.show('bv-modal-example-company')">Update Company</b-button>
            <b-modal id="bv-modal-example-company" hide-footer>
              <form @submit.prevent="update_company_func()">
                <div class="form-group">
                  <label for="position">Name</label>
                  <input type="text" class="form-control" v-model="update_company.name">
                </div>
                <div class="form-group">
                  <label for="position">Industry</label>
                  <input type="text" class="form-control" v-model="update_company.industry">
                </div>
                <div class="form-group">
                  <label for="position">Location</label>
                  <input type="text" class="form-control" v-model="update_company.location">
                </div>
                <div class="form-group">
                  <label for="position">Description</label>
                  <b-form-textarea id="textarea" v-model="update_company.description" placeholder="Enter something..."
                    rows="3" max-rows="6"></b-form-textarea>
                </div>
                <div class="form-group">
                  <label for="position">Logo File</label>
                  <b-form-file id="file-default" v-model="update_company.logo"></b-form-file>
                  <div class="mt-3">Selected file: {{ update_company.logo ? update_company.logo.name : '' }}</div>
                </div>
                <div class="row">
                  <div class="row ml-4 ">
                    <b-button type="submit" @click="$bvModal.hide('bv-modal-example-company')"
                      variant="success">Submit</b-button>
                  </div>
                </div>
              </form>
            </b-modal>
          </b-card>
        </div>
      </div>
    </div>
    <div class="card mt-1">
      <div class="button-group2">
        <button @click="navigateTo('home')">Ana Sayfa</button>
        <button @click="navigateTo('about')">Hakkında</button>
        <button @click="navigateTo('jobs')">İş İlanları</button>
        <button @click="navigateTo('team')">İş Arkadaşları</button>
      </div>
    </div>
    <div class="card mt-4">
      <div class="card-body">
        <!-- Job Positions Section -->
        <div v-if="activeSection === 'jobs'">
          <div class="row">
            <div class="col-10">
              <h2 class="card-title">Job Positions</h2>
            </div>
            <div class="col-2 mt-2" v-if="owner">
              <b-button variant="outline-success" @click="create_job_toggle = !create_job_toggle">Create</b-button>
            </div>
          </div>
          <div class="ml-2" v-if="owner && create_job_toggle">
            <div class="border-bottom">
              <div class="row px-4 pb-2">
                <b-form @submit.prevent="create_job_opening()">
                  <div class="row">
                    <div class="col-3">
                      <b-form-input v-model="create_job_openings.title" placeholder="Title" required></b-form-input>
                      <b-form-input v-model="create_job_openings.type" class="mt-2" placeholder="Type"
                        required></b-form-input>
                    </div>
                    <div class="col-6 ">
                      <b-form-textarea id="textarea" v-model="create_job_openings.description"
                        placeholder="Description..." rows="3" max-rows="6"></b-form-textarea>
                    </div>
                    <div class="col-3">
                      <b-form-input v-model="create_job_openings.location" placeholder="Location" required></b-form-input>
                      <b-button type="submit" variant="success" class="mt-2 ml-4">Submit</b-button>
                    </div>
                  </div>
                </b-form>
              </div>
            </div>
          </div>
          <ul class="list-unstyled">
            <li v-for="open in open_job" :key="open.id">
              <div class="card mt-4">
                <div class="card-body-about">
                  <div class="row">
                    <h4 class="card-title">{{ open.title }}</h4>
                    <p class="ml-2 mt-3">{{ open.type }}</p>
                    <h6 class="font-weight-lighter ml-4 mt-3">{{ open.location }}</h6>
                    <p class="ml-4 mt-3">{{ open.date.split("T")[0] }}</p>
                  </div>
                  <div v-for="(item) in job_applications" :key="item.id">
                    <div class="row ml-2 my-2 border-bottom" v-if="item.job.id == open.id">
                      <h6 class="mt-2">{{ item.user.first_name }} {{ item.user.last_name }}</h6>
                      <p class="ml-4 mt-2"> {{ item.user.headline }}</p>
                      <a :href="`/user/${item.user.user_name}`" class="mt-2 ml-4">Visit profile</a>
                      <b-button variant="link" class="ml-5" @click="accept_job_req(item.id)"
                        v-if="owner">Accept</b-button>
                    </div>
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>
        <div v-else-if="activeSection === 'about'">
          <div class="card mt-4">
            <div class="card-body-about">
              <!-- About Section -->
              <h4 class="card-title">About</h4>
              <p>{{ company_info.description }}</p>
            </div>
          </div>
        </div>

        <!-- Team Members Section -->
        <div v-else-if="activeSection === 'team'">
          <div class="row">
            <div class="col-11">
              <h4 class="card-title">Team Members</h4>
            </div>
          </div>
          <ul class="list-unstyled">
            <li v-for="member in teamMembers" :key="member.id">
              <div class="card mt-4">
                <div class="row">
                  <div class="col-2">
                    <img :src="member.photo" class="rounded-circle" alt="Profile Image" />
                  </div>
                  <div class="col-10">
                    <h4>{{ member.name }}</h4>
                    <h5>{{ member.position }}</h5>
                    <h6>{{ member.location }}</h6>
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>
        <!-- Default Section (Home) -->
        <div v-else>
          <h2>Home Section</h2>
          <!-- Content for the home section goes here -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'CompanyPage',
  data() {
    return {
      owner: true,
      company_info: [],
      job_applications: [],
      create_job_toggle: false,
      open_job: [],
      create_job_openings: {
        title: "",
        location: "",
        description: "",
        type: ""
      },
      update_company: {
        name: "",
        industry: "",
        location: "",
        description: "",
        logo: ""
      },
      activeSection: 'home',
      companyData: {
        name: "Company Name",
        headline: "Company Headline",
        email: "company@example.com",
        location: "Company Location",
        about: "Company description goes here"
      },
      connectionCount: 0,
      jobPositions: [
        {
          id: 1,
          title: "Job Position 1",
          department: "Department 1",
          location: "Location 1"
        },
        {
          id: 2,
          title: "Job Position 2",
          department: "Department 2",
          location: "Location 2"
        },
        {
          id: 3,
          title: "Job Position 3",
          department: "Department 3",
          location: "Location 3"
        }
      ],
      teamMembers: [
        {
          id: 1,
          name: "Team Member 1",
          position: "Position 1",
          location: "Location 1",
          photo: "https://example.com/profile1.png"
        },
        {
          id: 2,
          name: "Team Member 2",
          position: "Position 2",
          location: "Location 2",
          photo: "https://example.com/profile2.png"
        },
        {
          id: 3,
          name: "Team Member 3",
          position: "Position 3",
          location: "Location 3",
          photo: "https://example.com/profile3.png"
        }
      ],
      posts: [
        {
          id: 1,
          title: "Post Title 1",
          content: "Post content goes here"
        },
        {
          id: 2,
          title: "Post Title 2",
          content: "Post content goes here"
        },
        {
          id: 3,
          title: "Post Title 3",
          content: "Post content goes here"
        }
      ]
    };
  },
  mounted() {

    this.reload()
  },
  methods: {
    reload() {
      this.get_company_info()
      this.get_job_application()
      this.get_job_open()
    },
    navigateTo(section) {
      this.activeSection = section;
    }, 
    get_company_info() {
      const company_id = this.$route.path.split("/")[3]
      axios.get(`https://software.ardapektezol.com/api/company/${company_id}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.company_info = res.data.data.companies[0]
      })
    },
    update_company_func() {
      console.log(this.update_company.logo)
      if (this.update_company.name == "") {
        this.update_company.name = this.company_info.name
      }
      if (this.update_company.industry == "") {
        this.update_company.industry = this.company_info.industry
      }
      if (this.update_company.location == "") {
        this.update_company.location = this.company_info.location
      }
      if (this.update_company.description == "") {
        this.update_company.description = this.company_info.description
      }
      if (this.update_company.logo == "") {
        this.update_company.logo = this.company_info.logo
        const company_id = this.$route.path.split("/")[3]
        axios.put(`https://software.ardapektezol.com/api/company/${company_id}`, {
          name: this.update_company.name,
          industry: this.update_company.industry,
          location: this.update_company.location,
          description: this.update_company.description,
          logo: this.update_company.logo
        }, {
          headers: {
            Authorization: this.$cookies.get("token")
          }
        }).then(res => {
          this.reload()
        })
      } else {
        const company_id = this.$route.path.split("/")[3]
        const file = this.update_company.logo
        const reader = new FileReader()
        let rawImg;
        reader.onloadend = () => {
          rawImg = reader.result;
          axios.put(`https://software.ardapektezol.com/api/company/${company_id}`, {
            name: this.update_company.name,
            industry: this.update_company.industry,
            location: this.update_company.location,
            description: this.update_company.description,
            logo: rawImg
          }, {
            headers: {
              Authorization: this.$cookies.get("token")
            }
          }).then(res => {
            this.reload()
          })
        }
        reader.readAsDataURL(file);
      }

    },
    get_job_application() {
      const company_id = this.$route.path.split("/")[3]
      axios.get(`https://software.ardapektezol.com/api/company/${company_id}/job`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.job_applications = res.data.data.applications
        console.log(res.data.data.applications);
        //this.job_applications = res.data.data.companies[0]
      })
    },
    get_job_open() {
      const company_id = this.$route.path.split("/")[3]
      axios.get(`https://software.ardapektezol.com/api/jobs/${company_id}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.open_job = res.data.data.openings
        console.log(res.data.data.openings);
        //this.job_applications = res.data.data.companies[0]
      })
    },
    accept_job_req(application_id) {
      axios.put(`https://software.ardapektezol.com/api/jobs/${application_id}`,{}, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => { 
        this.reload()
        //this.job_applications = res.data.data.companies[0]
      })
    },
    create_job_opening() {
      const company_id = this.$route.path.split("/")[3]
      axios.post(`https://software.ardapektezol.com/api/company/${company_id}/job`, {
        title: this.create_job_openings.title,
        location: this.create_job_openings.location,
        description: this.create_job_openings.description,
        type: this.create_job_openings.type
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => { 
        //this.job_applications = res.data.data.companies[0]
        this.create_job_toggle = false
        this.reload()
      })
    }

  },

};
</script>

<style scoped>
.button-group {
  margin-bottom: 20px;
}

.button-group button {
  margin-right: 10px;
}

.card {
  padding: 5px;
  border: 2px solid #ccc;
}

.header-photo {
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

.button-group2 button {
  margin-right: 10px;
  padding: 10px 20px;
  font-size: 16px;
  background-color: transparent;
  color: rgb(152, 152, 152);
  cursor: pointer;
  transition: color 0.3s;
}

.button-group2 button:hover {
  color: #000000;
}

.button-group2 button.active {
  color: #007bff;
  background-color: rgb(0, 255, 89);
  /* Set the background color to green */
}
</style>
