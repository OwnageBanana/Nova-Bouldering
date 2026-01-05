export default {
  root: true,
  env: {
    node: true,
  },
  extends: [
    '@antfu',
    // 'plugin:vue/essential',
    // 'plugin:vue/strongly-recommended',
    // 'plugin:vue/recommended'
    // 'eslint:recommended'
  ],
  parser: 'vue-eslint-parser',
  parserOptions: {
    parser: 'babel-eslint',
    ecmaFeatures: {
      jsx: true,
    },
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    semi: 2,
    'vue/multi-word-component-names': 'off',
  },
}
