<template>
  <v-btn
    v-on:click="copyText"
  >
    {{ label }}
  </v-btn>
</template>

<script>
import { useSnackbar } from '@/helpers/useSnackbar'

export default {
  name: 'ButtonComponent',
  setup() {
    const { showSnackbar } = useSnackbar()
    return { showSnackbar }
  },
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
        this.showSnackbar('Copied!')
      }
    },
  },
}
</script>
