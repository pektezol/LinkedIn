import axios from "axios";

class UserService {
  getUser(param) { 
    return axios.get(`https://software.ardapektezol.com/api/users/${param}`)
  }
}

export default new UserService();