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
                <img src="../assets/images/linkedin.png" :alt="job.company" class="rounded-circle"  style = "position:relative; top:-15px;" />
              </div>
              <div class="job-info">
                <h5 class="card-title">{{ job.title }}</h5>
                <p class="card-text company-info">
                  {{ job.company }}
                </p>
                <p class="card-text company-info">
                  {{ job.location }}
                </p>
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
            <p class="card-text company-info">{{ selectedJob.company }}</p>
            <p class="card-text company-info">{{ selectedJob.location }}</p>
            <p class="card-text">{{ selectedJob.description }}</p>
            <p class="card-text">{{ selectedJob.requirements }}</p>
            <p class="card-text">{{ selectedJob.responsibilities }}</p>
            <div class="posted-by">
              <div class="posted-by-img">
                <img src="../assets/images/linkedin.png"  :alt="selectedJob.postedBy.name" class="img-fluid rounded-circle" style = "position:relative; top:-4px;"/>
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
export default {
  data() {
    return {
      jobs: [
        {
          id: 1,
          title: "Web Developer",
          employee: "Frontend Developer",
          company: "ABC Company",
          location: "San Francisco",
          logo: "path/to/logo1.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          requirements: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          responsibilities: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          postedBy: {
            name: "John Doe",
            photo: "path/to/photo1.png",
            position: "HR Manager",
          },
        },
        {
          id: 2,
          title: "Software Engineer",
          employee: "Full Stack Developer",
          company: "XYZ Corporation",
          location: "New York",
          logo: "path/to/logo2.png",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elitLorem ipsum dolor sit amet, consectetur adipiscing elit",
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
          employee: "Data Analyst",
          company: "123 Industries",
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
        // Add more job objects as needed
      ],
      selectedJob: null,
    };
  },
  methods: {
    selectJob(job) {
      this.selectedJob = job;
    },
    applyJob(job) {
      // Perform job application logic here
      console.log(`Applying for ${job.title} at ${job.company}`);
      // You can make an API call here to submit the application or display a confirmation message
    },
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