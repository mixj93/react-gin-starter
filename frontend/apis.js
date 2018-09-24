import axios from 'axios';

const baseUrl = process.env.NODE_ENV === 'development' ? 'http://localhost:5678/api/' : '/api/'

axios.defaults.baseURL = baseUrl;

axios.interceptors.response.use(function(response) {
  return response.data;
}, function(error) {
  console.error(error);
  return Promise.reject(error);
})

export const getPing = () => {
  return axios.get(`/ping`);
}
