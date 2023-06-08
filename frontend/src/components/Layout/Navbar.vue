<template>
  <div class="">
    <!-- navbar -->
    <nav class="navbar navbar-expand-lg navbar-light navbarbg sticky-top d-flex mb-3 border-bottom">
      <div class="container">
        <a class="navbar-brand" href="#">
          <i class="fab fa-linkedin fa-lg" style="color: #0a66c2; font-size: 2.4rem"></i>
        </a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        </button>

        <form class="form-inline">
          <a href="/">
            <img src="https://cdn-icons-png.flaticon.com/512/174/174857.png" alt="" class="mr-2" style="width: 32px;">
          </a>
          <input class="form-control searche" type="search" placeholder="Arama Yap" aria-label="Search" v-model="search"/>
           
        </form>

        <div class="collapse navbar-collapse">
          <ul class="navbar-nav ml-auto mt-4">

            <li class="nav-item active">
              <a class="nav-link" href="/">
                <img class="transform translate-x-1" src="@/assets/images/nav-home.svg" alt="" srcset="">
                <p>Home</p>
              </a>
            </li>

            <li class="nav-item">
              <a class="nav-link" href="/network">
                <img class="transform translate-x-6 ml-4" src="@/assets/images/nav-network.svg" alt="" srcset="">
                <p>My Networks</p>
              </a>
            </li>

            <li class="nav-item">
              <a class="nav-link" href="/job">
                <img class="transform translate-x-px" src="@/assets/images/nav-jobs.svg" alt="" srcset="">
                <p>Jobs</p>
              </a>
            </li>

            <!-- <li class="nav-item">
              <a class="nav-link" href="#">
                <i class="fas fa-comment-dots fa-lg text-center">
                  <img class="transform translate-x-6" src="@/assets/images/nav-messaging.svg" alt="" srcset="">
                  <p>Messaging</p>
                </i>
              </a>
            </li> -->

            <!-- <li class="nav-item">
              <a class="nav-link" href="#" @click="search()">
                <img class="transform translate-x-6" src="@/assets/images/nav-notifications.svg" alt="" srcset="">
                <p> Notifications</p>
              </a>
            </li> -->

            <li class="nav-item mr-3">
              <a class="nav-link" href="/profile">
                 
                <div class="flex flex-wrap">
                  <div class="w-full text-center">
                    <!--             <img :src="user.data.photoURL" alt="" srcset=""  ref="btnRef" class="transform translate-x-px rounded-full w-7 h-7 ease-linear transition-all duration-150">
 -->
                    <p class="flex mt-4">
                      Me
                      <icon class="language-selector__label-chevron mx-1 lazy-loaded" aria-hidden="true">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" preserveAspectRatio="xMinYMin meet"
                          focusable="false" class="lazy-loaded">
                          <path
                            d="M8.8 10.66L14 5.12a.07.07 0 00-.07-.12H2.07a.07.07 0 00-.07.12l5.2 5.54a1.1 1.1 0 001.6 0z"
                            fill="currentColor"></path>
                        </svg>
                      </icon>
                    </p>
                    <!-- <div ref="popoverRef" v-bind:class="{'hidden': !popoverShow, 'block': popoverShow}" class="bg-white  border border-gray-100 block z-50 font-normal leading-normal text-sm max-w-xs text-left no-underline break-words rounded-lg">
              <div >
                <div class="text-gray-700 p-6 hover:underline" @click="">
                  Sign Out
                </div>
              </div>
            </div> -->
                  </div>
                </div>
              </a>
            </li> 
          </ul>
        </div>
      </div>
    </nav>
    <div v-if="search_toggle">
      <div class="modalss">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="row" >
              <button  @click="search_toggle = false"> 
                <b-icon icon="x" aria-hidden="true"></b-icon>
              </button>
            </div>
            <b-list-group> 
              <div v-for="user in search_users_data" :key="user.id"> 
                <b-list-group-item class="d-flex align-items-center" :href="`/user/${user.user_name}`">
                  <b-avatar variant="success" icon="people-fill" class="mr-3"></b-avatar>
                  <span class="mr-auto">{{ user.full_name }}  </span>
                </b-list-group-item>
              </div>
              <div v-for="user in search_companies_data" :key="user.id"> 
                <b-list-group-item class="d-flex align-items-center" :href="`/company/${user.name}/${user.id}`">
                  <b-avatar :src="user.logo" size="2rem"></b-avatar>
                  <span class="mr-auto ml-2">{{ user.name }}  </span>
                </b-list-group-item>
              </div>
            </b-list-group>
          </div>
        </div>
      </div>
    </div>  
  </div>
</template>
  
<script>
import axios from 'axios';
export default {
  name: 'Navbar',
  data() {
    return {
      search: "",
      search_toggle: false,
      popoverShow: true,
      id: null,
      search_users_data: [], 
      search_companies_data: []
    }
  },
  watch: {
    async search() { 
      this.search_func() 
    },
     
  },
  methods: {
    search_func() {
      
      axios.get(`https://software.ardapektezol.com/api/search`, {
        params: {
          q: this.search
        }
      }).then(res => {
        this.search_users_data = [] 
        for (let index = 0; index < res.data.data.users.length; index++) {
           this.search_users_data.push({full_name: `${res.data.data.users[index].first_name} ${res.data.data.users[index].last_name}`, user_name: `${res.data.data.users[index].user_name}`})
        } 
        if (res.data.data.users.length != 0) {
          this.search_toggle = true
        } else {
          this.search_toggle = false
        }
        this.search_companies_data = [] 
        for (let index = 0; index < res.data.data.companies.length; index++) {
           this.search_companies_data.push({name: `${res.data.data.companies[index].name}`, id: `${res.data.data.companies[index].id}`, logo: `${res.data.data.companies[index].logo}`})
        } 
        if (res.data.data.companies.length != 0) {
          this.search_toggle = true
        } else {
          this.search_toggle = false
        }
      })
    }
  }
}
</script>

<style >
.modalss {
  position: absolute;
  top: 30px;
  right: 30%;
  bottom: 0;
  left: 0%;
  z-index: 10040;
  overflow: auto;
  overflow-y: auto;
}
</style>