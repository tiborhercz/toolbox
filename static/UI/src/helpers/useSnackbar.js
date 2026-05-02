import { reactive } from 'vue'

const state = reactive({ show: false, message: '' })

export function useSnackbar() {
  function showSnackbar(message) {
    state.message = message
    state.show = true
  }
  return { snackbarState: state, showSnackbar }
}
