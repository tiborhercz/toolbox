<template>
  <v-btn
    v-on:click="copyText"
  >
    {{ label }}
  </v-btn>
</template>

<script>
export default {
  name: 'Button',
  props: {
    label: {
      default: '',
      type: String,
    },
    copyValue: {
      default: '',
      type: String,
    },
  },
  methods: {
    copyText() {
      if (this.copyValue) {
        navigator.clipboard.writeText(this.copyValue)
      }
    },
  },
}
</script>
