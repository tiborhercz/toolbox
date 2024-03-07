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
          cols="10"
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
          cols="2"
          md="2"
        >
          <v-text-field
            v-model.number="indent"
            label="Indent"
            outlined
          />
        </v-col>
        <v-col
          cols="12"
          md="12"
        >
          <h2>Formatted JSON</h2>
          <json-formatter
            v-bind:value="inputValue"
            v-bind:indent="indent"
            v-on:formattedValue="setFormattedJson"
            v-on:errorMessage="setErrorMessage"
          />
          <basic-button
            v-bind:label="'copy'"
            v-bind:copy-value="formattedJson"
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
  name: 'JsonBeautify',
  components: {
    jsonFormatter,
    BasicButton,
  },
  props: {},
  data() {
    return {
      value: '',
      indent: 2,
      formattedJson: '{}',
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
    setFormattedJson(json) {
      this.formattedJson = json
    },
    setErrorMessage(errorMessage) {
      this.error = true
      this.errorMessages.push(errorMessage)
    },
  },
}
</script>
