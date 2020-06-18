<template>
  <b-alert :show="show" variant="danger">{{ error }}</b-alert>
</template>

<script>
import { mapState } from 'vuex'
export default {
  computed: {
    ...mapState(['error']),
    show() {
      return this.error !== null
    }
  }
}
</script>

<style></style>
