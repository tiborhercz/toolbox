<template>
  <v-row>
    <v-col
      cols="12"
      md="7"
    >
      <encode
        v-if="type === 'encode'"
        v-on:response="setResponse"
      />
      <decode
        v-if="type === 'decode'"
        v-on:response="setResponse"
      />
    </v-col>
    <v-col
      cols="12"
      md="7"
    >
      <v-textarea
        v-model="outputValue"
        name="input-7-1"
        label="Output"
        readonly
        outlined
      />
      <v-btn
        v-on:click="copyText(outputValue)"
      >
        Copy
      </v-btn>
    </v-col>
  </v-row>
</template>

<script>
import Encode from '@/components/tools/Base64/Encode'
import Decode from '@/components/tools/Base64/Decode'

export default {
  name: 'Base64EncodeDecode',
  components: {
    Encode,
    Decode,
  },
  props: {
    type: {
      default: '',
      type: String,
    },
  },
  data() {
    return {
      outputValue: '',
    }
  },
  watch: {
    $route() {
      this.outputValue = ''
    },
  },
  methods: {
    setResponse(value) {
      this.outputValue = value
    },
    copyText(value) {
      navigator.clipboard.writeText(value)
    },
  },
}
</script>
>
