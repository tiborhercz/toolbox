<template>
  <v-row>
    <v-col
      cols="12"
    >
      <p>
        Calculate the IPv4 address information of a given CIDR range.
      </p>
      <v-row class="mb-5">
        <v-col
          cols="12"
          sm="6"
          md="3"
        >
          <v-autocomplete
            v-bind:items="items"
            v-bind:search-input.sync="searchValue"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            placeholder="10.0.0.0/24"
            outlined
            dense
            v-on:keydown.enter="getCidr()"
          />
          <v-btn
            v-on:click="getCidr()"
          >
            Send
          </v-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col
          cols="12"
          sm="6"
          md="6"
        >
          <v-simple-table>
            <template v-slot:default>
              <thead>
                <tr>
                  <th
                    width="50%"
                    class="text-left"
                  >
                    Name
                  </th>
                  <th
                    width="50%"
                    class="text-left"
                  >
                    Value
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in ipv4Data"
                  v-bind:key="item.name"
                >
                  <td>{{ item.name }}</td>
                  <td>{{ item.value }}</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import { validateIpv4Cidr, validIpv4 } from '@/utils'
import Ipv4Api from '@/api/ipv4'

const DEFAULT_ITEMS = ['0.0.0.0/8', '0.0.0.0/16', '0.0.0.0/24']

function DEFAULT_IPV4DATA() {
  return {
    subnetMask: {
      name: 'Subnetmask',
      value: null,
    },
    firstIp: {
      name: 'First ip address',
      value: null,
    },
    lastIp: {
      name: 'Last ip address',
      value: null,
    },
    totalIpAddresses: {
      name: 'Total ip addresses',
      value: null,
    },
  }
}

export default {
  name: 'Cidr',
  components: {},
  props: {},
  data() {
    return {
      ip: ['0', '0', '0', '0'],
      cidrValues: ['8', '16', '24'],
      items: ['0.0.0.0/8', '0.0.0.0/16', '0.0.0.0/24'],
      isValidCidrIp: false,
      value: null,
      error: false,
      errorMessages: [],
      ipv4Data: DEFAULT_IPV4DATA(),
    }
  },
  computed: {
    searchValue: {
      get() {
        return this.value
      },
      set(newValue) {
        this.error = false
        this.errorMessages = []
        this.ipv4Data = DEFAULT_IPV4DATA()

        if (newValue === '') {
          this.items = DEFAULT_ITEMS
          return
        }

        this.value = newValue
        this.items.push(newValue)

        if ((this.isIpPartiallyValid(newValue)
            || (validIpv4(newValue)) || validateIpv4Cidr(newValue))
          && this.isCidrValid(newValue)
        ) {
          this.setIp(newValue)
          this.generateItems(newValue)
        }

        if (validateIpv4Cidr(newValue)) {
          this.getCidr()
        }
      },
    },
  },
  methods: {
    async getCidr() {
      if (validateIpv4Cidr(this.value)) {
        const { data } = await Ipv4Api.ipv4Cidr(this.value)
        this.ipv4Data.subnetMask.value = data.data.subnetMask
        this.ipv4Data.firstIp.value = data.data.firstIp
        this.ipv4Data.lastIp.value = data.data.lastIp
        this.ipv4Data.totalIpAddresses.value = data.data.totalIpAddresses
      } else {
        this.error = true
        this.errorMessages.push(`${this.value} is not a valid ipv4 address with cidr. Example: 10.0.0.0/24`)
      }
    },
    isCidrValid(value) {
      const slashCount = (value.match(/\//g) || []).length
      if (slashCount > 1) {
        return false
      }

      return true
    },
    isIpPartiallyValid(value) {
      if (value === null) {
        return false
      }

      const dotCount = (value.match(/\./g) || []).length
      if (dotCount > 3) {
        return false
      }

      const parts = value.split('.')
      const lastIndex = parts.length - 1

      if (parts[lastIndex] === '') {
        parts.pop()
      }

      if (parts.length > 4) {
        return false
      }

      if (!parts.every(this.validIpPart)) {
        return false
      }

      return true
    },
    setIp(value) {
      const parts = value.split('.')

      const maxI = 3
      for (let i = 0; i < parts.length && i <= maxI; i++) {
        const part = parts[i]
        if (i === maxI && part.indexOf('/') > -1) {
          this.ip[i] = part.substring(0, part.indexOf('/'))
        } else {
          this.ip[i] = part
        }
      }
    },
    validIpPart(value) {
      if (value.length > 3 || value.length === 0) {
        return false
      }

      if (!value.match(/^\d+$/)) {
        return false
      }

      if (Number(value) < 0 || Number(value) > 255) {
        return false
      }

      return true
    },
    generateItems() {
      const items = []
      this.cidrValues.forEach((elm) => {
        const cidrIp = `${this.ip.join('.')}/${elm}`
        items.push(cidrIp)
      })

      this.items = items
    },
  },
}
</script>
