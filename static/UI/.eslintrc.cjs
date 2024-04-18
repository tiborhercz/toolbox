require('@rushstack/eslint-patch/modern-module-resolution')

const path = require('node:path')
const createAliasSetting = require('@vue/eslint-config-airbnb/createAliasSetting')

/* eslint-env node */
module.exports = {
  root: true,
  extends: [
    'plugin:vue/vue3-essential',
    'plugin:vue/vue3-strongly-recommended',
    'plugin:vue/vue3-recommended',
    'eslint:recommended',
    '@vue/eslint-config-airbnb',
  ],
  parserOptions: {
    ecmaVersion: 'latest',
  },
  settings: {
    ...createAliasSetting({
      '@': `${path.resolve(__dirname, './src')}`,
    }),
  },
  rules: {
    semi: ['error', 'never'],
    'vue/v-bind-style': ['error', 'longform'],
    'vue/v-on-style': ['error', 'longform'],
    'vue/v-slot-style': ['error', 'longform'],
    'no-plusplus': ['error', { allowForLoopAfterthoughts: true }],
    'import/no-unresolved': 0,
    'import/extensions': ['error', {
      js: 'never',
      mjs: 'never',
      jsx: 'never',
      ts: 'never',
      tsx: 'never',
      vue: 'never',
    }],
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
  },
}
