<template>
  <!-- Right-Bottom-Mesajlaşma - Start -->
  <div class="">
    <div class="fixed">
      <nav class="navbar navbar-expand-xl navbar-light bg-white position-sticky fixed-bottom border ml-auto "
        style=" width: 300px; height: 50px; position: -webkit-sticky; border-radius: 5px;  right: 15px;">
        <b-button block variant="primary" @click="message_list_toggle = !message_list_toggle">Messages</b-button> 
      </nav>
    </div> 
    <div class="fixed2" v-if="message_list_toggle">
      <nav class="navbar navbar-expand-xl navbar-light bg-white position-sticky fixed-bottom border ml-auto ">
        <div class="row">
          <b-list-group
            style=" width: 298px; height: 400px; position: -webkit-sticky; border-radius: 5px;  right: 15px; "> 
            <div v-for="item in message_list.connections" :key="item.index"> 
              <b-list-group-item>
                <b-avatar class="mr-3"></b-avatar>
                <b-button variant="link text-dark" @click="select_user_for_message(item.sender.user_name)"><span
                    class="mr-auto">{{ item.sender.first_name }} {{
                      item.sender.last_name }}</span> </b-button>
              </b-list-group-item>
            </div> 
          </b-list-group>
        </div>
      </nav>
    </div> 
    <div class="fixed3" v-if="message_list_toggle && message_box"> 
      <div class="page-content page-container fixed3" id="page-content">
        <div class="padding">
          <div class="row container d-flex justify-content-center">
            <div style="width: 450px!important;">
              <div class="card card-bordered"> 
                <div class="card-header">
                  <h4 class="card-title"><strong>{{ message_user_name }}</strong></h4>
                </div> 
                <div class="ps-container ps-theme-default ps-active-y" id="chat-content"
                  style="overflow-y: scroll !important; height:400px !important;">
 
                  <div v-for="item in message_box_list" :key="item.id">
                    <div v-for="(message) in item.messages.slice().reverse()" :key="message.id">

                      <div v-if="message.sent == false"> 
                        <div class="media media-chat">
                          <div class="media-body">
                            <p>{{message.message}}</p>
                            <p class="meta"><time >{{ message.date.split("T")[1].slice(0, 5) }}</time></p>
                          </div>
                        </div>
                      </div>

                      <div v-else> 
                        <div class="media media-chat media-chat-reverse">
                          <div class="media-body">
                            <p>{{message.message}}</p>
                            <p class="meta"><time >{{ message.date.split("T")[1].slice(0, 5) }}</time></p>
                          </div>
                        </div>
                      </div>

                    </div> 
                  </div> 
                  <div class="ps-scrollbar-x-rail" style="left: 0px; bottom: 0px;">
                    <div class="ps-scrollbar-x" tabindex="0" style="left: 0px; width: 0px;"></div>
                  </div>
                  <div class="ps-scrollbar-y-rail" style="top: 0px; height: 0px; right: 2px;">
                    <div class="ps-scrollbar-y" tabindex="0" style="top: 0px; height: 2px;"></div>
                  </div>
                </div> 
                <div class="publisher bt-1 border-light"> 
                  <div class="row">
                    <div class="col-10">
                      <input class="publisher-input" type="text" placeholder="Write something" v-model="message_item">
                    </div>
                    <div class="col-2">
                      <b-button @click="message_send(message_user_name)" variant="outline-primary">Send</b-button>
                    </div>
                  </div> 
                </div> 
              </div>
            </div>
          </div>
        </div>
      </div> 
    </div> 
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "message",
  data() {
    return {
      message_list: null,
      message_box_list: null,
      message_box: false,
      message_list_toggle: true,
      user_messages: null,
      message_user_name: null,
      message_item: ""
    }
  },
  created() {
    this.get_message_list()
    
  },
  methods: {
    get_message_list() {
      axios.get(`https://software.ardapektezol.com/api/allconnections`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log(res.data.data);
        this.message_list = res.data.data
      })
    },
    select_user_for_message(data) {
      this.message_list_toggle = true,
        this.message_user_name = data
      axios.get(`https://software.ardapektezol.com/api/messages/${this.message_user_name}`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log(res.data.data);
        this.message_box_list = res.data.data
        this.message_box = true
         
      })
    },
    message_send(send_to) { 
      axios.post(`https://software.ardapektezol.com/api/messages/${send_to}`, {
        message: this.message_item
      }, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => { 
        console.log(res.data.data);
        this.user_messages = res.data.data
        this.message_box = true
        this.message_item = null
        this.select_user_for_message(this.message_user_name)
      })
    }
  }
}
</script>

<style>
.fixed {
  position: fixed;

  left: 80%;
  top: 95%;
  z-index: 2000;
}

.fixed2 {
  position: fixed;

  left: 80%;
  top: 50%;
  z-index: 2000;
  ;
}

.fixed3 {
  position: fixed;

  left: 55%;
  top: 38%;
  z-index: 2000;
}

.card-bordered {
  border: 1px solid #ebebeb;
}

.card {
  border: 0;
  border-radius: 0px;
  margin-bottom: 30px;
  -webkit-box-shadow: 0 2px 3px rgba(0, 0, 0, 0.03);
  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.03);
  -webkit-transition: .5s;
  transition: .5s;
}

.padding {
  padding: 3rem !important
}

body {
  background-color: #f9f9fa
}

.card-header:first-child {
  border-radius: calc(.25rem - 1px) calc(.25rem - 1px) 0 0;
}


.card-header {
  display: -webkit-box;
  display: flex;
  -webkit-box-pack: justify;
  justify-content: space-between;
  -webkit-box-align: center;
  align-items: center;

  background-color: transparent;
  border-bottom: 1px solid rgba(77, 82, 89, 0.07);
}

.card-header .card-title {
  padding: 0;
  border: none;
}

h4.card-title {
  font-size: 17px;
}

.card-header>*:last-child {
  margin-right: 0;
}

.card-header>* {
  margin-left: 8px;
  margin-right: 8px;
}

.btn-secondary {
  color: #4d5259 !important;
  background-color: #e4e7ea;
  border-color: #e4e7ea;
  color: #fff;
}

.btn-xs {
  font-size: 11px;
  padding: 2px 8px;
  line-height: 18px;
}

.btn-xs:hover {
  color: #fff !important;
}




.card-title {
  font-family: Roboto, sans-serif;
  font-weight: 300;
  line-height: 1.5;
  margin-bottom: 0;
  padding: 15px 20px;
  border-bottom: 1px solid rgba(77, 82, 89, 0.07);
}


.ps-container {
  position: relative;
}

.ps-container {
  -ms-touch-action: auto;
  touch-action: auto;
  overflow: hidden !important;
  -ms-overflow-style: none;
}

.media-chat {
  padding-right: 64px;
  margin-bottom: 0;
}

.media {
  padding: 16px 12px;
  -webkit-transition: background-color .2s linear;
  transition: background-color .2s linear;
}

.media .avatar {
  flex-shrink: 0;
}

.avatar {
  position: relative;
  display: inline-block;
  width: 36px;
  height: 36px;
  line-height: 36px;
  text-align: center;
  border-radius: 100%;
  background-color: #f5f6f7;
  color: #8b95a5;
  text-transform: uppercase;
}

.media-chat .media-body {
  -webkit-box-flex: initial;
  flex: initial;
  display: table;
}

.media-body {
  min-width: 0;
}

.media-chat .media-body p {
  position: relative;
  padding: 6px 8px;
  margin: 4px 0;
  background-color: #f5f6f7;
  border-radius: 3px;
  font-weight: 100;
  color: #9b9b9b;
}

.media>* {
  margin: 0 8px;
}

.media-chat .media-body p.meta {
  background-color: transparent !important;
  padding: 0;
  opacity: .8;
}

.media-meta-day {
  -webkit-box-pack: justify;
  justify-content: space-between;
  -webkit-box-align: center;
  align-items: center;
  margin-bottom: 0;
  color: #8b95a5;
  opacity: .8;
  font-weight: 400;
}

.media {
  padding: 16px 12px;
  -webkit-transition: background-color .2s linear;
  transition: background-color .2s linear;
}

.media-meta-day::before {
  margin-right: 16px;
}

.media-meta-day::before,
.media-meta-day::after {
  content: '';
  -webkit-box-flex: 1;
  flex: 1 1;
  border-top: 1px solid #ebebeb;
}

.media-meta-day::after {
  content: '';
  -webkit-box-flex: 1;
  flex: 1 1;
  border-top: 1px solid #ebebeb;
}

.media-meta-day::after {
  margin-left: 16px;
}

.media-chat.media-chat-reverse {
  padding-right: 12px;
  padding-left: 64px;
  -webkit-box-orient: horizontal;
  -webkit-box-direction: reverse;
  flex-direction: row-reverse;
}

.media-chat {
  padding-right: 64px;
  margin-bottom: 0;
}

.media {
  padding: 16px 12px;
  -webkit-transition: background-color .2s linear;
  transition: background-color .2s linear;
}

.media-chat.media-chat-reverse .media-body p {
  float: right;
  clear: right;
  background-color: #48b0f7;
  color: #fff;
}

.media-chat .media-body p {
  position: relative;
  padding: 6px 8px;
  margin: 4px 0;
  background-color: #f5f6f7;
  border-radius: 3px;
}


.border-light {
  border-color: #f1f2f3 !important;
}

.bt-1 {
  border-top: 1px solid #ebebeb !important;
}

.publisher {
  position: relative;
  display: -webkit-box;
  display: flex;
  -webkit-box-align: center;
  align-items: center;
  padding: 12px 20px;
  background-color: #f9fafb;
}

.publisher>*:first-child {
  margin-left: 0;
}

.publisher>* {
  margin: 0 8px;
}

.publisher-input {
  -webkit-box-flex: 1;
  flex-grow: 1;
  border: none;
  outline: none !important;
  background-color: transparent;
}

button,
input,
optgroup,
select,
textarea {
  font-family: Roboto, sans-serif;
  font-weight: 300;
}

.publisher-btn {
  background-color: transparent;
  border: none;
  color: #8b95a5;
  font-size: 16px;
  cursor: pointer;
  overflow: -moz-hidden-unscrollable;
  -webkit-transition: .2s linear;
  transition: .2s linear;
}

.file-group {
  position: relative;
  overflow: hidden;
}

.publisher-btn {
  background-color: transparent;
  border: none;
  color: #cac7c7;
  font-size: 16px;
  cursor: pointer;
  overflow: -moz-hidden-unscrollable;
  -webkit-transition: .2s linear;
  transition: .2s linear;
}

.file-group input[type="file"] {
  position: absolute;
  opacity: 0;
  z-index: -1;
  width: 20px;
}

.text-info {
  color: #48b0f7 !important;
}
</style>