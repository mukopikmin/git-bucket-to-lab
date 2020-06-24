<template>
  <b-button size="sm" variant="primary" :disabled="disabled" @click="action">
    <b-spinner v-if="migrating" small></b-spinner>
    <b-icon-box-seam v-else class="mr-1"></b-icon-box-seam>
    {{ label }}
  </b-button>
</template>

<script>
export default {
  props: {
    action: {
      type: Function,
      required: true
    },
    label: {
      type: String,
      required: true
    },
    migrating: Boolean,
    disable: Boolean
  },
  computed: {
    disabled() {
      return this.disable || this.migrating
    }
  }
}
</script>

<style></style>
