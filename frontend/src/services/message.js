import axios from "axios";

class MessageService {
  getUser(param) { 
    return axios.get(`https://software.ardapektezol.com/api/users/${param}`)
  } 
}

export default new MessageService();