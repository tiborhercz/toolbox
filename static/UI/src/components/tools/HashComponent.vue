<template>
  <v-row>
    <v-col>
      <p>
        Hash a value with the selected algorithm
      </p>
      <v-row>
        <v-col
          cols="12"
          md="7"
        >
          <v-textarea
            v-model="value"
            name="input-7-1"
            label="Input"
            variant="outlined"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            v-model="selectedHashAlgorithm"
            v-bind:items="hashAlgorithmsReversed"
            label="Select an algorithm"
            variant="outlined"
          />
        </v-col>
        <v-col
          cols="12"
          md="7"
        >
          <v-textarea
            v-model="outputValue"
            name="input-7-1"
            v-bind:label="selectedHashAlgorithm + ' Output'"
            readonly
            variant="outlined"
          />
          <v-btn
            v-on:click="copyText(outputValue)"
          >
            Copy
          </v-btn>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'Hash',
  components: {},
  props: {
    type: {
      default: '',
      type: String,
    },
  },
  data() {
    return {
      selectedHashAlgorithm: 'SHA512',
      hashAlgorithms: [],
      value: '',
      outputValue: '',
      error: false,
      errorMessages: [],
    }
  },
  computed: {
    hashAlgorithmsReversed() {
      return this.hashAlgorithms.slice(0).reverse()
    },
  },
  watch: {
    selectedHashAlgorithm() {
      this.hash()
    },
    value() {
      this.hash()
    },
  },
  mounted() {
    this.hashAlgorithms = wasmHashGetSupportedHashingAlgorithms() // eslint-disable-line
    this.hash()
  },
  methods: {
    hash() {
      try {
        this.error = false
        this.errorMessages = []

        let data = wasmHash(this.value, this.selectedHashAlgorithm) // eslint-disable-line

        this.outputValue = data
      } catch (error) {
        this.error = true
        this.errorMessages.push('Invalid JWT token.')
      }
    },
    copyText(value) {
      navigator.clipboard.writeText(value)
    },
  },
}
</script>
