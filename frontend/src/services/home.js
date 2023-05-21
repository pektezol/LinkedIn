import axios from "axios";

class HomeService {
  getPosts(token) { 
    return axios.get(`https://software.ardapektezol.com/api/post/`, {
      headers: {
        Authorization: token
      }
    })
  } 
  /* updatePosts(payload, token) { 
    return axios.post(`https://software.ardapektezol.com/api/post/update`, {
      
    }, {
      headers: {
        Authorization: token
      }
    })
  }  */
  deletePosts(id, token) { 
    return axios.post(`https://software.ardapektezol.com/api/post/delete`,{
      post_id: id
    }, {
      headers: {
        Authorization: token
      }
    })
  } 
}

export default new HomeService();