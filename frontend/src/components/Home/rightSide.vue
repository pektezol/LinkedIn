

  <template>
    <div>
       <div class="card-header" style="font-size: 20px">
      <span>Notifications</span>
      <span>
        <i class="fas fa-info-circle float-right">
        </i>
      </span>

    </div>
      <div class="card" style="width: 19rem">
    <div v-for="notification in orderedNotifications" :key="notification.id" class="notification-item">
      <div class="notification-icon">
        <i class="fas fa-bell"></i>
      </div>
      <div class="notification-content">
        <div class="notification-title">{{ notification.notification }}</div>
        <div class="notification-message">{{ notification.date.split("T")[0] }}</div>
      </div>
    </div>
  </div>

  <div class="text-center sticky-top pt-5" style="width: 19rem;">
    <ul class="list-inline">
      <li class="list-inline-item mt-3" style=" margin-top: 5rem;">
        <a href="https://about.linkedin.com/" class="link">
          Hakkında
        </a>
      </li>
      <li class="list-inline-item">
        <a href="https://www.linkedin.com/accessibility?lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3BT6H7rzDzR%2FSQEvO%2BuBKdSw%3D%3D&licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-compactfooter.accessibility"
          class="link">
          Erişilebilirlik
        </a>
      </li>
      <li class="list-inline-item">
        <a href="https://www.linkedin.com/help/linkedin?trk=footer_d_flagship3_feed&lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3BT6H7rzDzR%2FSQEvO%2BuBKdSw%3D%3D&licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-compactfooter.help"
          class="link">
          Yardım Merkezi
        </a>
      </li>
      <li class="list-inline-item">
        <a href="" class="link">
          Gizlilik ve Koşullar
        </a>
      </li>
      <li class="list-inline-item">
        <a href="https://www.linkedin.com/help/linkedin/answer/62931?lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3BT6H7rzDzR%2FSQEvO%2BuBKdSw%3D%3D&licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-compactfooter.ad_choices"
          class="link">
          Reklam Tercihleri
        </a>
      </li>
      <li class="list-inline-item">
        <a href="https://www.linkedin.com/ad/start?trk=n_nav_ads_rr&lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3BT6H7rzDzR%2FSQEvO%2BuBKdSw%3D%3D&licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-compactfooter.advertising"
          class="link">
          Reklam
        </a>
      </li>
    </ul>
    <div class="text-center" style="margin-top: 10px">
      <img src="https://www.logo.wine/a/logo/LinkedIn/LinkedIn-Logo.wine.svg" width="90px" height="20px" /><span
        style="font-size: 12px; margin-left: -15px"> © Linkedin Corporation |
        2022</span>
    </div>
  </div>
    </div>
</template>
  

<script>
import axios from "axios";
export default {
  data() {
    return {
      notifications: [],
      length_of_notifications: null
    };
  },
  mounted(){
    this.get_notifications()
  },
  computed: {
    orderedNotifications() {
      return this.notifications.slice().reverse();
    },
  },
  methods: {
    get_notifications(){
      axios.get(`https://software.ardapektezol.com/api/notifications`, {
        headers: {
          Authorization: this.$cookies.get("token")
        }
      }).then(res => {
        console.log(res.data);
        this.notifications = res.data.data
        this.length_of_notifications = res.data.data.length
        if (this.length_of_notifications >= 5) {
          for (let index = 5; index < this.length_of_notifications; index++) {
            this.notifications.pop() 
          }
        }
      })
    }
  }
};
</script>

<style scoped>




.notification-container {
  position: fixed;
  bottom: 20px;
  right: 20px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  z-index: 9999;
}

.notification-item {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
}

.notification-icon {
  margin-right: 10px;
}

.notification-title {
  font-weight: bold;
}

.notification-message {
  color: #666;
}
</style>