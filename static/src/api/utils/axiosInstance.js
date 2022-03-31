import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: `http://${process.env.VUE_APP_API_HOST}:${process.env.VUE_APP_API_PORT}/v1`,
})

export default axiosInstance
