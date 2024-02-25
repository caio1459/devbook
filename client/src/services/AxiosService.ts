import axios, { AxiosInstance } from "axios";

class AxiosService {
  makeRequest(): AxiosInstance {
    return axios.create({
      baseURL: "http://localhost:5000/api",
    });
  }
}

export default new AxiosService();
