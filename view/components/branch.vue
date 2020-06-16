<template>
  <span>
    <b-icon-shuffle variant="primary" class="mr-1"></b-icon-shuffle>
    {{ name }}
    <b-badge pill class="hash" variant="secondary">{{ shortSha }}</b-badge>
  </span>
</template>

<script>
export default {
  props: ['name', 'sha'],
  computed: {
    shortSha() {
      return this.sha.slice(0, 5)
    }
  }
}
</script>

<style>
.hash {
  float: right;
}
</style>
