<template>
  <v-row>
    <v-col cols="12">
      <p class="mb-2">
        Look up DNS records for a domain using DNS-over-HTTPS.
      </p>
      <v-row class="mb-5">
        <v-col
          cols="12"
          sm="6"
          md="4"
        >
          <v-text-field
            v-model="domain"
            placeholder="example.com"
            label="Domain"
            variant="outlined"
            density="compact"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            v-on:keydown.enter="lookup()"
          />
        </v-col>
        <v-col
          cols="12"
          sm="4"
          md="2"
        >
          <v-autocomplete
            v-model="selectedRecordType"
            v-bind:items="recordTypes"
            label="Record type"
            variant="outlined"
            density="compact"
          />
        </v-col>
        <v-col
          cols="12"
          sm="4"
          md="4"
        >
          <v-btn
            class="mr-2"
            v-bind:loading="loading"
            v-on:click="lookup()"
          >
            Lookup
          </v-btn>
          <v-btn
            v-bind:loading="propagationLoading"
            v-on:click="checkPropagation()"
          >
            Check Propagation
          </v-btn>
        </v-col>
      </v-row>

      <v-row v-if="records.length > 0">
        <v-col
          cols="12"
          md="8"
        >
          <v-table>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class="text-left">Type</th>
                  <th class="text-left">Name</th>
                  <th class="text-left">Value</th>
                  <th class="text-left">TTL</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(record, i) in records"
                  v-bind:key="i"
                >
                  <td>{{ record.type }}</td>
                  <td>{{ record.name }}</td>
                  <td>{{ record.data }}</td>
                  <td>{{ record.ttl }}</td>
                </tr>
              </tbody>
            </template>
          </v-table>
        </v-col>
      </v-row>

      <div v-if="propagationResults.length > 0">
        <v-row class="mb-4">
          <v-col
            cols="12"
            md="8"
          >
            <v-alert
              v-if="propagationSummary && propagationSummary.consistent"
              type="success"
              variant="tonal"
            >
              All providers returned consistent records
            </v-alert>
            <v-alert
              v-else-if="propagationSummary"
              type="warning"
              variant="tonal"
            >
              <div class="font-weight-bold mb-2">Differences detected</div>
              <v-table density="compact">
                <template v-slot:default>
                  <thead>
                    <tr>
                      <th class="text-left">Type</th>
                      <th class="text-left">Value</th>
                      <th
                        v-for="provider in comparedProviders"
                        v-bind:key="provider"
                        class="text-center"
                      >
                        {{ provider }}
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="diff in propagationDiffs"
                      v-bind:key="diff.record"
                    >
                      <td>{{ diff.type }}</td>
                      <td class="py-1 text-break">{{ diff.value }}</td>
                      <td
                        v-for="provider in comparedProviders"
                        v-bind:key="provider"
                        class="text-center"
                      >
                        <v-icon
                          v-if="diff.present.includes(provider)"
                          color="success"
                          icon="mdi-check-circle"
                          title="Present"
                        />
                        <v-icon
                          v-else
                          color="error"
                          icon="mdi-close-circle"
                          title="Missing"
                        />
                      </td>
                    </tr>
                  </tbody>
                </template>
              </v-table>
            </v-alert>
          </v-col>
        </v-row>

        <v-row
          v-for="(result, idx) in propagationResults"
          v-bind:key="idx"
          class="mb-4"
        >
          <v-col
            cols="12"
            md="8"
          >
            <h3 class="mb-2">{{ result.provider }}</h3>
            <v-alert
              v-if="result.error"
              type="error"
              density="compact"
              class="mb-2"
            >
              {{ result.error }}
            </v-alert>
            <v-table v-else-if="result.records && result.records.length > 0">
              <template v-slot:default>
                <thead>
                  <tr>
                    <th class="text-left">Type</th>
                    <th class="text-left">Name</th>
                    <th class="text-left">Value</th>
                    <th class="text-left">TTL</th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(record, i) in result.records"
                    v-bind:key="i"
                  >
                    <td>{{ record.type }}</td>
                    <td>{{ record.name }}</td>
                    <td>{{ record.data }}</td>
                    <td>{{ record.ttl }}</td>
                  </tr>
                </tbody>
              </template>
            </v-table>
            <p v-else>No records found</p>
          </v-col>
        </v-row>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed, ref } from 'vue'

const recordTypes = ['ALL', 'A', 'AAAA', 'MX', 'NS', 'TXT', 'CNAME', 'SOA', 'CAA', 'DMARC']

const domain = ref('')
const selectedRecordType = ref('ALL')
const records = ref([])
const propagationResults = ref([])
const propagationSummary = ref(null)
const loading = ref(false)
const propagationLoading = ref(false)
const error = ref(false)
const errorMessages = ref([])

const comparedProviders = computed(() => propagationResults.value
  .filter((result) => !result.error)
  .map((result) => result.provider))

// diff.record is formatted as "TYPE:value"
const propagationDiffs = computed(() => (propagationSummary.value?.diffs || [])
  .map((diff) => {
    const separator = diff.record.indexOf(':')
    return {
      ...diff,
      type: diff.record.slice(0, separator),
      value: diff.record.slice(separator + 1),
    }
  })
  .sort((a, b) => a.record.localeCompare(b.record)))

const lookup = async () => {
  if (!domain.value) {
    error.value = true
    errorMessages.value = ['Please enter a domain']
    return
  }

  error.value = false
  errorMessages.value = []
  records.value = []
  propagationResults.value = []
  propagationSummary.value = null
  loading.value = true

  try {
    const result = await wasmDnsLookup(domain.value, selectedRecordType.value) // eslint-disable-line
    const parsed = JSON.parse(result)
    records.value = parsed || []
  } catch (e) {
    error.value = true
    errorMessages.value = [`DNS lookup failed: ${e}`]
  } finally {
    loading.value = false
  }
}

const checkPropagation = async () => {
  if (!domain.value) {
    error.value = true
    errorMessages.value = ['Please enter a domain']
    return
  }

  error.value = false
  errorMessages.value = []
  records.value = []
  propagationResults.value = []
  propagationSummary.value = null
  propagationLoading.value = true

  try {
    const result = await wasmDnsPropagation(domain.value, selectedRecordType.value) // eslint-disable-line
    const parsed = JSON.parse(result)
    propagationResults.value = parsed.results || []
    propagationSummary.value = parsed.summary || null
  } catch (e) {
    error.value = true
    errorMessages.value = [`Propagation check failed: ${e}`]
  } finally {
    propagationLoading.value = false
  }
}
</script>
