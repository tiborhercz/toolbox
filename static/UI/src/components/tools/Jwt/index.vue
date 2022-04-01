<template>
  <v-row>
    <v-col>
      <p>
        Get the header and payload from a JWT token
      </p>
      <v-row>
        <v-col
          cols="12"
          md="7"
        >
          <v-textarea
            v-model="inputValue"
            name="input-7-1"
            label="JWT token"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            outlined
          />
        </v-col>
        <v-col
          cols="12"
          md="7"
        >
          <h2>JWT Header</h2>
          <json-formatter v-bind:json="headerValue" />
          <basic-button
            v-bind:label="'copy'"
            v-bind:copy-value="headerValue"
          />
        </v-col>
        <v-col
          cols="12"
          md="7"
        >
          <h2>JWT Payload</h2>
          <json-formatter v-bind:json="payloadValue" />
          <basic-button
            v-bind:label="'copy'"
            v-bind:copy-value="payloadValue"
          />
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import BasicButton from '@/components/Basic/Button'
import JsonFormatter from '@/components/CodeFormatters/JsonFormatter'

export default {
  name: 'Jwt',
  components: {
    BasicButton,
    JsonFormatter,
  },
  props: {
    type: {
      default: '',
      type: String,
    },
  },
  data() {
    return {
      value: '',
      headerValue: '',
      payloadValue: '',
      error: false,
      errorMessages: [],
    }
  },
  computed: {
    inputValue: {
      get() {
        return this.value
      },
      set(newValue) {
        this.value = newValue
        this.headerValue = ''
        this.payloadValue = ''

        try {
          if (newValue !== '') {
            this.error = false
            this.errorMessages = []

            let data = wasmJwtDecode(newValue) // eslint-disable-line
            data = JSON.parse(data)

            this.headerValue = data.header
            this.payloadValue = data.payload
          }
        } catch (error) {
          this.error = true
          this.errorMessages.push('Invalid JWT token.')
        }
      },
    },
  },
  methods: {
    copyText(value) {
      navigator.clipboard.writeText(value)
    },
  },
}
</script>
