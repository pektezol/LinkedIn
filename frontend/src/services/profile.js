import axios from "axios";

class ProfileService { 
  getProfile(token) {
    return axios.get(`https://software.ardapektezol.com/api/profile`, {
      headers: {
        Authorization: token
      }
    })
  }
}

export default new ProfileService();