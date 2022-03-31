<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>JWT Decode</h1>
      </v-col>
    </v-row>
    <jwt />
  </v-container>
</template>

<script>
import Jwt from '@/components/tools/Jwt'

export default {
  name: 'Base64View',
  components: {
    Jwt,
  },
  props: {
    type: {
      default: '',
      type: String,
    },
  },
}
</script>
