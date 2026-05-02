<template>
  <v-row>
    <v-col cols="12">
      <p class="mb-4">
        Calculate Unix file permissions. Toggle checkboxes to build the permission set, or type an octal value directly.
      </p>

      <v-row class="mb-4" align="center">
        <v-col cols="12" sm="4" md="3">
          <v-text-field
            v-model="octalInput"
            label="Octal value"
            variant="outlined"
            density="compact"
            placeholder="755"
            maxlength="4"
            v-on:input="onOctalInput"
          />
        </v-col>
        <v-col cols="12" sm="8" md="6">
          <v-chip class="mr-2" label>
            <strong class="mr-1">Symbolic:</strong> {{ symbolic }}
          </v-chip>
          <v-chip label>
            <strong class="mr-1">chmod</strong> {{ octal }}
          </v-chip>
        </v-col>
      </v-row>

      <v-table>
        <thead>
          <tr>
            <th>Category</th>
            <th class="text-center">Read (4)</th>
            <th class="text-center">Write (2)</th>
            <th class="text-center">Execute (1)</th>
            <th class="text-center">Octal</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="category in categories" :key="category.key">
            <td>{{ category.label }}</td>
            <td class="text-center">
              <v-checkbox
                v-model="category.read"
                density="compact"
                hide-details
                v-on:change="onCheckboxChange"
              />
            </td>
            <td class="text-center">
              <v-checkbox
                v-model="category.write"
                density="compact"
                hide-details
                v-on:change="onCheckboxChange"
              />
            </td>
            <td class="text-center">
              <v-checkbox
                v-model="category.execute"
                density="compact"
                hide-details
                v-on:change="onCheckboxChange"
              />
            </td>
            <td class="text-center">
              <strong>{{ categoryOctal(category) }}</strong>
            </td>
          </tr>
        </tbody>
      </v-table>

      <v-row class="mt-6">
        <v-col cols="12">
          <p class="text-subtitle-2 mb-2">Common presets</p>
          <v-btn
            v-for="preset in presets"
            :key="preset.label"
            variant="tonal"
            size="small"
            class="mr-2 mb-2"
            v-on:click="applyPreset(preset.value)"
          >
            {{ preset.label }} ({{ preset.value }})
          </v-btn>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, computed } from 'vue'

const categories = ref([
  { key: 'owner', label: 'Owner', read: true, write: true, execute: true },
  { key: 'group', label: 'Group', read: true, write: false, execute: true },
  { key: 'others', label: 'Others', read: true, write: false, execute: true },
])

const presets = [
  { label: 'Read only', value: '444' },
  { label: 'Standard file', value: '644' },
  { label: 'Standard dir', value: '755' },
  { label: 'Executable', value: '755' },
  { label: 'Private', value: '600' },
  { label: 'Full access', value: '777' },
]

const octalInput = ref('755')

function categoryOctal(cat) {
  return (cat.read ? 4 : 0) + (cat.write ? 2 : 0) + (cat.execute ? 1 : 0)
}

const octal = computed(() => {
  return categories.value.map(categoryOctal).join('')
})

const symbolic = computed(() => {
  return categories.value.map(cat => {
    return (cat.read ? 'r' : '-') + (cat.write ? 'w' : '-') + (cat.execute ? 'x' : '-')
  }).join('')
})

function onCheckboxChange() {
  octalInput.value = octal.value
}

function applyOctal(value) {
  const digits = value.padStart(3, '0').slice(-3).split('').map(Number)
  categories.value.forEach((cat, i) => {
    const d = digits[i] ?? 0
    cat.read = (d & 4) !== 0
    cat.write = (d & 2) !== 0
    cat.execute = (d & 1) !== 0
  })
}

function onOctalInput() {
  const val = octalInput.value.replace(/[^0-7]/g, '')
  octalInput.value = val
  if (val.length >= 3) {
    applyOctal(val)
  }
}

function applyPreset(value) {
  octalInput.value = value
  applyOctal(value)
}

applyOctal('755')
</script>
