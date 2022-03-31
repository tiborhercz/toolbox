import axios from './utils/axiosInstance'

export default {
  ipv4Cidr(ipv4Address) {
    const requestBody = {
      ipv4Address,
    }
    return axios.post('/ipv4/cidr', requestBody)
  },
}
