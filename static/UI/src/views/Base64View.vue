<script setup>
import Base64 from '@/components/tools/Base64/Base64Component'

const props = defineProps({
  type: {
    type: String,
    default: '',
  },
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>Base64 {{ props.type }}</h1>
      </v-col>
    </v-row>
    <base64
      v-bind:type="props.type"
    />
  </v-container>
</template>
