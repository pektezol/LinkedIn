import axios from "axios";

class CompanyService {
  getCompany(param) { 
    return axios.get(`https://software.ardapektezol.com/api/company/${param}`)
  } 
  updateCompany(param) { 
    return axios.get(`https://software.ardapektezol.com/api/company/${param}`)
  }
  createCompany(param) { 
    return axios.get(`https://software.ardapektezol.com/api/company/${param}`)
  }
}

export default new CompanyService();