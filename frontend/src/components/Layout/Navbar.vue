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
          <input class="form-control searche" type="search" placeholder="Arama Yap" aria-label="Search" />
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
                <img class="transform translate-x-6" src="@/assets/images/nav-network.svg" alt="" srcset="" >
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

            <li class="nav-item">
              <a class="nav-link" href="#">
                <img class="transform translate-x-6" src="@/assets/images/nav-notifications.svg" alt="" srcset="">
                <p> Notifications</p>
              </a>
            </li>

            <li class="nav-item mr-3">
              <a class="nav-link" href="/profile">
                <img class="rounded-circle ml-2 text-center"
                  src="https://media-exp1.licdn.com/dms/image/C4D03AQFzRzfBHzBpbA/profile-displayphoto-shrink_100_100/0/1661208519486?e=1669248000&v=beta&t=F_N3_KlkhDP4RbePX7fUwfN9TtrmwSAdb-1OW3VvUA4"
                  style="width: 25px; height: 25px;" />
                <div class="flex flex-wrap">
                  <div class="w-full text-center">
                    <!--             <img :src="user.data.photoURL" alt="" srcset=""  ref="btnRef" class="transform translate-x-px rounded-full w-7 h-7 ease-linear transition-all duration-150">
 -->
                    <p class="flex" >
                      Me
                      <icon @click="" class="language-selector__label-chevron mx-1 lazy-loaded" aria-hidden="true">
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

            <li class="nav-item border-left h-25">
              <a class="nav-link ml-4 mr-2" href="#"><i class="fas fa-th fa-lg text-center">
                  <div class="menutext">
                    İş
                  </div>
                </i>
              </a>
            </li>

            <li class="nav-item">
              <a class="nav-link" href="#"><i class="fas fa-chalkboard fa-lg text-center">
                  <div class="menutext">
                    İş İlanı Yayınla
                  </div>
                </i>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </div>
</template>
  
<script>
export default {
  name: 'Navbar',
  data() {
    return {
      popoverShow: true
    }
  },

}
</script>

<style ></style>