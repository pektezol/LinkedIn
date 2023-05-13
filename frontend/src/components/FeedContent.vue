<template>
  <div class="feed_content bg-gray-50 flex justify-around">
    <!-- **LEFT CONTENT** -->
    <FeedLeftSide/>
    <!--** CENTER ** -->
    <section class="feed_content-center flex flex-col">
      <div class="feed_content-center--post flex flex-col border-2 border bg-white rounded-lg -ml-40 px-6 h-28 py-3">
        <div class="flex justify-between mb-2">
          <img class="w-12 h-12 mr-2 rounded-full" :src="user.data.photoURL" alt="" srcset="">
          <form @submit.prevent>
            <input @click="toggleInputModal()" class="feed_content-center--searchbar w-96 px-4 outline-none border border-gray-200 rounded-3xl h-12 bg-white-100 " type="text" placeholder="Start a post" v-model="post">
            <button type="submit" class="hidden">Send</button>
          </form>
          <!-- Modal Section -->
          <div class="post-modal">
            <div v-if="showInputModal" class="overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none ">
              <div class="relative w-auto my-16 mx-auto max-w-xl">
                <!--content-->
                <div class="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                  <!--modal header-->
                  <div class="flex items-start justify-between  border-b border-solid border-blueGray-200 rounded-t">
                    <h3 class="text-xl mt-3 ml-2 text-gray-700">
                      Create a post
                    </h3>
                    <button class=" border-0 text-black float-right text-3xl leading-none outline-none focus:outline-none" v-on:click="toggleInputModal()">
                      <span class="text-black h-12 w-6 text-3xl flex outline-none focus:outline-none">
                        ×
                      </span>
                    </button>
                  </div>
                  <!--modal body-->
                  <div class="relative p-5 flex flex-col">
                    <div class="flex">
                      <img class="w-12 h-12 mr-2  rounded-full" :src="user.data.photoURL" alt="" srcset="">
                      <div class="feed_content-center--posts-info">
                        <h3 class="feed_content-center--posts-name text-md font-medium text-gray-900">{{user.data.displayName}}</h3>
                        <p class="feed_content-center--posts-description  text-xs">{{user.data.email}}</p>
                      </div>
                    </div>
                    <form @submit.prevent='submitPost(), toggleInputModal()'>
                      <input class="feed_content-center--text-area w-96 mb-16 px-4 outline-none h-12 bg-white-100 text-black-700 "
                      type="text" placeholder="What do you want to talk about ?" v-model="post">
                    </form>
                    <div v-if="imageData">
                      <img :src="picture" alt="preview" class="object-cover h-96 w-full">
                    </div>
                    <div v-else>
                      <img src="" class="object-cover invisible">
                    </div>
                  </div>
                  <!--modal footer-->
                  <div class="flex items-center justify-between p-2 border-t border-solid border-blueGray-200 rounded-b">
                    <ul class="flex justify-between mt-2">
                      <li class="flex ml-4  cursor-pointer" v-on:click="toggleModal(), toggleInputModal()">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#4B5563" class="mercado-match" width="24" height="24" focusable="false">
                          <path d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm1 13a1 1 0 01-.29.71L16 14l-2 2-6-6-4 4V7a1 1 0 011-1h14a1 1 0 011 1zm-2-7a2 2 0 11-2-2 2 2 0 012 2z"></path>
                        </svg>
                        <p  class="ml-2">Photo</p>
                      </li>
                      <li class="flex ml-4 cursor-pointer">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#4B5563" class="mercado-match" width="24" height="24" focusable="false">
                          <path d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm-9 12V8l6 4z"></path>
                        </svg>
                        <p  class="ml-2">Video</p>
                      </li>
                      <li class="flex ml-4 cursor-pointer">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#4B5563" class="mercado-match" width="24" height="24" focusable="false">
                          <path d="M3 3v15a3 3 0 003 3h12a3 3 0 003-3V3zm13 1.75A1.25 1.25 0 1114.75 6 1.25 1.25 0 0116 4.75zm-8 0A1.25 1.25 0 116.75 6 1.25 1.25 0 018 4.75zM19 18a1 1 0 01-1 1H6a1 1 0 01-1-1V9h14zm-5.9-3a1 1 0 00-1-1H12a3.12 3.12 0 00-1 .2l-1-.2v-3h3.9v1H11v1.15a3.7 3.7 0 011.05-.15 1.89 1.89 0 012 1.78V15a1.92 1.92 0 01-1.84 2H12a1.88 1.88 0 01-2-1.75 1 1 0 010-.25h1a.89.89 0 001 1h.1a.94.94 0 001-.88z"></path>
                        </svg>
                        <p  class="ml-2">Event</p>
                      </li>
                      <li class="flex ml-4 cursor-pointer">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#4B5563" class="mercado-match" width="24" height="24" focusable="false">
                          <path d="M21 3v2H3V3zm-6 6h6V7h-6zm0 4h6v-2h-6zm0 4h6v-2h-6zM3 21h18v-2H3zM13 7H3v10h10z"></path>
                        </svg>
                        <p class="ml-2">Write Article</p>
                      </li>
                    </ul>
                    <button class="text-white bg-blue-700 border rounded-3xl border-solid border-blue-700 hover:border-blue-700
                      hover:bg-blue-900  active:bg-blue-600
                      lowercase text-lg font-medium px-4 py-0.5 outline-none focus:outline-none mr-2 mb-1 ease-linear transition-all duration-150"
                      type="button" v-on:click="submitPost(), toggleInputModal()">
                      Post
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div>
          <ul class="flex justify-between mt-2">
            <li class="flex  cursor-pointer" v-on:click="toggleInputModal()">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#70b5f9" class="mercado-match" width="24" height="24" focusable="false">
                <path d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm1 13a1 1 0 01-.29.71L16 14l-2 2-6-6-4 4V7a1 1 0 011-1h14a1 1 0 011 1zm-2-7a2 2 0 11-2-2 2 2 0 012 2z"></path>
              </svg>
              <p  class="ml-2">Photo</p>
            </li>
            <li class="flex cursor-pointer">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#7fc15e" class="mercado-match" width="24" height="24" focusable="false">
                <path d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm-9 12V8l6 4z"></path>
              </svg>
              <p  class="ml-2">Video</p>
            </li>
            <li class="flex cursor-pointer">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#e7a33e" class="mercado-match" width="24" height="24" focusable="false">
                <path d="M3 3v15a3 3 0 003 3h12a3 3 0 003-3V3zm13 1.75A1.25 1.25 0 1114.75 6 1.25 1.25 0 0116 4.75zm-8 0A1.25 1.25 0 116.75 6 1.25 1.25 0 018 4.75zM19 18a1 1 0 01-1 1H6a1 1 0 01-1-1V9h14zm-5.9-3a1 1 0 00-1-1H12a3.12 3.12 0 00-1 .2l-1-.2v-3h3.9v1H11v1.15a3.7 3.7 0 011.05-.15 1.89 1.89 0 012 1.78V15a1.92 1.92 0 01-1.84 2H12a1.88 1.88 0 01-2-1.75 1 1 0 010-.25h1a.89.89 0 001 1h.1a.94.94 0 001-.88z"></path>
              </svg>
              <p  class="ml-2">Event</p>
            </li>
            <li class="flex cursor-pointer">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="#f5987e" class="mercado-match" width="24" height="24" focusable="false">
                <path d="M21 3v2H3V3zm-6 6h6V7h-6zm0 4h6v-2h-6zm0 4h6v-2h-6zM3 21h18v-2H3zM13 7H3v10h10z"></path>
              </svg>
              <p class="ml-2">Write Article</p>
            </li>
            <!-- Modal Section -->
            <div class="choose-image-modal">
              <div v-if="showModal" class="overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none ">
                <div class="relative w-auto my-16 mx-auto max-w-lg">
                  <!--content-->
                  <div class="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                    <!--header-->
                    <div class="flex items-start justify-between p-3 border-b border-solid border-blueGray-200 rounded-t">
                      <h3 class="text-xl text-gray-700">
                        Edit your photo
                      </h3>
                      <button class="p-1 border-0 text-black float-right text-3xl leading-none outline-none focus:outline-none" v-on:click="toggleModal()">
                        <span class="text-black h-12 w-6 text-3xl flex outline-none focus:outline-none">
                          ×
                        </span>
                      </button>
                    </div>
                    <!--body-->
                    <div class="relative p-10 flex-auto">
                      <div class="flex flex-col items-center justify-center">
                        <input type="file" accept="image/*" ref="fileInput" class="invisible" @change="previewImage">
                        <button @click="chooseFile" class="items-center border-0 px-2.5 py-1  text-base
                         text-center text-blue-700 rounded font-medium hover:bg-blue-100 hover:rounded  cursor-pointer"
                         for="image-upload">Select images to share</button>
                        <label for="image-upload"></label>
                      </div>
                    </div>
                    <!--footer-->
                    <div class="flex items-center justify-end p-2 border-t border-solid border-blueGray-200 rounded-b">
                      <button class="text-blue-700 bg-transparent border rounded-3xl border-solid border-blue-700 hover:border-blue-700
                      hover:bg-blue-100 hover:text-blue active:bg-blue-600
                      lowercase text-lg font-medium px-4 py-0.5 outline-none focus:outline-none mr-2 mb-1 ease-linear transition-all duration-150"
                      type="button" v-on:click="toggleModal()">
                        cancel
                      </button>
                      <button  class="text-white  rounded-3xl bg-blue-700 font-medium
                        px-5 py-2 text-sm outline-none focus:outline-none mr-2 mb-1 ease-linear transition-all duration-150"
                       type="button" v-on:click="toggleModal(), toggleInputModal()">
                        Done
                      </button>
                    </div>
                  </div>
                </div>
              </div>
              <div v-if="showModal" class="opacity-25 fixed inset-0 z-40 bg-black"></div>
            </div>
          </ul>
        </div>
      </div>
      <div class="feed_content-center--posts">
        <li class=" flex -ml-40 border border-2 rounded-2xl bg-white py-4 mt-10 w-auto "
          v-for="(item, index) in posts" :key="index">
          <img class="w-12 h-12 mr-2 ml-4 rounded-full" :src="item.photoURL" alt="display picture">
          <div class="feed_content-center--posts-info">
            <h3 class="feed_content-center--posts-name text-md font-medium text-gray-700 px-2">{{item.displayName}}</h3>
            <p class="feed_content-center--posts-description  text-xs px-2">{{item.email}}</p>
            <p class="feed_content-center--posts-message font-normal text-black-50 mt-4 mb-4  text-justify items-start w-96 h-auto flex-wrap">{{item.posts}} </p>
            <div class="feed_content-center--posts-image" v-if="item.photo">
              <img :src="item.photo" alt="preview" srcset="" class="object-cover h-96 w-96 rounded-md mb-4">
            </div>
            <div class="feed_content-center--posts-image" v-else>
              <img src="" class="object-cover invisible">
            </div>
          </div>
        </li>
      </div>
    </section>
    <FeedRightSide/>
  </div>
</template>

<script>
import FeedRightSide from '@/components/FeedRightSide.vue'
import FeedLeftSide from '@/components/FeedLeftSide.vue'
import 'firebase/storage'
// import router from '@/router/index'
export default {
  name: 'regular-modal',
  components: {
    FeedRightSide,
    FeedLeftSide
  },
  data () {
    return {
      post: '',
      posts: [
        { displayName: 'Toluwanimi Amole', photoURL: 'https://lh3.googleusercontent.com/a-/AOh14GjsDcVtdLcVnkqnS_U8NVGg2wt-VNCn29PqnGqThw=s96-c', email: 'amolejesunifemi@gmail.com', posts: 'Nice work' }
      ],
      showModal: false,
      showInputModal: false,
      picture: '',
      imageData: null,
      // uploadValue: null
      // userPost: [],
      user: {
        data: {
          photoURL: 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200',
          displayName: 'VolkanO6',
          email: 'volkanoztoklu@gmail.com'
        }
      }
    }
  },

  methods: {
    chooseFile () {
      this.$refs.fileInput.click()
    },
    toggleModal () {
      this.showModal = !this.showModal
    },
    toggleInputModal () {
      this.showInputModal = !this.showInputModal
    },
    submitPost () {},
    previewImage () {},
    onUpload () {}
  },
  created () {}
}
</script>
