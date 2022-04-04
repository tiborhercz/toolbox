<template>
  <div>
    <p>
      Decode a base64 string to text
    </p>
    <v-textarea
      v-model="inputValue"
      name="input-7-1"
      label="Decode"
      outlined
    />
  </div>
</template>

<script>
export default {
  name: 'Base64Decode',
  components: {},
  data() {
    return {
      value: '',
      apiResponseValue: '',
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

        this.error = false
        this.errorMessages = []

        try {
          if (newValue !== '') {
            const data = wasmBase64Process(this.value, true, false) // eslint-disable-line
            this.$emit('response', data)
          } else {
            this.$emit('response', '')
          }
        } catch (error) {
          this.$emit('error', 'Could not decode base64')
        }
      },
    },
  },
}
</script>
