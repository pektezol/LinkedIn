import axios from "axios";

class ProfileService {
  getProfile(token) {
    return axios.get(`https://software.ardapektezol.com/api/profile`, {
      headers: {
        Authorization: token
      }
    })
  }
  updateProfile(payload, token) {
    return axios.post(`https://software.ardapektezol.com/api/profile`, {
      user_name: payload.userName,
      first_name: payload.firstName,
      last_name: payload.lastName,
      email: payload.email,
      password: payload.password,
      date_of_birth: payload.dateOfBirth
    }, {
      headers: {
        Authorization: token
      }
    })
  }
  connections(token) {
    return axios.get(`https://software.ardapektezol.com/api/connections`, { 
      headers: {
        Authorization: token
      }
    })
  }
  connectionRequest(user_name, token) {
    return axios.post(`https://software.ardapektezol.com/api/connectin/request`, {
      user_name: user_name, 
    }, {
      headers: {
        Authorization: token
      }
    })
  }
  connectionAccept(user_name, token) {
    return axios.post(`https://software.ardapektezol.com/api/connection/accept`, {
      user_name: user_name, 
    }, {
      headers: {
        Authorization: token
      }
    })
  }
  connectionIgnore(user_name, token) {
    return axios.post(`https://software.ardapektezol.com/api/connection/ignore`, {
      user_name: user_name, 
    }, {
      headers: {
        Authorization: token
      }
    })
  }
  notifications(token) {
    return axios.get(`https://software.ardapektezol.com/api/notification`, {
      headers: {
        Authorization: token
      }
    })
  }
}

export default new ProfileService();