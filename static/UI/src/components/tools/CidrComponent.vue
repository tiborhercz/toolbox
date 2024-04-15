<template>
  <v-row>
    <v-col
      cols="12"
    >
      <p class="mb-2">
        Calculate the IPv4 address information of a given CIDR range.
      </p>
      <v-row class="mb-5">
        <v-col
          cols="12"
          sm="6"
          md="3"
        >
          <v-autocomplete
            v-model:search="searchValue"
            v-bind:items="items"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            placeholder="10.0.0.0/24"
            variant="outlined"
            density="compact"
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
          <v-table>
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
          </v-table>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { validateIpv4Cidr, validIpv4 } from '@/utils'

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

const ipAddress = reactive(['0', '0', '0', '0'])
const cidrValues = reactive(['8', '16', '24'])
const items = ref(['0.0.0.0/8', '0.0.0.0/16', '0.0.0.0/24'])
const value = ref('')
const error = ref(false)
const errorMessages = ref([])
const ipv4Data = ref(DEFAULT_IPV4DATA())

const setIp = (ip) => {
  const parts = ip.split('.')

  const maxI = 3
  for (let i = 0; i < parts.length && i <= maxI; i++) {
    const part = parts[i]
    if (i === maxI && part.indexOf('/') > -1) {
      ipAddress[i] = part.substring(0, part.indexOf('/'))
    } else {
      ipAddress[i] = part
    }
  }
}

const isCidrValid = (cidrValue) => {
  const slashCount = (cidrValue.match(/\//g) || []).length
  if (slashCount > 1) {
    return false
  }

  return true
}

const validIpPart = (ipPart) => {
  if (ipPart.length > 3 || ipPart.length === 0) {
    return false
  }

  if (!ipPart.match(/^\d+$/)) {
    return false
  }

  if (Number(ipPart) < 0 || Number(ipPart) > 255) {
    return false
  }

  return true
}

const isIpPartiallyValid = (ip) => {
  if (ip === null) {
    return false
  }

  const dotCount = (ip.match(/\./g) || []).length
  if (dotCount > 3) {
    return false
  }

  const parts = ip.split('.')
  const lastIndex = parts.length - 1

  if (parts[lastIndex] === '') {
    parts.pop()
  }

  if (parts.length > 4) {
    return false
  }

  if (!parts.every(validIpPart)) {
    return false
  }

  return true
}

const generateItems = () => {
  const localItems = []
  cidrValues.forEach((elm) => {
    const cidrIp = `${ipAddress.join('.')}/${elm}`
    localItems.push(cidrIp)
  })

  items.value = localItems
}

const getCidr = () => {
  console.log(value.value)
  if (validateIpv4Cidr(value.value)) {
    let data = wasmIpv4Process(value.value) // eslint-disable-line
    console.log(data)
    data = JSON.parse(data)

    ipv4Data.value.subnetMask.value = data.subnetMask
    ipv4Data.value.firstIp.value = data.firstIp
    ipv4Data.value.lastIp.value = data.lastIp
    ipv4Data.value.totalIpAddresses.value = data.totalIpAddresses
  } else {
    error.value = true
    errorMessages.value.push(`${value.value} is not a valid ipv4 address with cidr. Example: 10.0.0.0/24`)
  }
}

const searchValue = computed({
  get() {
    return value.value
  },
  set(newValue) {
    error.value = false
    errorMessages.value = []
    ipv4Data.value = DEFAULT_IPV4DATA()

    if (newValue === '') {
      items.value = DEFAULT_ITEMS
      value.value = ''
      return
    }

    value.value = newValue
    items.value.push(newValue)

    if ((isIpPartiallyValid(newValue)
        || (validIpv4(newValue)) || validateIpv4Cidr(newValue))
      && isCidrValid(newValue)
    ) {
      setIp(newValue)
      generateItems(newValue)
    }

    if (validateIpv4Cidr(newValue)) {
      getCidr()
    }
  },
})
</script>
