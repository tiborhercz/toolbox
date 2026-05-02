<script setup>
import { ref, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useTheme } from 'vuetify'
import Drawer from '@/components/DrawerComponent'
import { useSnackbar } from '@/helpers/useSnackbar'

const { snackbarState } = useSnackbar()

const theme = useTheme()

const darkMode = ref(true)

function toggleDarkMode() {
  darkMode.value = !theme.global.current.value.dark
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('dark_theme', theme.global.current.value.dark ? 'true' : 'false')
}

onMounted(() => {
  const LocalDarkTheme = localStorage.getItem('dark_theme')
  if (LocalDarkTheme) {
    theme.global.name.value = LocalDarkTheme === 'true' ? 'dark' : 'light'
    darkMode.value = LocalDarkTheme === 'true'
  }
})
</script>

<style lang="scss">
@use "@/assets/scss/main.scss";
</style>

<template>
  <v-app>
    <v-navigation-drawer
      app
    >
      <v-container>
        <v-row>
          <v-col>
            <router-link class="nostyle" to="/">
              <h1>Toolbox</h1>
            </router-link>
          </v-col>
        </v-row>
      </v-container>
      <Drawer />
      <template v-slot:append>
        <div class="pa-2">
          <v-switch
            v-model="darkMode"
            label="Dark theme"
            v-on:change="toggleDarkMode"
          />
        </div>
      </template>
    </v-navigation-drawer>
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
    <v-snackbar v-model="snackbarState.show" :timeout="1500" color="success" location="bottom right">
      {{ snackbarState.message }}
    </v-snackbar>
  </v-app>
</template>
