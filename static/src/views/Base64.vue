<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>Base64 {{ type }}</h1>
      </v-col>
    </v-row>
    <base64
      v-bind:type="type"
    />
  </v-container>
</template>

<script>
import Base64 from '@/components/tools/Base64'

export default {
  name: 'Base64View',
  components: {
    Base64,
  },
  props: {
    type: {
      default: '',
      type: String,
    },
  },
}
</script>
