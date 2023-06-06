<template>
  <div class="card">
    <div class="card-body text-center">
      <span><button class="mediatext btn btn-light mr-2" @click="toggleModal">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" width="24" height="24"
            focusable="false" fill="#70B5F9">
            <path
              d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm1 13a1 1 0 01-.29.71L16 14l-2 2-6-6-4 4V7a1 1 0 011-1h14a1 1 0 011 1zm-2-7a2 2 0 11-2-2 2 2 0 012 2z">
            </path>
          </svg>
          <span class="ml-1 mediatext">
            Fotoğraf
          </span>
        </button>
      </span>

      <span><button class="mediatext btn btn-light mr-2" @click="toggleModal">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill="" width="24"
            height="24" focusable="false">
            <path d="M21 3v2H3V3zm-6 6h6V7h-6zm0 4h6v-2h-6zm0 4h6v-2h-6zM3 21h18v-2H3zM13 7H3v10h10z">
            </path>
          </svg><span class="ml-1 mediatext">
            Yazı yaz
          </span>
        </button>
      </span>

      <b-modal ref="my-modal" hide-footer title="Create Post">

        <b-form @submit.prevent="onSubmit">

          <b-form-textarea id="textarea" v-model="post_form.text" placeholder="Write something..." rows="3"
            max-rows="6"></b-form-textarea>

          <b-form-file v-model="post_form.file" :state="Boolean(post_form.file)" placeholder="Choose a file or drop it here..."
            drop-placeholder="Drop file here..."></b-form-file>
          <div class="mt-3">Selected file: {{ post_form.file ? post_form.file.name : '' }}</div>

          <b-button class="mt-2" type="submit" variant="outline-success" block>Submit</b-button>
          <b-button class="mt-2" variant="outline-success" @click="encode_file()" block>aa</b-button>
        </b-form>
        <img :src="img" alt=""> 
      </b-modal>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios'; 
export default {
  name: "mainUpper",
  data() {
    return {
      post_form: {
        text: null,
        file: null
      },
      img: null
    }
  },
  methods: {
    encode_file() {
      const file = this.post_form.file
      const reader = new FileReader()
      let rawImg;
      reader.onloadend = () => {
        rawImg = reader.result;
        console.log(rawImg);
        this.img=rawImg
      }
      reader.readAsDataURL(file);
      this.post_form.file = file
      console.log(this.post_form.file) 
    },
    onSubmit() {
      axios.post(`https://software.ardapektezol.com/api/posts`, {
        text: this.post_form.text,
        image: this.post_form.file
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log(res);
        router.push("/network")
      })
      this.toggleModal()
    },
    toggleModal() {
      // We pass the ID of the button that we want to return focus to
      // when the modal has hidden
      this.$refs['my-modal'].toggle('#toggle-btn')
    }
  }
}
</script>
  