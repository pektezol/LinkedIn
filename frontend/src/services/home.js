import axios from "axios";

class HomeService {
  getPosts() { 
    return axios.get(`https://software.ardapektezol.com/api/post/`)
  } 
}

export default new HomeService();