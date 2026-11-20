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
          md="6"
          class="d-flex flex-wrap align-start ga-2"
        >
          <v-btn
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
          <v-menu>
            <template v-slot:activator="{ props }">
              <v-btn
                v-bind="props"
                variant="text"
                append-icon="mdi-menu-down"
              >
                Flush Cache
              </v-btn>
            </template>
            <v-list density="compact">
              <v-list-item
                v-for="link in flushLinks"
                v-bind:key="link.name"
                v-bind:href="link.url"
                v-bind:title="link.name"
                target="_blank"
                rel="noopener noreferrer"
                append-icon="mdi-open-in-new"
              />
            </v-list>
          </v-menu>
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
                  <colgroup>
                    <col style="width: 1%">
                    <col style="width: 1%">
                    <col>
                  </colgroup>
                  <tbody>
                    <template
                      v-for="group in providerDiffs"
                      v-bind:key="group.provider"
                    >
                      <tr>
                        <th colspan="3">
                          <div class="d-flex flex-wrap align-center justify-space-between ga-2">
                            <span class="font-weight-bold">{{ group.provider }}</span>
                            <span
                              v-if="group.error"
                              class="font-weight-regular text-medium-emphasis"
                            >
                              <v-icon
                                class="mr-1"
                                icon="mdi-alert-circle-outline"
                                size="small"
                              />
                              Not compared, the lookup failed
                            </span>
                            <span
                              v-else-if="group.rows.length === 0"
                              class="font-weight-regular text-medium-emphasis"
                            >
                              <v-icon
                                class="mr-1"
                                color="success"
                                icon="mdi-check-circle"
                                size="small"
                              />
                              No differences
                            </span>
                            <span
                              v-else
                              class="d-flex ga-2"
                            >
                              <v-chip
                                v-if="group.missing > 0"
                                color="error"
                                size="small"
                                variant="tonal"
                              >
                                {{ group.missing }} missing
                              </v-chip>
                              <v-chip
                                v-if="group.extra > 0"
                                color="info"
                                size="small"
                                variant="tonal"
                              >
                                {{ group.extra }} extra
                              </v-chip>
                            </span>
                          </div>
                        </th>
                      </tr>
                      <tr
                        v-for="row in group.rows"
                        v-bind:key="group.provider + row.record"
                      >
                        <td
                          class="pl-4 pl-sm-8 text-no-wrap"
                          v-bind:title="row.title"
                        >
                          <v-icon
                            size="small"
                            v-bind:color="row.missing ? 'error' : 'info'"
                            v-bind:icon="row.missing ? 'mdi-minus-circle' : 'mdi-plus-circle'"
                          />
                          <span class="d-none d-sm-inline ml-1">{{ row.missing ? 'Missing' : 'Extra' }}</span>
                        </td>
                        <td class="text-no-wrap">{{ row.type }}</td>
                        <td class="py-1 text-break">{{ row.value }}</td>
                      </tr>
                    </template>
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

// Each resolver has its own page for flushing a domain from its cache. Only
// the OpenDNS page takes the domain from the URL.
const flushLinks = computed(() => {
  const openDns = new URL('https://cachecheck.opendns.com/')
  if (domain.value.trim()) {
    openDns.searchParams.set('d', domain.value.trim())
  }

  return [
    { name: 'Cloudflare', url: 'https://one.one.one.one/purge-cache/' },
    { name: 'Google', url: 'https://dns.google/cache' },
    { name: 'OpenDNS', url: openDns.href },
  ]
})

// A record that not every provider returned is listed under the providers on
// the smaller side: as missing where most of the others have it, as extra
// where most of the others do not. diff.record is formatted as "TYPE:value".
const providerDiffs = computed(() => propagationResults.value.map((result) => {
  const diffs = (!result.error && propagationSummary.value?.diffs) || []
  const rows = diffs
    .map((diff) => {
      const missing = diff.missing.includes(result.provider)
      const [same, others] = missing ? [diff.missing, diff.present] : [diff.present, diff.missing]
      const separator = diff.record.indexOf(':')
      return {
        record: diff.record,
        type: diff.record.slice(0, separator),
        value: diff.record.slice(separator + 1),
        missing,
        minority: same.length <= others.length,
        title: `${missing ? 'Returned' : 'Not returned'} by ${others.join(', ')}`,
      }
    })
    .filter((row) => row.minority)
    .sort((a, b) => a.record.localeCompare(b.record))

  return {
    provider: result.provider,
    error: Boolean(result.error),
    rows,
    missing: rows.filter((row) => row.missing).length,
    extra: rows.filter((row) => !row.missing).length,
  }
}))

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
