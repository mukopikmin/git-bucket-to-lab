<template>
  <b-button size="sm" variant="primary" :disabled="migrating" @click="action">
    <b-spinner v-if="migrating" small></b-spinner>
    <b-icon-box-seam v-else class="mr-1"></b-icon-box-seam>
    Migrate
  </b-button>
</template>

<script>
export default {
  props: ['action', 'migrating']
}
</script>

<style></style>
