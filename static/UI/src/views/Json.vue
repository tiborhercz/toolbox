<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>JSON beatify</h1>
      </v-col>
    </v-row>
    <json />
  </v-container>
</template>

<script>
import Json from '@/components/tools/Json'

export default {
  name: 'Base64View',
  components: {
    Json,
  },
}
</script>
