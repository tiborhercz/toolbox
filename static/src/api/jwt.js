import axios from './utils/axiosInstance'

const jwtPrefix = '/jwt'

export default {
  jwtDecode(value) {
    const requestBody = {
      value,
    }
    return axios.post(`${jwtPrefix}/decode`, requestBody)
  },
}
