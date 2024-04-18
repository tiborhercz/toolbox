<template>
  <div>
    <p>
      Encode text to a base64 string
    </p>
    <v-textarea
      v-model="inputValue"
      name="input-7-1"
      label="Encode"
      outlined
    />
  </div>
</template>

<script>
export default {
  name: 'Base64Encode',
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
            const data = wasmBase64Process(this.value, false, false) // eslint-disable-line
            this.$emit('response', data)
          } else {
            this.$emit('response', '')
          }
        } catch (error) {
          this.$emit('error', 'Could not encode base64')
        }
      },
    },
  },
}
</script>
