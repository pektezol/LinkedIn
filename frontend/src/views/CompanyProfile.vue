<template>
  <div class="container">

    <div class="card mt-4">
      <div class="card-body">
        <div>
          <div class="header-photo">
            <img alt="Header Photo" src="https://example.com/header-image.png" />
          </div>
          <div class="profile-photo ml-5">
            <img class="img-fluid rounded-circle mb-1" src="../assets/images/profilPhoto.jpeg" alt="Profile Image" />
          </div>
          <b-card>
            <b-card-body>
              <div class="row">
                <div class="col">
                  <h5 class="card-title">{{ companyData.name }}</h5>
                  <h6>{{ companyData.headline }}</h6>
                  <h6 class="text-primary">Connection Count: {{ connectionCount }}</h6>
                </div>
                <div class="col">
                  <ul class="list-unstyled">
                    <li><i class="fa fa-envelope-o mr-2"></i>Email: {{ companyData.email }}</li>
                    <li><i class="fa fa-map-marker mr-2"></i>Location: {{ companyData.location }}</li>
                  </ul>
                </div>
              </div>
              <div class="button-group">
                <b-button variant="primary" @click="connectWithCompany">Connect</b-button>
                <b-button variant="primary">Message</b-button>
              </div>
            </b-card-body>
          </b-card>
        </div>
      </div>
    </div>

    <div class="card mt-1">
     <div class="button-group2">
      <button @click="navigateTo('home')">Ana Sayfa</button>
      <button @click="navigateTo('about')">Hakkında</button>
      <button @click="navigateTo('posts')">Gönderiler</button>
      <button @click="navigateTo('jobs')">İş İlanları</button>
    </div>
    </div>

    <div class="card mt-4">
      <div class="card-body-about">
        <!-- About Section -->
        <h4 class="card-title">About</h4>
        <p>{{ companyData.about }}</p>
      </div>
    </div>


    <div class="card mt-4">
      <div class="card-body">
        <!-- Job Positions Section -->
        <div v-if="activeSection === 'jobs'">
          <div class="row">
            <div class="col-11">
              <h4 class="card-title">Job Positions</h4>
            </div>
            <div class="col ml-4">
            </div>
          </div>
          <ul class="list-unstyled">
            <li v-for="position in jobPositions" :key="position.id">
              <div class="border-bottom">
                <div class="row">
                  <div class="col-12">
                    <h4>{{ position.title }}</h4>
                    <h5>{{ position.department }}</h5>
                    <h6 class="font-weight-lighter">{{ position.location }}</h6>
                  </div>
                </div>
              </div>
            </li>
          </ul>
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
              <div class="border-bottom">
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

        <!-- Posts Section -->
        <div v-else-if="activeSection === 'posts'">
          <div class="row">
            <div class="col-11">
              <h4 class="card-title">Posts</h4>
            </div>
          </div>
          <ul class="list-unstyled">
            <li v-for="post in posts" :key="post.id">
              <div class="border-bottom">
                <div class="row">
                  <div class="col-12">
                    <h4>{{ post.title }}</h4>
                    <p>{{ post.content }}</p>
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
export default {
  name: 'CompanyPage',
  data() {
    return {
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
  methods: {
    navigateTo(section) {
      this.activeSection = section;
    },
    connectWithCompany() {
      // Connect with the company logic
    }

  },
  mounted() {
    this.connectionCount = Math.floor(Math.random() * 100); // Random connection count
  }
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
  background-color: rgb(0, 255, 89); /* Set the background color to green */
}


</style>
