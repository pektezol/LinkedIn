import axios from "axios";

class UserService {
  getUser(param) { 
    return axios.get(`https://software.ardapektezol.com/api/users/${param}`)
  }
  getProfile(payload) {
    return axios.post('https://software.ardapektezol.com/api/auth/login', {
        user_name: payload.userName, 
        password: payload.password
    })
  } 
}

export default new UserService();