<template>
    <div class="container">
      <div class="card mt-4">
        <div class="card-body">
    <div>
      <div class="header-photo">
        <img alt="Header Photo" src="https://64.media.tumblr.com/682b7be9273636dffb1d8fbe3220628b/tumblr_pdz53tIneb1sx8ybdo10_1280.png"/>
      </div>
      <div class="profile-photo ml-5">
        <img class="img-fluid rounded-circle mb-1" src="../assets/images/profilPhoto.jpeg" alt="Profile Image" />
      </div> 
        <b-card>
          <b-card-body>
            <h5 class="card-title">{{ profile_data.first_name }} {{ profile_data.last_name }}</h5>
                <ul class="list-unstyled">
                  <li><i class="fa fa-envelope-o mr-2"></i>{{ profile_data.email }}</li>  
                  <li><i class="fa fa-map-marker mr-2"></i>{{ profile_data.location }}</li>
                </ul>
                <div class="button-group">
                  <button class="btn btn-primary">Connect</button>
                  <button class="btn btn-primary">Message</button>
                  <button class="btn btn-primary">Follow</button> 
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
            <button class="btn btn-secondary" @click="openModal1">Edit</button>
          </div>
        </div>
        <b-modal v-model="modalOpen1" @ok="saveChanges" @hidden="resetForm">
          <h5>Edit About</h5>
          <form>
            <div class="form-group">
              <label for="about">About</label>
              <textarea class="form-control" v-model="editedAbout"></textarea>
            </div>
          </form>
        </b-modal>
      </div>
  
      <div class="card mt-4">
        <div class="card-body">
          <!-- Education Section -->
          <h4 class="card-title">Education</h4>
          <ul class="list-unstyled">
            <li v-for="education in educations" :key="education.id">
              <h6>{{ education.degree }}</h6>
              <p class="text-muted">{{ education.university }}</p>
            </li>
          </ul>
        </div>
      </div>
  
      <div class="card mt-4">
        <div class="card-body">
          <!-- Experience Section -->
          <h4 class="card-title">Experience</h4>
          <ul class="list-unstyled">
            <li v-for="experience in experiences" :key="experience.id">
              <h6>{{ experience.position }}</h6>
              <p class="text-muted">{{ experience.company }}</p>
              <p>{{ experience.description }}</p>
            </li>
          </ul>
        </div>
      </div>
  
      <div class="card mt-4">
        <div class="card-body">
          <!-- Skills Section -->
          <h4 class="card-title">Skills</h4>
          <ul class="list-unstyled">
            <li v-for="skill in skills" :key="skill.id">{{ skill.name }}</li>
          </ul>
        </div>
      </div>
  
    </div>
  </template>
  
  <script>
  import '@/assets/main.css'
  import axios from "axios";
  export default {
    name: 'LinkedInProfile',
    data() {
      return {
        profile_data: null,
        profileImage: 'profile-image.jpg',
        modalOpen: false,
        modalOpen1: false,
        nameEdit: '',
        surnameEdit: '',
        emailEdit: '',
        position: '',
        education: '',
        locationEdit: '',
        experiences: [
          {
            id: 1,
            position: 'Software Engineer',
            company: 'Company A',
            description: 'Responsible for developing web applications using Vue.js and Node.js.',
          },
          {
            id: 2,
            position: 'Front-end Developer',
            company: 'Company B',
            description: 'Developed responsive UI components using Bootstrap and jQuery.',
          },
        ],
        educations: [
          {
            id: 1,
            degree: 'Master of Science in Computer Science',
            university: 'University of XYZ',
          },
          {
            id: 2,
            degree: 'Bachelor of Engineering in Information Technology',
            university: 'ABC College',
          },
        ],
        skills: [
          { id: 1, name: 'JavaScript' },
          { id: 2, name: 'HTML' },
          { id: 3, name: 'CSS' },
        ],
        licenses: [
          { id: 1, name: 'License 1' },
          { id: 2, name: 'License 2' },
          { id: 3, name: 'License 3' },
        ],
      };
    },
    async mounted() {
      axios.get(`https://software.ardapektezol.com/api/profile`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        this.profile_data = res.data.data.user
        console.log("eee", this.profile_data);
      })
    },
    methods: {
  
      openModal() {
        // Set initial values for fields
        this.nameEdit = 'John';
        this.surnameEdit = 'Doe';
        this.email = 'johndoe@example.com';
        this.position = 'Software Engineer';
        this.education = 'Bachelor of Science in Computer Science';
        this.location = 'New York, USA';
        this.editedAbout = this.about;
        // Open the modal
        this.modalOpen = true; 
      }, 
      openModal1() {
        this.modalOpen1 = true;
      }, 
      saveChanges() {
        // Perform save or update logic here
        // Access the updated values from this.name, this.surname, etc.
        this.modalOpen = false;
        this.about = this.editedAbout;
        this.modalOpen1 = false;
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
  