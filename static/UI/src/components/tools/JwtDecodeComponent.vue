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
            variant="outlined"
          />
        </v-col>
        <v-col>
          <v-btn
            v-on:click="fillExampleJwt"
          >
            Example JWT
          </v-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col
          cols="12"
          md="7"
        >
          <h2>JWT Header</h2>
          <json-formatter v-bind:value="headerValue" />
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
          <json-formatter v-bind:value="payloadValue" />
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
  name: 'JwtDecodeComponent',
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
      exampleJwt: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c',
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
    fillExampleJwt() {
      this.inputValue = this.exampleJwt
    },
    copyText(value) {
      navigator.clipboard.writeText(value)
    },
  },
}
</script>
