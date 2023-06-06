<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4">
        <div class="profile-box">
          <h4 class="profile-box-title">Job Postings Based on Your Profile</h4>
        </div>
        <div v-for="(job, index) in jobs" :key="job.id" class="card mb-3">
          <div class="card-body" @click="selectJob(job, index)">
            <div class="company-details">
              <div class="company-logo">
               
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
                <img src="../assets/images/hRPhotoArda_.jpeg" :alt="selectedJob.postedBy.name" class="img-fluid rounded-circle" style="position:relative; top:-4px;" />
              </div>
              <div class="posted-by-info">
                <p class="card-text mb-1">
                  <strong>Posted By: </strong>{{ selectedJob.postedBy.name }}
                </p>
                <p class="card-text mb-0">
                  <strong>Position: </strong>{{ selectedJob.postedBy.position }}
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
        {
          id: 1,
          title: "Senior Embedded Software Developer",
          type: "Backend Developer",
          company: {
            id: 2,
            name: "Siemens Company",
            logo: ""
          },
          location: "San Francisco",
          logo: "path/to/logo1.png",
          description: "",
          requirements: "",
          responsibilities: "",
          postedBy: {
            name: "John Doe",
            photo: "path/to/photo1.png",
            position: "HR Manager",
          },
          date: "01/01/2020"
        },
        {
          id: 2,
          title: "Software Engineer",
          employee: "Full Stack Developer",
          company: "Evil Co",
          location: "New York",
          logo: "../assets/images/linkedin.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elit",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Jane Smith",
            photo: "path/to/photo2.png",
            position: "Senior Software Engineer",
          },
        },
        {
          id: 3,
          title: "Data Analyst",
          employee: "Data Analyitica",
          company: "Alpha Industries",
          location: "Chicago",
          logo: "path/to/logo3.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Alice Johnson",
            photo: "path/to/photo3.png",
            position: "Data Science Manager",
          },
        },
        {
          id: 4,
          title: "Graphic Designer",
          company: "Samsung",
          location: "London",
          logo: "path/to/logo4.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "David Brown",
            photo: "path/to/photo4.png",
            position: "Creative Director",
          },
        },
        {
          id: 5,
          title: "UI/UX Designer",
          company: "Dijital Dusler Co.",
          location: "Los Angeles",
          logo: "path/to/logo5.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Emily Brown",
            photo: "path/to/photo5.png",
            position: "Lead Designer",
          },
        },
        {
          id: 6,
          title: "Product Manager",
          company: "Beyin Takimi Solutions",
          location: "San Francisco",
          logo: "path/to/logo6.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Michael Clark",
            photo: "path/to/photo6.png",
            position: "Product Director",
          },
        },
        {
          id: 7,
          title: "Data Scientist",
          company: "Big Data Analytics",
          location: "New York",
          logo: "path/to/logo7.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Sophia Roberts",
            photo: "path/to/photo7.png",
            position: "Data Science Lead",
          },
        },
        {
          id: 8,
          title: "Marketing Specialist",
          company: "Digital Marketing Agency",
          location: "Chicago",
          logo: "path/to/logo8.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Daniel Wilson",
            photo: "path/to/photo8.png",
            position: "Marketing Manager",
          },
        },
        {
          id: 9,
          title: "Business Analyst",
          company: "Global Consulting",
          location: "London",
          logo: "path/to/logo9.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Olivia Adams",
            photo: "path/to/photo9.png",
            position: "Business Analyst Lead",
          },
        },
        {
          id: 10,
          title: "Sales Manager",
          company: "E-commerce Solutions",
          location: "Sydney",
          logo: "path/to/logo10.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "Ryan Taylor",
            photo: "path/to/photo10.png",
            position: "Sales Director",
          },
        },
      ],
      selectedJob: null,
    };
  },
  watch: {
    selectJob() {
      console.log("sdfgsdgfsdf")
    }
  },
  mounted(){
    this.get_job()
  },
  methods: {
    selectJob(job, index) {
      this.selectedJob = job;
    },
    applyJob(job) {
      // Implement your apply job functionality here
      console.log("Applying for job:", job.title);
    },
    get_job() { 
      axios.get(`https://software.ardapektezol.com/api/jobs`)
      .then(res => {
         console.log(res.data.data.openings);
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