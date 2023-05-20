import axios from "axios";

class AuthService {
  user() {
    console.log("girdiii");
    return axios.get(`https://software.ardapektezol.com/api/users/pektezol`)
  }
  userLogin(payload) {
    return axios.post('https://software.ardapektezol.com/api/auth/login', {
        user_name: payload.userName,
        first_name: payload.firstName,
        last_name: payload.lastName,
        email: payload.email,
        password: payload.password,
        date_of_birth: payload.dateOfBirth
    })
  }
  userRegister(payload) {
    return axios.post('https://software.ardapektezol.com/api/auth/register', {
        user_name: payload.userName,
        first_name: payload.firstName,
        last_name: payload.lastName,
        email: payload.email,
        password: payload.password,
        date_of_birth: payload.dateOfBirth
    })
  }
}

export default new AuthService();