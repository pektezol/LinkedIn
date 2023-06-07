<template>
  <div class="container">
    <div class="row border-bottom mb-3" v-if="users_companies.length != 0">
      <h4 class="ml-4">Your Company(ies)</h4>
      <div class="row">

        <div class="col mt-35" v-for="(company) in users_companies" :key="company.id">
          <b-card :title="company.name" :img-src="company.logo" img-alt="Image" img-top tag="article"
            style="max-width: 20rem;" class="mb-2">
            <p>Industry: {{ company.industry }}</p>
            <p>Location: {{ company.location }}</p>
            <b-button :href="`company/${company.name}/${company.id}`" variant="primary">Go Your Company Page</b-button>
          </b-card>
        </div>

      </div>
    </div>


    <div class="row">
      <div class="col-md-4">
        <div class="profile-box">
          <h4 class="profile-box-title">Job Postings Based on Your Profile</h4>
        </div>
        <div v-for="(job, index) in jobs" :key="job.id" class="card mb-3">
          <div class="card-body" @click="selectJob(job, index)">
            <div class="company-details">
              <div class="company-logo"> 
                 
                    <b-avatar :src="job.company.logo"></b-avatar>
              </div>
              <div class="job-info">
                <h5 class="card-title">{{ job.title }}</h5>
                <p class="card-text company-info">{{ job.company.name }}</p>
                <p class="card-text company-info">{{ job.location }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-md-8">
        <div v-if="selectedJob" class="card">
          <div class="card-body">
            <div class="card-title">
              <h5>{{ selectedJob.title }}</h5>
            </div>
            <p class="card-text company-info">{{ selectedJob.company.name }}</p>
            <p class="card-text company-info">{{ selectedJob.location }}</p>
            <p class="card-text">{{ selectedJob.description }}</p>
            <p class="card-text">{{ selectedJob.requirements }}</p>
            <p class="card-text">{{ selectedJob.responsibilities }}</p>
            <div class="posted-by">
              <div class="posted-by-img">
                <img :src="selectedJob.company.employer.profile_picture" :alt="selectedJob.company.employer.user_name"
                  class="img-fluid rounded-circle" style="position:relative; top:-4px;" />
              </div>
              <div class="posted-by-info">
                <p class="card-text mb-1">
                  <strong>Posted By: </strong> {{ selectedJob.company.employer.first_name }} {{
                    selectedJob.company.employer.last_name }}
                </p>
                <p class="card-text mb-0">
                  <strong>Position: </strong> {{ selectedJob.company.employer.headline }}
                </p>
              </div>
            </div>
            <button class="btn btn-primary" @click="applyJob(selectedJob)">Apply</button>
          </div>
        </div>
        <div v-else class="alert alert-info mt-4" role="alert">
          Click on a job to view more details
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';


export default {
  data() {
    return {
      jobs: [
        /* {
          id: 1,
          title: "Senior Embedded Software Developer",
          type: "Backend Developer",
          company: {
            id: 2,
            name: "Siemens Company",
            logo: "path/to/logo1.png",
            employer: {
              id: "1",
              user_name: "John Doe",
              first_name: "Volkan",
              last_name: "Öztoklu",
              headline: "Student of mef",
              profile_picture: "../assets/images/hRPhotoArda_.jpeg",
              position: "HR Manager",
            },
          },
          location: "San Francisco",
          description: "asfdasdfasdfasdfasdfasdfasfdasdfasdfasdfasdfasdfasdfasdfasdfasdf",
          date: "01/01/2020"
        }, */
      ],
      selectedJob: null,
      users_companies: []
    };
  },
  watch: {

  },
  mounted() {
    this.get_job()
    this.get_company()
  },
  methods: {
    selectJob(job, index) {
      this.selectedJob = job;
    },
    applyJob(job) {
      // Implement your apply job functionality here
      console.log("Applying for job:", job.title);
      axios.post(`https://software.ardapektezol.com/api/jobs/${job.id}`, {},{
            headers: {
              Authorization: this.$cookies.get("token")
            }
          })
        .then(res => {
          this.jobs = res.data.data.openings
          console.log(res.data.data.openings);
        })
    },
    get_job() {
      axios.get(`https://software.ardapektezol.com/api/jobs`)
        .then(res => {
          this.jobs = res.data.data.openings
          console.log(res.data.data.openings);
        })
    },
    get_company() {
      axios.get(`https://software.ardapektezol.com/api/company`)
        .then(res => {
          console.log(res.data.data);
          if (res.data.data.companies.length != 0) {
            for (let index = 0; index < res.data.data.companies.length; index++) {
              if (res.data.data.companies[index].employer.id == this.$cookies.get("user_id")) {
                this.users_companies.push(res.data.data.companies[index])
              }
            }
          }
        })
    }
  },
};
</script>
<style scoped>
.profile-box {
  background-color: #007bff;
  color: #fff;
  padding: 10px;
  margin-bottom: 20px;
}

.profile-box-title {
  margin: 0;
}

.card {
  cursor: pointer;
}

.alert-info {
  margin-top: 20px;
}

.company-details {
  display: flex;
  align-items: center;
}

.company-logo {
  width: 50px;
  height: 50px;
  margin-right: 10px;
}

.logo-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.job-info {
  display: flex;
  flex-direction: column;
}

.card-title {
  margin-top: 0;
  margin-bottom: 5px;
}

.company-info {
  margin-bottom: 0;
}

.posted-by {
  display: flex;
  align-items: center;
  margin-top: 10px;
}

.posted-by-img {
  width: 50px;
  height: 50px;
  margin-right: 10px;
}

.img-fluid {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
</style>