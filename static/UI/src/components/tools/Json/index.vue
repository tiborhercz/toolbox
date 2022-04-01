<template>
  <v-row>
    <v-col
      cols="12"
    >
      <p>
        Beautify JSON
      </p>
      <v-row>
        <v-col
          cols="12"
          md="7"
        >
          <v-textarea
            v-model="inputValue"
            name="input-7-1"
            label="JSON"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            outlined
          />
        </v-col>
        <v-col
          cols="12"
          md="12"
        >
          <h2>Formatted JSON</h2>
          <json-formatter
            v-bind:json="inputValue"
            v-on:errorMessage="setErrorMessage"
          />
          <basic-button
            v-bind:label="'copy'"
            v-bind:copy-value="inputValue"
          />
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import BasicButton from '@/components/Basic/Button'
import jsonFormatter from '@/components/CodeFormatters/JsonFormatter'

export default {
  name: 'Json',
  components: {
    jsonFormatter,
    BasicButton,
  },
  props: {},
  data() {
    return {
      value: '',
      error: false,
      errorMessages: [],
    }
  },
  computed: {
    inputValue: {
      get() {
        return this.value
      },
      async set(newValue) {
        this.value = newValue
        this.error = false
        this.errorMessages = []
      },
    },
  },
  methods: {
    setErrorMessage(errorMessage) {
      this.error = true
      this.errorMessages.push(errorMessage)
    },
  },
}
</script>
