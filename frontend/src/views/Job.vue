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
                <p class="card-text company-info">{{ job.company }}</p>
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
            <p class="card-text company-info">{{ selectedJob.company }}</p>
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
export default {
  data() {
    return {
      jobs: [
        {
          id: 1,
          title: "Senior Embedded Software Developer",
          employee: "Backend Developer",
          company: "Siemens Company",
          location: "San Francisco",
          logo: "path/to/logo1.png",
          description: "For more than 35 years, ground breaking technologies and business models engineered and developed by Siemens Turkey R&D department contributes to the global success of Siemens. Our R&D activities focused on the fields of electrification, automation and digitalization, provides end-to-end support to our global customers on their unique digital transformation journey. We aim to increase our strength on solution and product development with world class SW development experts, with the vision to be the digitalization partner for our global customers. We need game changers like you, to bring the next level of smart manufacturing and infrastructure from mere concept to reality, develop tomorrow’s smart cities, and address the most critical digitalization challenges across a comprehensive range of industries.",
          requirements: "BSc., MSc. Or PhD in Computer Science, Electronics, Control Engineering or related areas. Proven 5+ years of experience in embedded software development. Strong knowledge in C / C++ Strong knowledge in object-oriented programming, analysis, and design. Experience in development of industrial communication network protocols, tools, and libraries (Switching & Routing Protocols, Netlink, Netfilter, Iproute2, Quagga, ConfD, Netconf, SNMP, etc.) is a must. Development experience in Linux environment is a must Experience with creating and building custom Linux OS (Yocto, Windriver, Isar, etc.). Experience in Linux kernel driver development and U-boot. Knowledge and experience in computer architecture is a plus. Knowledge and experience in Real-time concepts and RTOS are a plusKnowledge in Switching ASICs (Marvell, Broadcom, Mellanox, . etc.) is a plus. Knowledge and experience in virtualization technologies (Linux KVM, Libvirt, VMware, VirtualBox, Vagrant, etc.) is a plus. Experience with source control systems (GIT, ClearCase etc.) Very good command of English.",
          responsibilities: "What We Offer: Speak up Culture, Respectful Workplace,Being part of a global work environment,Remote working and remote living flexibility,Attractive remuneration package,Excellent recognition tools providing spot awards,Learning & Development opportunities for both personal and professional growth,Leave days for parents and a variety of flexible working models that allow time off for yourself and your family,Creche allowance for mothers,Share matching programs to become a shareholder of Siemens AG,Find more benefits here,",
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
  methods: {
    selectJob(job, index) {
      this.selectedJob = job;
    },
    applyJob(job) {
      // Implement your apply job functionality here
      console.log("Applying for job:", job.title);
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