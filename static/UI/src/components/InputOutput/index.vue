<template>
  <v-col>
    <v-row>
      <v-col
        cols="12"
        md="7"
      >
        <v-textarea
          v-model="inputValue"
          name="input-7-1"
          label="Input"
          outlined
        />
      </v-col>
      <v-col
        cols="12"
        md="7"
      >
        <v-textarea
          v-model="outputValue"
          name="input-7-1"
          v-bind:label="'Output'"
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
  </v-col>
</template>

<script>
export default {
  name: 'InputOutput',
  components: {},
  props: {
    type: {
      default: '',
      type: String,
    },
  },
  data() {
    return {
      inputValue: '',
      outputValue: '',
    }
  },
  computed: {
  },
  methods: {
    copyText(value) {
      navigator.clipboard.writeText(value)
    },
  },
}
</script>
